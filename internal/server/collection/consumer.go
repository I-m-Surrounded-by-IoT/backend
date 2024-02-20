package collection

import (
	"strings"

	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
	logkafka "github.com/zijiren233/logrus-kafka-hook"
)

func NewLogConsumer(
	k *conf.KafkaConfig,
) sarama.ConsumerGroup {
	opts := []logkafka.KafkaOptionFunc{
		logkafka.WithKafkaSASLHandshake(true),
		logkafka.WithKafkaSASLUser(k.User),
		logkafka.WithKafkaSASLPassword(k.Password),
	}
	if k.User != "" || k.Password != "" {
		opts = append(opts,
			logkafka.WithKafkaSASLEnable(true),
			logkafka.WithKafkaSASLUser(k.User),
			logkafka.WithKafkaSASLPassword(k.Password),
		)
	}
	client, err := logkafka.NewKafkaClient(
		strings.Split(k.Brokers, ","),
		opts...,
	)
	if err != nil {
		log.Fatalf("failed to create kafka client: %v", err)
	}
	consumerGroup, err := sarama.NewConsumerGroupFromClient(
		"device-collection-report",
		client,
	)
	if err != nil {
		log.Fatalf("failed to create kafka consumer group: %v", err)
	}
	return consumerGroup
}
