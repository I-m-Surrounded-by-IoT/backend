package notify

import (
	"context"
	"fmt"

	"github.com/I-m-Surrounded-by-IoT/backend/api/email"
	"github.com/I-m-Surrounded-by-IoT/backend/api/message"
	"github.com/I-m-Surrounded-by-IoT/backend/api/notify"
	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	registryClient "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/registry"
	log "github.com/sirupsen/logrus"
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
	subject := fmt.Sprintf("device %d %s", req.DeviceId, "online")
	body := formatDeviceOnlineBody(req.DeviceId, req.Timestamp)
	methods := maps.Values(resp.UserNotificationMethods)
	emails := make([]string, len(methods))
	for i, m := range methods {
		emails[i] = m.Email
	}
	emailPayload := &email.SendEmailReq{
		To:      emails,
		Subject: subject,
		Body:    body,
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
			Timestamp:   req.Timestamp,
			MessageType: message.MessageType_TYPE_DEVICE_ONLINE,
			Title:       subject,
			Content:     body,
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
	subject := fmt.Sprintf("device %d %s", req.DeviceId, "offline")
	body := formatDeviceOfflineBody(req.DeviceId, req.Timestamp)
	methods := maps.Values(resp.UserNotificationMethods)
	emails := make([]string, len(methods))
	for i, m := range methods {
		emails[i] = m.Email
	}
	emailPayload := &email.SendEmailReq{
		To:      emails,
		Subject: subject,
		Body:    body,
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
			Timestamp:   req.Timestamp,
			MessageType: message.MessageType_TYPE_DEVICE_OFFLINE,
			Title:       subject,
			Content:     body,
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
