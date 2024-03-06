package collection

import (
	"context"

	"github.com/I-m-Surrounded-by-IoT/backend/service"
	collection "github.com/I-m-Surrounded-by-IoT/backend/service/collection"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
)

type CollectionConsumerServer struct {
	consumerGroup sarama.ConsumerGroup
	ctx           context.Context
	cancel        context.CancelFunc
	handler       sarama.ConsumerGroupHandler
}

func NewCollectionConsumerServer(
	kc sarama.Client,
	c *collection.CollectionConsumer,
) *CollectionConsumerServer {
	consumerGroup, err := sarama.NewConsumerGroupFromClient(
		"collection",
		kc,
	)
	if err != nil {
		log.Fatalf("failed to create kafka consumer group: %v", err)
	}
	return &CollectionConsumerServer{
		consumerGroup: consumerGroup,
		handler:       c,
	}
}

func (l *CollectionConsumerServer) Start(ctx context.Context) error {
	log.Infof("start log consumer...")
	l.ctx, l.cancel = context.WithCancel(ctx)
	err := l.consumerGroup.Consume(l.ctx, []string{service.KafkaTopicDeviceReport}, l.handler)
	if err != nil {
		log.Errorf("failed to consume: %v", err)
	}
	return err
}

func (l *CollectionConsumerServer) Stop(ctx context.Context) error {
	l.cancel()
	return l.consumerGroup.Close()
}
