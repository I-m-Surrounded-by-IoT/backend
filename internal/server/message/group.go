package message

import (
	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

func NewConsumerGroup(
	cli sarama.Client,
) sarama.ConsumerGroup {
	group, err := sarama.NewConsumerGroupFromClient(
		service.KafkaTopicMessage,
		cli,
	)
	if err != nil {
		logrus.Fatalf("failed to create kafka consumer group: %v", err)
	}
	return group
}
