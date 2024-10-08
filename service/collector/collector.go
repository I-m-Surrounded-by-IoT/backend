package collector

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	collectorApi "github.com/I-m-Surrounded-by-IoT/backend/api/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/api/notify"
	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/api/waterquality"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	registryClient "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/IBM/sarama"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-kratos/kratos/v2/registry"
	json "github.com/json-iterator/go"
	"github.com/panjf2000/ants/v2"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	logkafka "github.com/zijiren233/logrus-kafka-hook"
	"github.com/zijiren233/timewheel-redis"
)

const (
	device_offline_max_duration = time.Minute * 5
)

type CollectorService struct {
	deviceClient     device.DeviceClient
	kafkaClient      sarama.Client
	kafkaProducer    sarama.AsyncProducer
	mqttClient       mqtt.Client
	userClient       user.UserClient
	collectionClient collection.CollectionClient
	notifyClient     notify.NotifyClient
	timewheel        *timewheel.TimeWheel
	collectorApi.UnimplementedCollectorServer
}

func NewCollectorService(c *conf.CollectorConfig, k *conf.KafkaConfig, reg registry.Registrar, rc *conf.RedisConfig) *CollectorService {
	etcd := reg.(*registryClient.EtcdRegistry)
	discoveryUserConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///user",
		TimeOut:  "10s",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}

	discoveryNotifyConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///notify",
		TimeOut:  "10s",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}

	discoveryCollectionConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///collection",
		TimeOut:  "10s",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Username: rc.Username,
		Password: rc.Password,
		DB:       int(rc.Db),
	})

	s := &CollectorService{
		userClient:       user.NewUserClient(discoveryUserConn),
		notifyClient:     notify.NewNotifyClient(discoveryNotifyConn),
		collectionClient: collection.NewCollectionClient(discoveryCollectionConn),
	}

	s.timewheel = timewheel.NewTimeWheel(
		rdb,
		"collector-timewheel",
		s.twCallback,
		timewheel.WithRetrySleep(time.Second*30),
	)

	go s.timewheel.Run()

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

	cc, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///device",
	}, etcd)
	if err != nil {
		logrus.Fatalf("failed to create grpc conn: %v", err)
	}
	s.deviceClient = device.NewDeviceClient(cc)

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

func (s *CollectorService) twCallback(timer *timewheel.Timer) bool {
	id, err := strconv.ParseUint(timer.Id, 10, 64)
	if err != nil {
		log.Errorf("failed to parse device id: %v", err)
		return false
	}
	lastSeenresp, err := s.deviceClient.GetDeviceLastSeen(
		context.Background(),
		&device.GetDeviceLastSeenReq{
			Id: id,
		},
	)
	if err != nil {
		log.Errorf("failed to get device last seen: %v", err)
		return false
	}
	if time.UnixMilli(lastSeenresp.LastSeenAt).Add(device_offline_max_duration).Before(time.Now()) {
		lastReportResp, err := s.collectionClient.GetDeviceLastReport(
			context.Background(),
			&collection.GetDeviceLastReportReq{
				Id: id,
			},
		)
		if err != nil {
			log.Error("failed to get device last report: %w", err)
			return false
		}
		log.Debugf("device %d last report: %+v", id, lastReportResp)
		_, err = s.notifyClient.NotifyDeviceOffline(
			context.Background(),
			&notify.NotifyDeviceOfflineReq{
				DeviceId:   id,
				Async:      true,
				LastSeen:   lastSeenresp,
				LastReport: lastReportResp,
			},
		)
		if err != nil {
			log.Errorf("failed to notify device offline: %v", err)
		}
	}
	return true
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

	log.Debugf("receive report message: %s", m.Payload())

	ls, err := s.deviceClient.GetDeviceLastSeen(context.Background(), &device.GetDeviceLastSeenReq{
		Id: id,
	})
	if err != nil {
		log.Errorf("failed to get device last seen: %v", err)
	}

	data := &waterquality.Quality{}
	if err := json.Unmarshal(m.Payload(), data); err != nil {
		log.Errorf("failed to unmarshal report message: %v", err)
		return
	}

	// data, err := service.UnmarshalCollectionData(m.Payload())
	// if err != nil {
	// 	log.Errorf("failed to unmarshal report message: %s, %v", m.Payload(), err)
	// 	return
	// }

	err = service.KafkaTopicDeviceReportSend(s.kafkaProducer, id, &collection.CreateCollectionRecordReq{
		DeviceId:   id,
		Data:       data,
		ReceivedAt: time.Now().UnixMilli(),
	})
	if err != nil {
		log.Errorf("failed to send report message to kafka: %v", err)
		return
	}

	if ls != nil {
		if time.UnixMilli(ls.LastSeenAt).Add(device_offline_max_duration).Before(time.Now()) {
			_, err := s.notifyClient.NotifyDeviceOnline(
				context.Background(),
				&notify.NotifyDeviceOnlineReq{
					DeviceId: id,
					Async:    true,
					Seen:     ls,
					Report: &collection.CollectionRecord{
						DeviceId:   id,
						ReceivedAt: time.Now().UnixMilli(),
						Data:       data,
					},
				},
			)
			if err != nil {
				log.Errorf("failed to notify device online: %v", err)
			}
		}
	}

	_ = ants.Submit(func() {
		if err := s.UpdateDeviceLastSeen(
			context.Background(),
			id,
			time.Now(),
			"",
		); err != nil {
			log.Errorf("failed to update device last seen: %v", err)
		}
	})

	err = s.timewheel.AddTimer(
		strconv.FormatUint(id, 10),
		device_offline_max_duration,
		timewheel.WithForce(),
	)
	if err != nil {
		log.Errorf("failed to add timer: %v", err)
	}
}

