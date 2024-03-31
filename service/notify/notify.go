package notify

import (
	"bytes"
	"context"
	"fmt"
	"text/template"
	"time"

	"github.com/Boostport/mjml-go"
	"github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/api/email"
	"github.com/I-m-Surrounded-by-IoT/backend/api/message"
	"github.com/I-m-Surrounded-by-IoT/backend/api/notify"
	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	registryClient "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/service"
	notify_template "github.com/I-m-Surrounded-by-IoT/backend/service/notify/template"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/registry"
	log "github.com/sirupsen/logrus"
	"github.com/zijiren233/stream"
	"golang.org/x/exp/maps"
)

type NotifyService struct {
	emailClient   email.EmailClient
	messageClient message.MessageClient
	userClient    user.UserClient
	kafkaClient   sarama.Client
	kafkaProducer sarama.AsyncProducer
	notify.UnimplementedNotifyServer
}

func NewNotifyService(dc *conf.NotifyConfig, k *conf.KafkaConfig, reg registry.Registrar) *NotifyService {
	etcd := reg.(*registryClient.EtcdRegistry)
	discoveryEmailConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///email",
		TimeOut:  "10s",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}
	emailClient := email.NewEmailClient(discoveryEmailConn)

	discoveryUserConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///user",
		TimeOut:  "10s",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}
	userClient := user.NewUserClient(discoveryUserConn)

	discoveryMessageConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///message",
		TimeOut:  "10s",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}
	messageClient := message.NewMessageClient(discoveryMessageConn)

	kafkaClient, err := utils.DailKafka(k)
	if err != nil {
		log.Fatalf("failed to create kafka conn: %v", err)
	}

	kafkaProducer, err := sarama.NewAsyncProducerFromClient(kafkaClient)
	if err != nil {
		log.Fatalf("failed to create kafka producer: %v", err)
	}

	// discoveryDeviceConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
	// 	Endpoint: "discovery:///device",
	// 	TimeOut:  "10s",
	// }, etcd)
	// if err != nil {
	// 	log.Fatalf("failed to create grpc conn: %v", err)
	// }
	// deviceClient := device.NewDeviceClient(discoveryDeviceConn)

	return &NotifyService{
		emailClient:   emailClient,
		messageClient: messageClient,
		userClient:    userClient,
		kafkaClient:   kafkaClient,
		kafkaProducer: kafkaProducer,
	}
}

var (
	deviceOnlineTemplate  *template.Template
	deviceOfflineTemplate *template.Template
)

type device_online_payload struct {
	DeviceId uint64
	Time     string
	*collection.GeoPoint
	IP string

	QualityTime string
	Temperature float32
	Ph          float32

	Year int
}

func init() {
	body, err := mjml.ToHTML(
		context.Background(),
		stream.BytesToString(notify_template.DeviceOnline),
		mjml.WithMinify(true),
	)
	if err != nil {
		log.Fatalf("failed to parse device online template: %v", err)
	}
	deviceOnlineTemplate, err = template.New("").Parse(body)
	if err != nil {
		log.Fatalf("failed to parse device online template: %v", err)
	}

	body, err = mjml.ToHTML(
		context.Background(),
		stream.BytesToString(notify_template.DeviceOffline),
		mjml.WithMinify(true),
	)
	if err != nil {
		log.Fatalf("failed to parse device offline template: %v", err)
	}
	deviceOfflineTemplate, err = template.New("").Parse(body)
	if err != nil {
		log.Fatalf("failed to parse device offline template: %v", err)
	}
}

