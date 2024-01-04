package database

import (
	"context"

	"github.com/I-m-Surrounded-by-IoT/backend/service/database"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type DeviceLogConsumer struct {
	consumerGroup sarama.ConsumerGroup
	ctx           context.Context
	cancel        context.CancelFunc
	deviceLog     sarama.ConsumerGroupHandler
}

func NewDeviceLogConsumer(
	consumerGroup sarama.ConsumerGroup,
	c *database.DeviceLogConsumer,
) *DeviceLogConsumer {
	return &DeviceLogConsumer{
		consumerGroup: consumerGroup,
		deviceLog:     c,
	}
}

func (l *DeviceLogConsumer) Start(ctx context.Context) error {
	logrus.Infof("start log consumer...")
	l.ctx, l.cancel = context.WithCancel(ctx)
	err := l.consumerGroup.Consume(l.ctx, []string{"device-log"}, l.deviceLog)
	if err != nil {
		logrus.Errorf("failed to consume: %v", err)
		return err
	}
	return err
}

func (l *DeviceLogConsumer) Stop(ctx context.Context) error {
	l.cancel()
	return l.consumerGroup.Close()
}
