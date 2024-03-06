package mail

import (
	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/IBM/sarama"

	log "github.com/sirupsen/logrus"
)

var _ sarama.ConsumerGroupHandler = (*MailConsumer)(nil)

type MailConsumer struct {
	ms *MailService
}

func NewMailConsumer(ms *MailService) *MailConsumer {
	return &MailConsumer{
		ms: ms,
	}
}

func (s *MailConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *MailConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *MailConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Infof("device log consumer started")
	msgCh := claim.Messages()
	for {
		select {
		case msg := <-msgCh:
			data, err := service.KafkaTopicMailUnmarshal(msg.Value)
			if err != nil {
				log.Errorf("failed to unmarshal device log (%s): %v", msg.Value, err)
				continue
			}

			_, err = s.ms.SendMail(session.Context(), data)
			if err != nil {
				log.Errorf("failed to create device log: %v", err)
				continue
			}
		case <-session.Context().Done():
			log.Infof("device log consumer closed")
			return nil
		}
	}
}
