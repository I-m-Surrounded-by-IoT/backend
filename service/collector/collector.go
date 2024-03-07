package collector

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	collection "github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	collectorApi "github.com/I-m-Surrounded-by-IoT/backend/api/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	registryClient "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/IBM/sarama"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kratos/kratos/v2/registry"
	json "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	logkafka "github.com/zijiren233/logrus-kafka-hook"
)

type CollectorService struct {
	deviceClient  device.DeviceClient
	reg           registry.Registrar
	er            *registryClient.EtcdRegistry
	kafkaClient   sarama.Client
	kafkaProducer sarama.AsyncProducer
	mqttClient    mqtt.Client
	// TODO: grpc server
	collectorApi.UnimplementedCollectorServer
}

func NewCollectorService(c *conf.CollectorConfig, k *conf.KafkaConfig, reg registry.Registrar) *CollectorService {
	s := &CollectorService{
		reg: reg,
	}

	opt := mqtt.NewClientOptions().
		AddBroker(c.Mqtt.Addr).
		SetUsername("collector").
		SetClientID(c.Mqtt.ClientId).
		SetPassword(c.Mqtt.Password).
		SetAutoReconnect(true).
		SetOrderMatters(false).
		SetCleanSession(true)
	s.mqttClient = mqtt.NewClient(opt)
	if token := s.mqttClient.Connect(); token.WaitTimeout(time.Second*5) && token.Error() != nil {
		log.Fatalf("failed to connect mqtt server: %v", token.Error())
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
			[]string{service.KafkaTopicDeviceLog},
			logkafka.WithHookMustHasFields([]string{"device_id"}),
			logkafka.WithHookKeyFormatter(&kafkaDeviceLogKeyFormatter{}),
			logkafka.WithHookValueFormatter(&kafkaDeviceLogValueFormatter{}),
			logkafka.WithDynamicTopicsFunc(func(entry *log.Entry) []string {
				deviceIdI, ok := entry.Data["device_id"]
				if !ok {
					return nil
				}
				deviceId, ok := deviceIdI.(uint64)
				if !ok {
					return nil
				}
				return []string{
					fmt.Sprintf("%s-%d", service.KafkaTopicDeviceLog, deviceId),
				}
			}),
		)
		if err != nil {
			log.Fatalf("failed to create kafka hook: %v", err)
		}
		log.Infof("add kafka hook to logrus")
		log.AddHook(lkh)
	}

	producer, err := sarama.NewAsyncProducerFromClient(s.kafkaClient)
	if err != nil {
		log.Fatalf("failed to create kafka producer: %v", err)
	}
	s.kafkaProducer = producer

	if tk := s.mqttClient.Subscribe("$share/collector-device-report/device/+/report", 2, s.handlerDeviceReport); !tk.WaitTimeout(5 * time.Second) {
		log.Fatalf("failed to subscribe topic: %v", tk.Error())
	}

	if tk := s.mqttClient.Subscribe("$share/collector-device-conn/event/device/conn", 2, s.handlerDeviceConn); !tk.WaitTimeout(5 * time.Second) {
		log.Fatalf("failed to subscribe topic: %v", tk.Error())
	}

	log.Info("collector service started")

	return s
}

func (s *CollectorService) handlerDeviceReport(c mqtt.Client, m mqtt.Message) {
	log := log.WithField("topic", m.Topic())

	splited := strings.Split(m.Topic(), "/")
	if len(splited) != 3 {
		log.Errorf("invalid topic: %v", m.Topic())
		return
	}

	id, err := strconv.ParseUint(splited[1], 10, 64)
	if err != nil {
		log.Errorf("failed to parse device id: %v", err)
		return
	}

	log = log.WithField("device_id", id)

	if err := s.UpdateDeviceLastSeen(
		context.Background(),
		id,
		time.Now(),
		"",
	); err != nil {
		log.Errorf("failed to update device last seen: %v", err)
	}

	data := &collection.CollectionData{}
	if err := json.Unmarshal(m.Payload(), data); err != nil {
		log.Errorf("failed to unmarshal report message: %s, %v", m.Payload(), err)
		return
	}

	log.Infof("receive report message: %+v", data)

	if err := s.UpdateDeviceLastReport(context.Background(), id, time.Now(), data); err != nil {
		log.Errorf("failed to update device last report: %v", err)
	}

	err = service.KafkaTopicDeviceReportSend(s.kafkaProducer, id, data)
	if err != nil {
		log.Errorf("failed to send report message to kafka: %v", err)
		return
	}
}

type connMessage struct {
	Time     time.Time `json:"timestamp"`
	Event    string    `json:"event"`
	ClientID string    `json:"clientid"`
	Peername string    `json:"peername"`
}

func (s *CollectorService) handlerDeviceConn(c mqtt.Client, m mqtt.Message) {
	log := log.WithField("topic", m.Topic())

	msg := connMessage{}

	if err := json.Unmarshal(m.Payload(), &msg); err != nil {
		log.Errorf("failed to unmarshal message: %v", err)
		return
	}

	log.Infof("receive device conn message: %+v", msg)

	before, fater, found := strings.Cut(msg.ClientID, "-")
	if !found {
		log.Errorf("invalid client id: %v", msg.ClientID)
		return
	}
	if before != "device" {
		log.Errorf("invalid client id: %v", msg.ClientID)
		return
	}
	deviceID, err := strconv.ParseUint(fater, 10, 64)
	if err != nil {
		log.Errorf("failed to parse device id: %v", err)
		return
	}

	log = log.WithField("device_id", deviceID)

	if err := s.UpdateDeviceLastSeen(
		context.Background(),
		deviceID,
		time.Now(),
		msg.Peername,
	); err != nil {
		log.Errorf("failed to update device last seen: %v", err)
	}

	switch msg.Event {
	case "client.connected":
		// service.KafkaTopicMailSend(s.kafkaProducer, &mail.SendMailReq{})
	case "client.disconnected":
	}
}

func (c *CollectorService) UpdateDeviceLastSeen(ctx context.Context, id uint64, t time.Time, ip string) error {
	_, err := c.deviceClient.UpdateDeviceLastSeen(ctx, &device.UpdateDeviceLastSeenReq{
		Id: id,
		LastSeen: &device.DeviceLastSeen{
			LastSeenAt: t.UnixMilli(),
			LastSeenIp: ip,
		},
	})
	return err
}

func (c *CollectorService) UpdateDeviceLastReport(ctx context.Context, id uint64, t time.Time, data *collection.CollectionData) error {
	req := &device.UpdateDeviceLastReportReq{
		Id: id,
		LastReport: &device.DeviceLastReport{
			LastReportAt: t.UnixMilli(),
			Timestamp:    data.Timestamp,
			Temperature:  data.Temperature,
		},
	}
	if data.GeoPoint != nil {
		req.LastReport.Lat = data.GeoPoint.Lat
		req.LastReport.Lon = data.GeoPoint.Lon
	}
	_, err := c.deviceClient.UpdateDeviceLastReport(ctx, req)
	return err
}

// func (c *CollectorService) ReportImmediately(ctx context.Context, req *collectorApi.ReportImmediatelyReq) (*collectorApi.Empty, error) {
// 	conn, ok := c.getOnlineDeviceConn(req.Id)
// 	if !ok {
// 		return nil, fmt.Errorf("device not online")
// 	}
// 	msg := collector.Message{
// 		Type:    collector.MessageType_ReportImmediately,
// 		Payload: &collector.Message_Empty{},
// 	}
// 	b, err := proto.Marshal(&msg)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = conn.Send(b)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &collectorApi.Empty{}, nil
// }
