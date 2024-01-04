package database

import (
	"log"
	"strings"

	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/IBM/sarama"
	logkafka "github.com/zijiren233/logrus-kafka-hook"
)

func NewLogConsumer(
	k *conf.KafkaConfig,
) sarama.ConsumerGroup {
	client, err := logkafka.NewKafkaClient(
		strings.Split(k.Brokers, ","),
		logkafka.WithKafkaSASLEnable(true),
		logkafka.WithKafkaSASLHandshake(true),
		logkafka.WithKafkaSASLUser(k.User),
		logkafka.WithKafkaSASLPassword(k.Password),
	)
	if err != nil {
		log.Fatalf("failed to create kafka client: %v", err)
	}
	consumerGroup, err := sarama.NewConsumerGroupFromClient(
		"log",
		client,
	)
	if err != nil {
		log.Fatalf("failed to create kafka consumer group: %v", err)
	}
	return consumerGroup
}
