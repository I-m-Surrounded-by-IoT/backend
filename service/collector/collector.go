package collector

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	collectorApi "github.com/I-m-Surrounded-by-IoT/backend/api/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/api/database"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	registryClient "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/proto/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	tcpconn "github.com/I-m-Surrounded-by-IoT/backend/utils/tcpConn"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	log "github.com/sirupsen/logrus"
	logkafka "github.com/zijiren233/logrus-kafka-hook"
	"google.golang.org/protobuf/proto"
)

type CollectorService struct {
	db           database.DatabaseClient
	grpcEndpoint string
	er           *registryClient.EtcdRegistry
	cr           *registryClient.ConsulRegistry
	// TODO: grpc server
	collectorApi.UnimplementedCollectorServer
}

func (c *CollectorService) SetGrpcEndpoint(endpoint string) {
	c.grpcEndpoint = endpoint
}

func (c *CollectorService) ServeTcp(ctx context.Context, conn net.Conn) error {
	log.Infof("receive connection from collector: %v", conn.RemoteAddr())
	Conn := tcpconn.NewConn(conn)
	defer Conn.Close()
	err := Conn.SayHello()
	if err != nil {
		return fmt.Errorf("say hello to collector failed: %w", err)
	}
	b, err := Conn.NextMessage()
	if err != nil {
		return fmt.Errorf("receive message from collector failed: %w", err)
	}
	msg := collector.Message{}
	err = proto.Unmarshal(b, &msg)
	if err != nil {
		return fmt.Errorf("unmarshal message from collector failed: %w", err)
	}
	if msg.Type != collector.MessageType_ReportMac {
		return fmt.Errorf("invalid first message type: %v", msg.Type)
	}

	device, err := c.db.FirstOrCreateDevice(ctx, &database.CreateDeviceReq{
		Mac: msg.GetMac(),
	})
	if err != nil {
		return fmt.Errorf("find or create device failed: %w", err)
	}

	log := log.WithField("device_id", device.DeviceId)

	devicdService := &registry.ServiceInstance{
		ID:   device.Mac,
		Name: fmt.Sprintf("device-%v", device.DeviceId),
		Metadata: map[string]string{
			"endpoint": c.grpcEndpoint,
		},
	}

	log.Infof("register device to registry: %v", devicdService)
	var reg registry.Registrar
	switch {
	case c.er != nil:
		reg = etcd.New(c.er.Client(), etcd.Context(ctx))
	case c.cr != nil:
		reg = consul.New(c.cr.Client(), consul.WithHealthCheck(false))
	default:
		log.Errorf("etcd or consul registry is nil")
		return fmt.Errorf("etcd or consul registry is nil")
	}

	err = reg.Register(ctx, devicdService)
	if err != nil {
		return fmt.Errorf("register device failed: %w", err)
	}
	log.Infof("register device to registry: %v", devicdService)

	defer func() {
		log.Infof("deregister device from registry: %v", devicdService)
		err := reg.Deregister(context.Background(), devicdService)
		if err != nil {
			log.Errorf("deregister device failed: %v", err)
		}
	}()

	msg.Type = collector.MessageType_ReportImmediately
	msg.Payload = &collector.Message_Empty{}
	b, err = proto.Marshal(&msg)
	if err != nil {
		return fmt.Errorf("marshal message to collector failed: %w", err)
	}
	err = Conn.Send(b)
	if err != nil {
		return fmt.Errorf("send message to collector failed: %w", err)
	}

	for {
		b, err := Conn.NextMessage()
		if err != nil {
			return fmt.Errorf("receive message from collector failed: %w", err)
		}
		msg = collector.Message{}
		err = proto.Unmarshal(b, &msg)
		if err != nil {
			return fmt.Errorf("unmarshal message from collector failed: %w", err)
		}
		switch msg.Type {
		case collector.MessageType_Heartbeat:
			log.Infof("receive heartbeat message from collector: %v", msg.Payload)
			continue
		case collector.MessageType_Report:
			payload := msg.GetReportPayload()
			if payload == nil {
				log.Errorf("invalid report payload: %v", msg.Payload)
				continue
			}
			log.Infof("receive report message from collector: %v", payload)
			_, err = c.db.CreateCollectionInfo(ctx, &database.CollectionInfo{
				DeviceId:  device.DeviceId,
				Timestamp: payload.Timestamp,
				GeoPoint: &database.GeoPoint{
					Lng: payload.GeoPoint.Longitude,
					Lat: payload.GeoPoint.Latitude,
				},
				Temperature: payload.Temperature,
			})
			if err != nil {
				return fmt.Errorf("create collection info failed: %w", err)
			}
		case collector.MessageType_ReportLog:
			payload := msg.GetLogPayload()
			if payload == nil {
				log.Errorf("invalid report log payload: %v", msg.Payload)
				continue
			}
			switch payload.Level {
			case collector.LogLevel_LogLevelDebug:
				log.Debugf("device report: time: %v, message: %v", time.UnixMicro(int64(payload.Timestamp)).Format(time.DateTime), payload.Message)
			case collector.LogLevel_LogLevelInfo:
				log.Infof("device report: time: %v, message: %v", time.UnixMicro(int64(payload.Timestamp)).Format(time.DateTime), payload.Message)
			case collector.LogLevel_LogLevelWarning:
				log.Warnf("device report: time: %v, message: %v", time.UnixMicro(int64(payload.Timestamp)).Format(time.DateTime), payload.Message)
			case collector.LogLevel_LogLevelError:
				log.Errorf("device report: time: %v, message: %v", time.UnixMicro(int64(payload.Timestamp)).Format(time.DateTime), payload.Message)
			default:
				log.Errorf("device report invalid log level: %v, message: %v", payload.Level, payload.Message)
			}

		default:
			log.Errorf("invalid message type: %v", msg.Type)
			continue
		}
	}
}

func NewCollectorService(c *conf.CollectorConfig, k *conf.KafkaConfig, reg registry.Registrar) *CollectorService {
	s := &CollectorService{}
	switch reg := reg.(type) {
	case *registryClient.EtcdRegistry:
		cc, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
			Endpoint: "discovery:///database",
			Tls:      c.DatabaseTls,
		}, reg)
		if err != nil {
			panic(err)
		}
		s.db = database.NewDatabaseClient(cc)
		s.er = reg
	case *registryClient.ConsulRegistry:
		cc, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
			Endpoint: "discovery:///database",
			Tls:      c.DatabaseTls,
		}, reg)
		if err != nil {
			panic(err)
		}
		s.db = database.NewDatabaseClient(cc)
		s.cr = reg
	default:
		panic("invalid registry")
	}

	if k != nil && k.Brokers != "" {
		opts := []logkafka.KafkaOptionFunc{
			logkafka.WithKafkaSASLHandshake(true),
			logkafka.WithKafkaSASLUser(k.User),
			logkafka.WithKafkaSASLPassword(k.Password),
		}
		if k.User != "" || k.Password != "" {
			opts = append(opts,
				logkafka.WithKafkaSASLEnable(true),
			)
		}
		lkh, err := logkafka.NewLogKafkaHook(
			strings.Split(k.Brokers, ","),
			[]string{"device-log"},
			opts,
			logkafka.WithLogKafkaHookMustHasFields([]string{"device_id"}),
			logkafka.WithLogKafkaHookKeyFormatter(new(kafkaLogKeyFormatter)),
		)
		if err != nil {
			log.Fatalf("failed to create kafka hook: %v", err)
		}
		log.Infof("add kafka hook to logrus")
		log.AddHook(lkh)
	} else {
		log.Warnf("kafka config is empty")
	}

	return s
}