type connMessage struct {
	Timestamp int64  `json:"timestamp"`
	Event     string `json:"event"`
	ClientID  string `json:"clientid"`
	Peername  string `json:"peername"`
}

func (s *CollectorService) handlerDeviceConn(c mqtt.Client, m mqtt.Message) {
	log := log.WithField("topic", m.Topic())

	msg := connMessage{}

	if err := json.Unmarshal(m.Payload(), &msg); err != nil {
		log.Errorf("failed to unmarshal message: %s, err: %v", m.Payload(), err)
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

	defer func() {
		if err := s.UpdateDeviceLastSeen(
			context.Background(),
			deviceID,
			time.UnixMilli(msg.Timestamp),
			msg.Peername,
		); err != nil {
			log.Errorf("failed to update device last seen: %v", err)
		}
	}()

	// switch msg.Event {
	// case "client.connected":
	// 	ls, err := s.deviceClient.GetDeviceLastSeen(context.Background(), &device.GetDeviceLastSeenReq{
	// 		Id: deviceID,
	// 	})
	// 	if err != nil {
	// 		log.Errorf("failed to get device last seen: %v", err)
	// 		return
	// 	}
	// 	if time.Since(time.UnixMilli(ls.LastSeenAt)).Seconds() > 1 {
	// 		_, err := s.notifyClient.NotifyDeviceOnline(
	// 			context.Background(),
	// 			&notify.NotifyDeviceOnlineReq{
	// 				DeviceId:  deviceID,
	// 				Timestamp: msg.Timestamp,
	// 				Async:     true,
	// 				Ip:        msg.Peername,
	// 			},
	// 		)
	// 		if err != nil {
	// 			log.Errorf("failed to notify device online: %v", err)
	// 		}
	// 	}
	// case "client.disconnected":
	// 	err = s.timewheel.AddTimer(
	// 		strconv.FormatUint(deviceID, 10),
	// 		time.Minute*3,
	// 		timewheel.WithForce(),
	// 	)
	// 	if err != nil {
	// 		log.Errorf("failed to add timer: %v", err)
	// 	}
	// }
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

func (c *CollectorService) ReportNow(ctx context.Context, req *collectorApi.ReportNowReq) (*collectorApi.Empty, error) {
	log := log.WithField("topic", fmt.Sprintf("device/%d/report_now", req.DeviceId))
	log = log.WithField("device_id", req.DeviceId)
	if token := c.mqttClient.Publish(fmt.Sprintf("device/%d/report_now", req.DeviceId), 2, false, "report_now"); !token.WaitTimeout(time.Second * 5) {
		log.Errorf("failed to publish report now data: %v", token.Error())
		return nil, fmt.Errorf("failed to publish report now data: %v", token.Error())
	}
	return &collectorApi.Empty{}, nil
}

func (c *CollectorService) BoatControl(ctx context.Context, req *collectorApi.BoatControlReq) (*collectorApi.Empty, error) {
	log := log.WithField("topic", fmt.Sprintf("device/%d/boat_control", req.DeviceId))
	log = log.WithField("device_id", req.DeviceId)
	cmd := ""
	switch req.Command {
	case collectorApi.BoatControlCommand_FORWARD:
		cmd = "forward"
	case collectorApi.BoatControlCommand_LEFT:
		cmd = "left"
	case collectorApi.BoatControlCommand_RIGHT:
		cmd = "right"
	default:
		log.Errorf("invalid command: %v", req.Command)
		return nil, fmt.Errorf("invalid command: %v", req.Command)
	}
	if token := c.mqttClient.Publish(fmt.Sprintf("device/%d/boat_control", req.DeviceId), 2, false, cmd); !token.WaitTimeout(time.Second * 5) {
		log.Errorf("failed to publish boat control command: %v", token.Error())
		return nil, fmt.Errorf("failed to publish boat control command: %v", token.Error())
	}
	return &collectorApi.Empty{}, nil
}
