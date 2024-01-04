package database

import (
	"strconv"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/service/database/model"
	"github.com/IBM/sarama"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

var _ sarama.ConsumerGroupHandler = (*DeviceLogConsumer)(nil)

type DeviceLogConsumer struct {
	db *dbUtils
}

func NewDeviceLogConsumer(dbs *DatabaseService) *DeviceLogConsumer {
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

type DeviceLog struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
	Time  string `json:"time"`
}

func (s *DeviceLogConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	log.Infof("device log consumer started")
	msgCh := claim.Messages()
	var deviceLog DeviceLog
	for {
		select {
		case msg := <-msgCh:
			log.Infof("receive device log message: key: %v, value: %v", string(msg.Key), string(msg.Value))
			deviceID, err := strconv.ParseUint(string(msg.Key), 10, 64)
			if err != nil {
				log.Errorf("failed to parse device id (%s): %v", msg.Key, err)
				continue
			}
			err = jsoniter.Unmarshal(msg.Value, &deviceLog)
			if err != nil {
				log.Errorf("failed to unmarshal device log (%s): %v", msg.Value, err)
				continue
			}
			level, err := log.ParseLevel(deviceLog.Level)
			if err != nil {
				log.Errorf("failed to parse log level (%s): %v", deviceLog.Level, err)
				continue
			}
			t, err := time.Parse(time.DateTime, deviceLog.Time)
			if err != nil {
				log.Errorf("failed to parse log time (%s): %v", deviceLog.Time, err)
				continue
			}
			err = s.db.CreateDeviceLog(&model.DeviceLog{
				DeviceID:  deviceID,
				Level:     level,
				Message:   deviceLog.Msg,
				Timestamp: t,
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
