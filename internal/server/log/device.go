package log

import (
	"context"

	"github.com/I-m-Surrounded-by-IoT/backend/service/log"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type DeviceLogServer struct {
	consumerGroup sarama.ConsumerGroup
	ctx           context.Context
	cancel        context.CancelFunc
	deviceLog     sarama.ConsumerGroupHandler
}

func NewDeviceLogServer(
	consumerGroup sarama.ConsumerGroup,
	c *log.DeviceLogConsumer,
) *DeviceLogServer {
	return &DeviceLogServer{
		consumerGroup: consumerGroup,
		deviceLog:     c,
	}
}

func (l *DeviceLogServer) Start(ctx context.Context) error {
	logrus.Infof("start log consumer...")
	l.ctx, l.cancel = context.WithCancel(ctx)
	err := l.consumerGroup.Consume(l.ctx, []string{"device-log"}, l.deviceLog)
	if err != nil {
		logrus.Errorf("failed to consume: %v", err)
	}
	return err
}

func (l *DeviceLogServer) Stop(ctx context.Context) error {
	l.cancel()
	return l.consumerGroup.Close()
}
