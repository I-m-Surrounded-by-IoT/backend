package collector

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	collection "github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	collectorApi "github.com/I-m-Surrounded-by-IoT/backend/api/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	registryClient "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/proto/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	tcpconn "github.com/I-m-Surrounded-by-IoT/backend/utils/tcpConn"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/registry"
	log "github.com/sirupsen/logrus"
	logkafka "github.com/zijiren233/logrus-kafka-hook"
	"google.golang.org/protobuf/proto"
)

type CollectorService struct {
	deviceClient  device.DeviceClient
	grpcEndpoint  string
	reg           registry.Registrar
	er            *registryClient.EtcdRegistry
	kafkaClient   sarama.Client
	kafkaProducer sarama.AsyncProducer
	dlcr          *DeviceStreamLogRegistor
	// TODO: grpc server
	collectorApi.UnimplementedCollectorServer
}

func NewCollectorService(c *conf.CollectorConfig, k *conf.KafkaConfig, reg registry.Registrar) *CollectorService {
	s := &CollectorService{
		reg:  reg,
		dlcr: NewDeviceStreamLogRegistor(),
	}
	switch reg := reg.(type) {
	case *registryClient.EtcdRegistry:
		cc, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
			Endpoint: "discovery:///device",
		}, reg)
		if err != nil {
			panic(err)
		}
		s.deviceClient = device.NewDeviceClient(cc)
		s.er = reg
	default:
		panic("invalid registry")
	}

	if k == nil || k.Brokers == "" {
		log.Fatal("kafka config is empty")
	} else {
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
		client, err := logkafka.NewKafkaClient(
			strings.Split(k.Brokers, ","),
			opts...,
		)
		if err != nil {
			log.Fatalf("failed to create kafka client: %v", err)
		}
		s.kafkaClient = client
		lkh, err := logkafka.NewLogKafkaHookFromClient(
			client,
			[]string{"log-device"},
			logkafka.WithHookMustHasFields([]string{"device_id"}),
			logkafka.WithHookKeyFormatter(&kafkaDeviceLogKeyFormatter{}),
			logkafka.WithHookValueFormatter(&kafkaDeviceLogValueFormatter{}),
		)
		if err != nil {
			log.Fatalf("failed to create kafka hook: %v", err)
		}
		log.Infof("add kafka hook to logrus")
		log.AddHook(lkh)
	}
	log.AddHook(s.dlcr)

	producer, err := sarama.NewAsyncProducerFromClient(s.kafkaClient)
	if err != nil {
		log.Fatalf("failed to create kafka producer: %v", err)
	}
	s.kafkaProducer = producer
	return s
}

func (c *CollectorService) SetGrpcEndpoint(endpoint string) {
	c.grpcEndpoint = endpoint
}

func (c *CollectorService) UpdateDeviceLastSeen(ctx context.Context, id uint64) {
	_, err := c.deviceClient.UpdateDeviceLastSeen(ctx, &device.UpdateDeviceLastSeenReq{
		Id:       id,
		LastSeen: time.Now().UnixMilli(),
	})
	if err != nil {
		log.Errorf("update device last seen failed: %v", err)
	}
}

func (c *CollectorService) RegisterDevice(ctx context.Context, d *device.DeviceInfo) (func() error, error) {
	devicdService := &registry.ServiceInstance{
		ID:   d.Mac,
		Name: fmt.Sprintf("device-%v", d.Id),
		Endpoints: []string{
			c.grpcEndpoint,
		},
	}
	return func() error {
		return c.reg.Deregister(context.Background(), devicdService)
	}, c.reg.Register(ctx, devicdService)
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

	d, err := c.deviceClient.GetDeviceInfoByMac(ctx, &device.GetDeviceInfoByMacReq{
		Mac: msg.GetMac(),
	})
	if err != nil {
		return fmt.Errorf("find or create device failed: %w", err)
	}

	log := log.WithField("device_id", d.Id)
	dlc, err := c.dlcr.RegisterDevice(d.Id)
	if err != nil {
		log.Errorf("register device log chan failed: %v", err)
		return fmt.Errorf("register device log chan failed: %w", err)
	}
	defer dlc.Close()
	defer c.dlcr.UnregisterDevice(d.Id, dlc)

	c.UpdateDeviceLastSeen(ctx, d.Id)

	dereg, err := c.RegisterDevice(ctx, d)
	if err != nil {
		log.Errorf("register device failed: %v", err)
		return fmt.Errorf("register device failed: %w", err)
	}
	defer func() {
		err := dereg()
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
		c.UpdateDeviceLastSeen(ctx, d.Id)
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
			info := &collection.CollectionRecord{
				DeviceId:  d.Id,
				Timestamp: payload.Timestamp,
				GeoPoint: &collection.GeoPoint{
					Lng: payload.GeoPoint.Longitude,
					Lat: payload.GeoPoint.Latitude,
				},
				Temperature: payload.Temperature,
			}
			bytes, err := proto.Marshal(info)
			if err != nil {
				log.Errorf("failed to marshal collection info: %v", err)
				continue
			}

			// TODO: 添加数据后处理服务
			topics := []string{"device-collection-report"}
			for _, topic := range topics {
				c.kafkaProducer.Input() <- &sarama.ProducerMessage{
					Topic: topic,
					Key:   sarama.StringEncoder(strconv.FormatUint(d.Id, 10)),
					Value: sarama.ByteEncoder(bytes),
				}
			}

		case collector.MessageType_ReportLog:
			payload := msg.GetLogPayload()
			if payload == nil {
				log.Errorf("invalid report log payload: %v", msg.Payload)
				continue
			}
			switch payload.Level {
			case collector.LogLevel_LogLevelDebug:
				log.Debugf("device report: time: %v, message: %v", time.UnixMilli(int64(payload.Timestamp)).Format(time.DateTime), payload.Message)
			case collector.LogLevel_LogLevelInfo:
				log.Infof("device report: time: %v, message: %v", time.UnixMilli(int64(payload.Timestamp)).Format(time.DateTime), payload.Message)
			case collector.LogLevel_LogLevelWarning:
				log.Warnf("device report: time: %v, message: %v", time.UnixMilli(int64(payload.Timestamp)).Format(time.DateTime), payload.Message)
			case collector.LogLevel_LogLevelError:
				log.Errorf("device report: time: %v, message: %v", time.UnixMilli(int64(payload.Timestamp)).Format(time.DateTime), payload.Message)
			default:
				log.Errorf("device report invalid log level: %v, message: %v", payload.Level, payload.Message)
			}

		default:
			log.Errorf("invalid message type: %v", msg.Type)
			continue
		}
	}
}

func (c *CollectorService) GetDeviceStreamLog(req *collectorApi.GetDeviceStreamLogReq, resp collectorApi.Collector_GetDeviceStreamLogServer) error {
	dlc, ok := c.dlcr.GetDeviceLogChans(req.Id)
	if !ok {
		return fmt.Errorf("device log chan not found")
	}
	ch, f, err := dlc.Watch(req.LevelFilter)
	if err != nil {
		return err
	}
	defer f()
	for {
		select {
		case log := <-ch:
			err := resp.Send(&collectorApi.GetDeviceStreamLogResp{
				Level:     log.Level,
				Message:   log.Message,
				Timestamp: log.Time.UnixMilli(),
			})
			if err != nil {
				return err
			}
		case <-resp.Context().Done():
			return nil
		}
	}
}
