package notify

import (
	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
)

var _ sarama.ConsumerGroupHandler = (*UserNotifyConsumer)(nil)

type UserNotifyConsumer struct {
	userClient user.UserClient
}

func NewUserNotifyConsumer(userClient user.UserClient) *UserNotifyConsumer {
	return &UserNotifyConsumer{
		userClient: userClient,
	}
}

func (s *UserNotifyConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *UserNotifyConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *UserNotifyConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Infof("user mail notify consumer started")
	// msgCh := claim.Messages()
	for {
		select {
		// case msg := <-msgCh:
		// data, err := service.KafkaTopicMailUnmarshal(msg.Value)
		// if err != nil {
		// 	log.Errorf("failed to unmarshal device log (%s): %v", msg.Value, err)
		// 	continue
		// }

		// if err != nil {
		// 	log.Errorf("failed to create device log: %v", err)
		// 	continue
		// }
		case <-session.Context().Done():
			log.Infof("device log consumer closed")
			return nil
		}
	}
}
