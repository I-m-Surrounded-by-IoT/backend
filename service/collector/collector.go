package collector

import (
	"context"
	"fmt"
	"net"
	"strings"

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
	"github.com/sirupsen/logrus"
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
	Conn := tcpconn.NewConn(conn)
	defer Conn.Close()
	logrus.Infof("receive connection from collector: %v", conn.RemoteAddr())
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

	log := logrus.WithField("device_id", device.DeviceId)

	devicdService := &registry.ServiceInstance{
		ID:   device.Mac,
		Name: fmt.Sprintf("device-%v", device.DeviceId),
		Metadata: map[string]string{
			"endpoint": c.grpcEndpoint,
		},
	}

	var reg registry.Registrar
	switch {
	case c.er != nil:
		reg = etcd.New(c.er.Client(), etcd.Context(ctx))
	case c.cr != nil:
		reg = consul.New(c.cr.Client(), consul.WithHealthCheck(false))
	default:
		return fmt.Errorf("invalid registry")
	}

	err = reg.Register(ctx, devicdService)
	if err != nil {
		return fmt.Errorf("register device failed: %w", err)
	}
	log.Infof("register device to registry: %v", devicdService)

	defer func() {
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
				log.Errorf("invalid message payload: %v", msg.Payload)
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
			})
			if err != nil {
				return fmt.Errorf("create collection info failed: %w", err)
			}
		default:
			log.Errorf("invalid message type: %v", msg.Type)
			continue
		}
	}
}

func NewCollectorService(c *conf.CollectorConfig, reg registry.Registrar) *CollectorService {
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

	if c.Kafka != nil && c.Kafka.Brokers != "" {
		lkh, err := logkafka.NewLogKafkaHook(
			strings.Split(c.Kafka.Brokers, ","),
			[]string{"device-log"},
			[]logkafka.KafkaOptionFunc{
				logkafka.WithKafkaSASLEnable(true),
				logkafka.WithKafkaSASLHandshake(true),
				logkafka.WithKafkaSASLUser(c.Kafka.User),
				logkafka.WithKafkaSASLPassword(c.Kafka.Password),
			},
			logkafka.WithLogKafkaHookMustHasFields([]string{"device_id"}),
			logkafka.WithLogKafkaHookKeyFormatter(new(kafkaLogKeyFormatter)),
		)
		if err != nil {
			logrus.Fatalf("failed to create kafka hook: %v", err)
		}
		logrus.Infof("add kafka hook to logrus")
		logrus.AddHook(lkh)
	} else {
		logrus.Warnf("kafka config is empty")
	}

	return s
}
