package log

import (
	"strconv"
	"time"

	logApi "github.com/I-m-Surrounded-by-IoT/backend/api/log"
	"github.com/I-m-Surrounded-by-IoT/backend/service/log/model"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
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
	var deviceLog logApi.DeviceLog
	for {
		select {
		case msg := <-msgCh:
			deviceID, err := strconv.ParseUint(string(msg.Key), 10, 64)
			if err != nil {
				log.Errorf("failed to parse device id (%s): %v", msg.Key, err)
				continue
			}
			err = proto.Unmarshal(msg.Value, &deviceLog)
			if err != nil {
				log.Errorf("failed to unmarshal device log (%s): %v", msg.Value, err)
				continue
			}
			err = s.db.CreateDeviceLog(&model.DeviceLog{
				DeviceID:  deviceID,
				Timestamp: time.UnixMilli(deviceLog.Timestamp),
				Message:   deviceLog.Message,
				Level:     log.Level(deviceLog.Level),
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
