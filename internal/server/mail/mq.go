package mail

import (
	"context"

	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/I-m-Surrounded-by-IoT/backend/service/mail"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type MailMQServer struct {
	group   sarama.ConsumerGroup
	ctx     context.Context
	cancel  context.CancelFunc
	handler sarama.ConsumerGroupHandler
}

func NewMailMQServerServer(
	cli sarama.Client,
	c *mail.MailConsumer,
) *MailMQServer {
	group, err := sarama.NewConsumerGroupFromClient(
		service.KafkaTopicMail,
		cli,
	)
	if err != nil {
		logrus.Fatalf("failed to create kafka consumer group: %v", err)
	}
	return &MailMQServer{
		group:   group,
		handler: c,
	}
}

func (l *MailMQServer) Start(ctx context.Context) error {
	logrus.Infof("start device log consumer...")
	l.ctx, l.cancel = context.WithCancel(ctx)
	err := l.group.Consume(l.ctx, []string{service.KafkaTopicDeviceLog}, l.handler)
	if err != nil {
		logrus.Errorf("failed to consume: %v", err)
	}
	return err
}

func (l *MailMQServer) Stop(ctx context.Context) error {
	l.cancel()
	return l.group.Close()
}
