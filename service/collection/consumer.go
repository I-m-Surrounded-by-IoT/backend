package collection

import (
	"context"
	"strconv"

	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/zijiren233/stream"
)

var _ sarama.ConsumerGroupHandler = (*CollectionConsumer)(nil)

type CollectionConsumer struct {
	s *CollectionService
}

func NewCollectionConsumer(s *CollectionService) *CollectionConsumer {
	return &CollectionConsumer{
		s: s,
	}
}

func (s *CollectionConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *CollectionConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *CollectionConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Infof("start consume device report...")
	msgCh := claim.Messages()
	for {
		select {
		case msg := <-msgCh:
			_, err := strconv.ParseUint(stream.BytesToString(msg.Key), 10, 64)
			if err != nil {
				log.Errorf("failed to parse device id (%s): %v", stream.BytesToString(msg.Key), err)
				continue
			}
			data, err := service.KafkaTopicDeviceReportUnmarshal(msg.Value)
			if err != nil {
				log.Errorf("failed to unmarshal device report (%s): %v", msg.Value, err)
				continue
			}
			_, err = s.s.CreateCollectionRecord(context.Background(), data)
			if err != nil {
				log.Errorf("failed to create collection record: %v", err)
				continue
			}
		case <-session.Context().Done():
			log.Infof("stop consume device report...")
			return nil
		}
	}
}
