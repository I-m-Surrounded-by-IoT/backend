package notify

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"text/template"
	"time"

	"github.com/Boostport/mjml-go"
	"github.com/I-m-Surrounded-by-IoT/backend/api/email"
	"github.com/I-m-Surrounded-by-IoT/backend/api/message"
	"github.com/I-m-Surrounded-by-IoT/backend/api/notify"
	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/api/waterquality"
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
	userTestEmailTemplate *template.Template
)

type user_test_email_payload struct {
	Username string
	Year     int
}

func init() {
	body, err := mjml.ToHTML(
		context.Background(),
		stream.BytesToString(notify_template.UserTestEmail),
		mjml.WithMinify(true),
	)
	if err != nil {
		log.Fatalf("failed to parse user test template: %v", err)
	}
	userTestEmailTemplate, err = template.New("").Parse(body)
	if err != nil {
		log.Fatalf("failed to parse user test template: %v", err)
	}
}

func (ns *NotifyService) NotifyTestEmail(ctx context.Context, req *notify.NotifyTestEmailReq) (*notify.Empty, error) {
	uif, err := ns.userClient.GetUserInfo(
		ctx,
		&user.GetUserInfoReq{
			Id: req.UserId,
			Fields: []string{
				"id",
				"username",
				"email",
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %w", err)
	}
	if uif.Email == "" {
		return nil, fmt.Errorf("user id: %s, name: %s has no email", uif.Id, uif.Username)
	}
	out := &bytes.Buffer{}
	err = userTestEmailTemplate.Execute(
		out,
		&user_test_email_payload{
			Username: uif.Username,
			Year:     time.Now().Year(),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}
	emailPayload := &email.SendEmailReq{
		To:      []string{uif.Email},
		Subject: "测试邮件",
		Body:    out.String(),
	}
	_, err = ns.emailClient.SendEmail(ctx, emailPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to send mail: %w", err)
	}
	return &notify.Empty{}, nil
}

var (
	deviceOnlineTemplate  *template.Template
	deviceOfflineTemplate *template.Template
)

func newMapUrl(deviceID uint64, geo *waterquality.GeoPoint, t time.Time) (string, error) {
	u, err := url.Parse("https://map.baidu.com/")
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("title", fmt.Sprintf("设备ID:%d上线", deviceID))
	q.Set("content", fmt.Sprintf("坐标:%f,%f\n时间:%s", geo.Lat, geo.Lon, t.Format("2006-01-02 15:04:05")))
	q.Set("autoOpen", "true")
	u.RawQuery = q.Encode()

	return fmt.Sprintf(
		"%s&latlng=%f,%f&l",
		u.String(),
		geo.Lat,
		geo.Lon,
	), nil
}

type device_online_offline_payload struct {
	DeviceId uint64
	Time     string
	*waterquality.GeoPoint
	IP string

	MapUrl string

	QualityTime string
	Temperature float32
	PH          float32
	TSW         float32
	TDS         float32
	Oxygen      float32
	Level       int64

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
	mapUrl, err := newMapUrl(req.DeviceId, req.Report.Data.GeoPoint, time.UnixMilli(req.Seen.LastSeenAt))
	if err != nil {
		return nil, fmt.Errorf("failed to create map url: %w", err)

	}
	out := &bytes.Buffer{}
	err = deviceOnlineTemplate.Execute(
		out,
		&device_online_offline_payload{
			DeviceId: req.DeviceId,
			Time:     time.UnixMilli(req.Seen.LastSeenAt).Format("2006-01-02 15:04:05"),
			GeoPoint: req.Report.Data.GeoPoint,
			IP:       req.Seen.LastSeenIp,
			MapUrl:   mapUrl,

			QualityTime: time.UnixMilli(req.Report.Data.Timestamp).Format("2006-01-02 15:04:05"),
			Temperature: req.Report.Data.Temperature,
			PH:          req.Report.Data.Ph,
			TSW:         req.Report.Data.Tsw,
			TDS:         req.Report.Data.Tds,
			Oxygen:      req.Report.Data.Oxygen,
			Level:       req.Report.Level,

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
	mapUrl, err := newMapUrl(req.DeviceId, req.LastReport.Data.GeoPoint, time.UnixMilli(req.LastSeen.LastSeenAt))
	if err != nil {
		return nil, fmt.Errorf("failed to create map url: %w", err)
	}
	out := &bytes.Buffer{}
	err = deviceOfflineTemplate.Execute(
		out,
		&device_online_offline_payload{
			DeviceId: req.DeviceId,
			Time:     time.UnixMilli(req.LastSeen.LastSeenAt).Format("2006-01-02 15:04:05"),
			GeoPoint: req.LastReport.Data.GeoPoint,
			IP:       req.LastSeen.LastSeenIp,
			MapUrl:   mapUrl,

			QualityTime: time.UnixMilli(req.LastReport.Data.Timestamp).Format("2006-01-02 15:04:05"),
			Temperature: req.LastReport.Data.Temperature,
			PH:          req.LastReport.Data.Ph,
			TSW:         req.LastReport.Data.Tsw,
			TDS:         req.LastReport.Data.Tds,
			Oxygen:      req.LastReport.Data.Oxygen,
			Level:       req.LastReport.Level,

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
