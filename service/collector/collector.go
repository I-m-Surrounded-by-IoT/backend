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
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/IBM/sarama"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kratos/kratos/v2/registry"
	json "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	logkafka "github.com/zijiren233/logrus-kafka-hook"
	"github.com/zijiren233/stream"
	"google.golang.org/protobuf/proto"
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

	producer, err := sarama.NewAsyncProducerFromClient(s.kafkaClient)
	if err != nil {
		log.Fatalf("failed to create kafka producer: %v", err)
	}
	s.kafkaProducer = producer

	tk := s.mqttClient.Subscribe("$share/collector/device/+/report", 2, func(c mqtt.Client, m mqtt.Message) {
		splited := strings.Split(m.Topic(), "/")
		if len(splited) != 3 {
			log.Errorf("invalid topic: %v", m.Topic())
			return
		}
		id, err := strconv.ParseUint(splited[1], 10, 64)
		if err != nil {
			log.Errorf("failed to parse device id, topic: %v, error: %v", m.Topic(), err)
			return
		}

		log := log.WithField("device_id", id)

		if err := s.UpdateDeviceLastSeen(context.Background(), id); err != nil {
			log.Errorf("failed to update device last seen: %v", err)
		}

		data := &collection.CollectionData{}
		if err := json.Unmarshal(m.Payload(), data); err != nil {
			log.Errorf("failed to unmarshal message: %v", err)
			return
		}

		log.Infof("receive report message: %+v", data)

		if err := s.UpdateDeviceLastReport(context.Background(), id, time.Now(), data); err != nil {
			log.Errorf("failed to update device last report: %v", err)
		}

		bytes, err := proto.Marshal(data)
		if err != nil {
			log.Errorf("failed to marshal info: %v", err)
			return
		}

		// TODO: 添加数据后处理服务
		topics := []string{"device-collection-report"}
		for _, topic := range topics {
			s.kafkaProducer.Input() <- &sarama.ProducerMessage{
				Topic: topic,
				Key:   sarama.StringEncoder(strconv.FormatUint(id, 10)),
				Value: sarama.ByteEncoder(bytes),
			}
		}
	})
	if !tk.WaitTimeout(5 * time.Second) {
		log.Fatalf("failed to subscribe topic: %v", tk.Error())
	}

	log.Info("collector service started")

	return s
}

func (c *CollectorService) UpdateDeviceLastSeen(ctx context.Context, id uint64) error {
	_, err := c.deviceClient.UpdateDeviceLastSeen(ctx, &device.UpdateDeviceLastSeenReq{
		Id: id,
		LastSeen: &device.DeviceLastSeen{
			LastSeenAt: time.Now().UnixMilli(),
		},
	})
	return err
}

func (c *CollectorService) UpdateDeviceLastReport(ctx context.Context, id uint64, t time.Time, data *collection.CollectionData) error {
	_, err := c.deviceClient.UpdateDeviceLastReport(ctx, &device.UpdateDeviceLastReportReq{
		Id: id,
		LastReport: &device.DeviceLastReport{
			LastReportAt: t.UnixMilli(),
			Timestamp:    data.Timestamp,
			Lat:          data.GeoPoint.Lat,
			Lon:          data.GeoPoint.Lon,
			Temperature:  data.Temperature,
		},
	})
	return err
}

func (c *CollectorService) GetDeviceStreamReport(req *collectorApi.GetDeviceStreamReportReq, resp collectorApi.Collector_GetDeviceStreamReportServer) error {
	cxt, cancel := context.WithCancel(resp.Context())
	defer cancel()
	topic := fmt.Sprintf("device/%d/report", req.Id)
	sub := c.mqttClient.Subscribe(topic, 2, func(client mqtt.Client, message mqtt.Message) {
		select {
		case <-cxt.Done():
			return
		default:
			data := &collection.CollectionData{}
			if err := json.Unmarshal(message.Payload(), data); err != nil {
				log.Errorf("failed to unmarshal message: %v", err)
				cancel()
				return
			}
			err := resp.Send(data)
			if err != nil {
				cancel()
			}
		}
	})
	select {
	case <-sub.Done():
		return sub.Error()
	case <-cxt.Done():
		if token := c.mqttClient.Unsubscribe(topic); !token.WaitTimeout(time.Second * 5) {
			log.Errorf("failed to unsubscribe topic: %v", token.Error())
		}
		return cxt.Err()
	}
}

func (c *CollectorService) GetDeviceStreamEvent(req *collectorApi.GetDeviceStreamEventReq, resp collectorApi.Collector_GetDeviceStreamEventServer) error {
	cxt, cancel := context.WithCancel(resp.Context())
	defer cancel()
	// req.EventFilter
	// c.mqttClient.SubscribeMultiple()
	topic := fmt.Sprintf("device/%d/#", req.Id)
	sub := c.mqttClient.Subscribe(topic, 2, func(client mqtt.Client, message mqtt.Message) {
		select {
		case <-cxt.Done():
			return
		default:
			err := resp.Send(&collectorApi.GetDeviceStreamEventResp{
				Message: stream.BytesToString(message.Payload()),
			})
			if err != nil {
				cancel()
			}
		}
	})
	select {
	case <-sub.Done():
		return sub.Error()
	case <-cxt.Done():
		if token := c.mqttClient.Unsubscribe(topic); !token.WaitTimeout(time.Second * 5) {
			log.Errorf("failed to unsubscribe topic: %v", token.Error())
		}
		return nil
	}
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
