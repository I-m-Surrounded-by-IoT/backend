package notify

import (
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
)

var _ sarama.ConsumerGroupHandler = (*NotifyConsumer)(nil)

type NotifyConsumer struct {
	s *NotifyService
}

func NewNotifyConsumer(s *NotifyService) *NotifyConsumer {
	return &NotifyConsumer{
		s: s,
	}
}

func (s *NotifyConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *NotifyConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *NotifyConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Infof("notify consumer started")
	// msgCh := claim.Messages()
	for {
		select {
		// case msg := <-msgCh:
		// 	s.s.No

		// 	if err != nil {
		// 		log.Errorf("failed to create device log: %v", err)
		// 		continue
		// 	}
		case <-session.Context().Done():
			log.Infof("device log consumer closed")
			return nil
		}
	}
}
