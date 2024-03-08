package email

import (
	"context"

	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/I-m-Surrounded-by-IoT/backend/service/email"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type EmailMQServer struct {
	group   sarama.ConsumerGroup
	ctx     context.Context
	cancel  context.CancelFunc
	handler sarama.ConsumerGroupHandler
}

func NewEmailMQServerServer(
	cli sarama.Client,
	c *email.EmailConsumer,
) *EmailMQServer {
	group, err := sarama.NewConsumerGroupFromClient(
		service.KafkaTopicEmail,
		cli,
	)
	if err != nil {
		logrus.Fatalf("failed to create kafka consumer group: %v", err)
	}
	return &EmailMQServer{
		group:   group,
		handler: c,
	}
}

func (l *EmailMQServer) Start(ctx context.Context) error {
	logrus.Infof("start email consumer...")
	l.ctx, l.cancel = context.WithCancel(ctx)
	err := l.group.Consume(l.ctx, []string{service.KafkaTopicEmail}, l.handler)
	if err != nil {
		logrus.Errorf("failed to consume: %v", err)
	}
	return err
}

func (l *EmailMQServer) Stop(ctx context.Context) error {
	l.cancel()
	return l.group.Close()
}
