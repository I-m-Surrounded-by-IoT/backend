package collection

import (
	"context"

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
	consumerGroup sarama.ConsumerGroup,
	c *collection.CollectionConsumer,
) *CollectionConsumerServer {
	return &CollectionConsumerServer{
		consumerGroup: consumerGroup,
		handler:       c,
	}
}

func (l *CollectionConsumerServer) Start(ctx context.Context) error {
	log.Infof("start log consumer...")
	l.ctx, l.cancel = context.WithCancel(ctx)
	err := l.consumerGroup.Consume(l.ctx, []string{"device-collection-report"}, l.handler)
	if err != nil {
		log.Errorf("failed to consume: %v", err)
	}
	return err
}

func (l *CollectionConsumerServer) Stop(ctx context.Context) error {
	l.cancel()
	return l.consumerGroup.Close()
}
