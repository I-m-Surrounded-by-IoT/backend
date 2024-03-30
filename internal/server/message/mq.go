package message

import (
	"context"

	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/I-m-Surrounded-by-IoT/backend/service/message"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type MessageMQServer struct {
	group   sarama.ConsumerGroup
	ctx     context.Context
	cancel  context.CancelFunc
	handler sarama.ConsumerGroupHandler
}

func NewMessageMQServerServer(
	cli sarama.Client,
	c *message.MessageConsumer,
) *MessageMQServer {
	group, err := sarama.NewConsumerGroupFromClient(
		service.KafkaTopicMessage,
		cli,
	)
	if err != nil {
		logrus.Fatalf("failed to create kafka consumer group: %v", err)
	}
	return &MessageMQServer{
		group:   group,
		handler: c,
	}
}

func (l *MessageMQServer) Start(ctx context.Context) error {
	logrus.Infof("start message consumer...")
	l.ctx, l.cancel = context.WithCancel(ctx)
	err := l.group.Consume(l.ctx, []string{service.KafkaTopicMessage}, l.handler)
	if err != nil {
		logrus.Errorf("failed to consume: %v", err)
	}
	return err
}

func (l *MessageMQServer) Stop(ctx context.Context) error {
	l.cancel()
	return l.group.Close()
}
