package log

import (
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/I-m-Surrounded-by-IoT/backend/service/log/model"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
)

var _ sarama.ConsumerGroupHandler = (*DeviceLogConsumer)(nil)

type DeviceLogConsumer struct {
	db *dbUtils
}

func NewDeviceLogConsumer(dbs *LogService) *DeviceLogConsumer {
	return &DeviceLogConsumer{
		db: dbs.db,
	}
}

func (s *DeviceLogConsumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *DeviceLogConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (s *DeviceLogConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Infof("device log consumer started")
	msgCh := claim.Messages()
	for {
		select {
		case msg := <-msgCh:
			data, err := service.KafkaTopicDeviceLogUnmarshal(msg.Value)
			if err != nil {
				log.Errorf("failed to unmarshal device log (%s): %v", msg.Value, err)
				continue
			}
			err = s.db.CreateDeviceLog(&model.DeviceLog{
				DeviceID:  data.DeviceId,
				Topic:     data.Topic,
				Timestamp: time.UnixMilli(data.Timestamp),
				Message:   data.Message,
				Level:     log.Level(data.Level),
			})
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
