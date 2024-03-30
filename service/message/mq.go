package message

import (
	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/IBM/sarama"

	log "github.com/sirupsen/logrus"
)

var _ sarama.ConsumerGroupHandler = (*MessageConsumer)(nil)

type MessageConsumer struct {
	ms *MessageService
}

func NewMessageConsumer(ms *MessageService) *MessageConsumer {
	return &MessageConsumer{
		ms: ms,
	}
}

func (s *MessageConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *MessageConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *MessageConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Infof("email consumer started")
	msgCh := claim.Messages()
	for {
		select {
		case msg := <-msgCh:
			data, err := service.KafkaTopicMessageUnmarshal(msg.Value)
			if err != nil {
				log.Errorf("failed to unmarshal email (%s): %v", msg.Value, err)
				continue
			}

			_, err = s.ms.SendMessage(session.Context(), data)
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
