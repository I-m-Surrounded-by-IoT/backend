package log

import (
	"context"

	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/I-m-Surrounded-by-IoT/backend/service/log"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type DeviceLogServer struct {
	group   sarama.ConsumerGroup
	ctx     context.Context
	cancel  context.CancelFunc
	handler sarama.ConsumerGroupHandler
}

func NewDeviceLogServer(
	group sarama.ConsumerGroup,
	c *log.DeviceLogConsumer,
) *DeviceLogServer {
	return &DeviceLogServer{
		group:   group,
		handler: c,
	}
}

func (l *DeviceLogServer) Start(ctx context.Context) error {
	logrus.Infof("start device log consumer...")
	l.ctx, l.cancel = context.WithCancel(ctx)
	err := l.group.Consume(l.ctx, []string{service.KafkaTopicDeviceLog}, l.handler)
	if err != nil {
		logrus.Errorf("failed to consume: %v", err)
	}
	return err
}

func (l *DeviceLogServer) Stop(ctx context.Context) error {
	l.cancel()
	return l.group.Close()
}