func (s *NotifyService) NotifyDeviceOnline(ctx context.Context, req *notify.NotifyDeviceOnlineReq) (*notify.Empty, error) {
	resp, err := s.userClient.ListFollowedUserNotificationMethodsByDevice(context.Background(), &user.ListFollowedUserNotificationMethodsByDeviceReq{
		DeviceId: req.DeviceId,
	})
	if err != nil {
		log.Errorf("failed to get followed user: %v", err)
		return nil, fmt.Errorf("failed to get followed user: %w", err)
	}
	if len(resp.UserNotificationMethods) == 0 {
		log.Debugf("no followed user for device %d", req.DeviceId)
		return nil, nil
	}
	subject := fmt.Sprintf("设备 %d 上线", req.DeviceId)
	out := &bytes.Buffer{}
	err = deviceOnlineTemplate.Execute(
		out,
		&device_online_payload{
			DeviceId: req.DeviceId,
			Time:     time.UnixMilli(req.Seen.LastSeenAt).Format("2006-01-02 15:04:05"),
			GeoPoint: req.Report.Data.GeoPoint,
			IP:       req.Seen.LastSeenIp,

			QualityTime: time.UnixMilli(req.Report.Data.Timestamp).Format("2006-01-02 15:04:05"),
			Temperature: req.Report.Data.Temperature,
			Ph:          req.Report.Data.Ph,

			Year: time.Now().Year(),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}
	methods := maps.Values(resp.UserNotificationMethods)
	emails := make([]string, len(methods))
	for i, m := range methods {
		emails[i] = m.Email
	}
	emailPayload := &email.SendEmailReq{
		To:      emails,
		Subject: subject,
		Body:    out.String(),
	}
	if req.Async {
		err = service.KafkaTopicEmailSend(s.kafkaProducer, emailPayload)
	} else {
		_, err = s.emailClient.SendEmail(ctx, emailPayload)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to send mail: %w", err)
	}
	messagePayload := &message.SendMessageReq{
		UserId: maps.Keys(resp.UserNotificationMethods),
		Payload: &message.MessagePayload{
			MessageType: message.MessageType_TYPE_DEVICE_ONLINE,
			Title:       subject,
			Content:     out.String(),
		},
	}
	if req.Async {
		err = service.KafkaTopicMessageSend(s.kafkaProducer, messagePayload)
	} else {
		_, err = s.messageClient.SendMessage(ctx, messagePayload)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}
	return &notify.Empty{}, nil
}

func (s *NotifyService) NotifyDeviceOffline(ctx context.Context, req *notify.NotifyDeviceOfflineReq) (*notify.Empty, error) {
	resp, err := s.userClient.ListFollowedUserNotificationMethodsByDevice(context.Background(), &user.ListFollowedUserNotificationMethodsByDeviceReq{
		DeviceId: req.DeviceId,
	})
	if err != nil {
		log.Errorf("failed to get followed user: %v", err)
		return nil, fmt.Errorf("failed to get followed user: %w", err)
	}
	if len(resp.UserNotificationMethods) == 0 {
		log.Debugf("no followed user for device %d", req.DeviceId)
		return nil, nil
	}
	subject := fmt.Sprintf("设备 %d 离线", req.DeviceId)
	out := &bytes.Buffer{}
	err = deviceOfflineTemplate.Execute(
		out,
		req,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}
	methods := maps.Values(resp.UserNotificationMethods)
	emails := make([]string, len(methods))
	for i, m := range methods {
		emails[i] = m.Email
	}
	emailPayload := &email.SendEmailReq{
		To:      emails,
		Subject: subject,
		Body:    out.String(),
	}
	if req.Async {
		err = service.KafkaTopicEmailSend(s.kafkaProducer, emailPayload)
	} else {
		_, err = s.emailClient.SendEmail(ctx, emailPayload)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to send mail: %w", err)
	}
	messagePayload := &message.SendMessageReq{
		UserId: maps.Keys(resp.UserNotificationMethods),
		Payload: &message.MessagePayload{
			MessageType: message.MessageType_TYPE_DEVICE_OFFLINE,
			Title:       subject,
			Content:     out.String(),
		},
	}
	if req.Async {
		err = service.KafkaTopicMessageSend(s.kafkaProducer, messagePayload)
	} else {
		_, err = s.messageClient.SendMessage(ctx, messagePayload)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}
	return &notify.Empty{}, nil
}
