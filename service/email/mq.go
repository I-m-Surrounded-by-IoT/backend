package email

import (
	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/IBM/sarama"

	log "github.com/sirupsen/logrus"
)

var _ sarama.ConsumerGroupHandler = (*EmailConsumer)(nil)

type EmailConsumer struct {
	ms *EmailService
}

func NewEmailConsumer(ms *EmailService) *EmailConsumer {
	return &EmailConsumer{
		ms: ms,
	}
}

func (s *EmailConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *EmailConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *EmailConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Infof("email consumer started")
	msgCh := claim.Messages()
	for {
		select {
		case msg := <-msgCh:
			data, err := service.KafkaTopicEmailUnmarshal(msg.Value)
			if err != nil {
				log.Errorf("failed to unmarshal email (%s): %v", msg.Value, err)
				continue
			}

			_, err = s.ms.SendEmail(session.Context(), data)
			if err != nil {
				log.Errorf("failed to create email: %v", err)
				continue
			}
		case <-session.Context().Done():
			log.Infof("email consumer closed")
			return nil
		}
	}
}
