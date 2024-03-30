package captcha

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/Boostport/mjml-go"
	captchaApi "github.com/I-m-Surrounded-by-IoT/backend/api/captcha"
	"github.com/I-m-Surrounded-by-IoT/backend/api/email"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	registryClient "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/service"
	captcha_template "github.com/I-m-Surrounded-by-IoT/backend/service/captcha/template"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	logkafka "github.com/zijiren233/logrus-kafka-hook"
	"github.com/zijiren233/stream"
)

type CaptchaService struct {
	kafkaProducer sarama.AsyncProducer
	emailClient   email.EmailClient
	cache         *CaptchaRcache
	captchaApi.UnimplementedCaptchaServer
}

func NewCaptchaService(c *conf.CaptchaConfig, k *conf.KafkaConfig, reg registry.Registrar, rc *conf.RedisConfig) *CaptchaService {
	etcd := reg.(*registryClient.EtcdRegistry)
	discoveryEmailConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///email",
		TimeOut:  "10s",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}
	emailClient := email.NewEmailClient(discoveryEmailConn)

	rdb := redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Username: rc.Username,
		Password: rc.Password,
		DB:       int(rc.Db),
	})

	s := &CaptchaService{
		cache:       NewCaptchaRcache(rdb),
		emailClient: emailClient,
	}

	if k == nil || k.Brokers == "" {
		log.Fatal("kafka config is empty")
	}
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

	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Fatalf("failed to create kafka producer: %v", err)
	}
	s.kafkaProducer = producer

	return s
}

var (
	captchaTemplate *template.Template
)

func init() {
	body, err := mjml.ToHTML(
		context.Background(),
		stream.BytesToString(captcha_template.Captcha),
		mjml.WithMinify(true),
	)
	if err != nil {
		log.Fatalf("failed to parse mjml: %v", err)
	}
	captchaTemplate = template.Must(template.New("").Parse(body))
}

func (cs *CaptchaService) SendEmailCaptcha(ctx context.Context, req *captchaApi.SendEmailCaptchaReq) (*captchaApi.Empty, error) {
	if req.UserId == "" {
		return nil, errors.New("user id is required")
	}
	if req.Email == "" {
		return nil, errors.New("email is required")
	}
	captcha, err := cs.cache.NewMailCaptcha(ctx, req.UserId, req.Email)
	if err != nil {
		return nil, err
	}
	out := &bytes.Buffer{}
	err = captchaTemplate.Execute(out, struct {
		Captcha string
		Year    int
	}{
		Captcha: captcha,
		Year:    time.Now().Year(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}
	emailReq := &email.SendEmailReq{
		To:      []string{req.Email},
		Subject: "验证码",
		Body:    out.String(),
	}
	if req.Async {
		err = service.KafkaTopicEmailSend(
			cs.kafkaProducer,
			emailReq,
		)
	} else {
		_, err = cs.emailClient.SendEmail(
			ctx,
			emailReq,
		)
	}
	if err != nil {
		return nil, err
	}
	return &captchaApi.Empty{}, nil
}

func (cs *CaptchaService) VerifyEmailCaptcha(ctx context.Context, req *captchaApi.VerifyEmailCaptchaReq) (*captchaApi.Empty, error) {
	if req.UserId == "" {
		return nil, errors.New("user id is required")
	}
	if req.Email == "" {
		return nil, errors.New("email is required")
	}
	return &captchaApi.Empty{}, cs.cache.VerifyEmailCaptcha(ctx, req.UserId, req.Email, req.Captcha)
}
