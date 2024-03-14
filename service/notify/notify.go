package notify

import (
	"context"
	"fmt"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/email"
	"github.com/I-m-Surrounded-by-IoT/backend/api/notify"
	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	registryClient "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/registry"
	log "github.com/sirupsen/logrus"
)

type NotifyService struct {
	emailClient   email.EmailClient
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
		userClient:    userClient,
		kafkaClient:   kafkaClient,
		kafkaProducer: kafkaProducer,
	}
}

func (s *NotifyService) NotifyDeviceOnline(ctx context.Context, req *notify.NotifyDeviceOnlineReq) (*notify.Empty, error) {
	resp, err := s.userClient.ListFollowedUserEmailsByDevice(context.Background(), &user.ListFollowedUserEmailsByDeviceReq{
		DeviceId: req.DeviceId,
	})
	if err != nil {
		log.Errorf("failed to get followed user: %v", err)
		return nil, fmt.Errorf("failed to get followed user: %w", err)
	}
	if len(resp.UserEmails) == 0 {
		log.Debugf("no followed user for device %d", req.DeviceId)
		return nil, nil
	}
	payload := &email.SendEmailReq{
		To:      resp.UserEmails,
		Subject: fmt.Sprintf("device %d %s", req.DeviceId, "online"),
		Body:    fmt.Sprintf("device %d %s at %s", req.DeviceId, "online", time.UnixMilli(req.Timestamp).Format(time.RFC3339)),
	}
	if req.Async {
		err = service.KafkaTopicEmailSend(s.kafkaProducer, payload)
	} else {
		_, err = s.emailClient.SendEmail(ctx, payload)
	}
	if err != nil {
		log.Errorf("failed to send mail: %v", err)
		return nil, fmt.Errorf("failed to send mail: %w", err)
	}
	return &notify.Empty{}, nil
}

func (s *NotifyService) NotifyDeviceOffline(ctx context.Context, req *notify.NotifyDeviceOfflineReq) (*notify.Empty, error) {
	resp, err := s.userClient.ListFollowedUserEmailsByDevice(context.Background(), &user.ListFollowedUserEmailsByDeviceReq{
		DeviceId: req.DeviceId,
	})
	if err != nil {
		log.Errorf("failed to get followed user: %v", err)
		return nil, fmt.Errorf("failed to get followed user: %w", err)
	}
	if len(resp.UserEmails) == 0 {
		log.Debugf("no followed user for device %d", req.DeviceId)
		return nil, nil
	}
	payload := &email.SendEmailReq{
		To:      resp.UserEmails,
		Subject: fmt.Sprintf("device %d %s", req.DeviceId, "offline"),
		Body:    fmt.Sprintf("device %d %s at %s", req.DeviceId, "offline", time.UnixMilli(req.Timestamp).Format(time.RFC3339)),
	}
	if req.Async {
		err = service.KafkaTopicEmailSend(s.kafkaProducer, payload)
	} else {
		_, err = s.emailClient.SendEmail(ctx, payload)
	}
	if err != nil {
		log.Errorf("failed to send mail: %v", err)
		return nil, fmt.Errorf("failed to send mail: %w", err)
	}
	return &notify.Empty{}, nil
}
