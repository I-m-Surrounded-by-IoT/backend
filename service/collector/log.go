package collector

import (
	"fmt"
	"strconv"

	logApi "github.com/I-m-Surrounded-by-IoT/backend/api/log"
	"github.com/sirupsen/logrus"
	"github.com/zijiren233/stream"
	"google.golang.org/protobuf/proto"
)

var _ logrus.Formatter = (*kafkaDeviceLogKeyFormatter)(nil)

type kafkaDeviceLogKeyFormatter struct{}

func getDeviceID(entry *logrus.Entry) (uint64, error) {
	deviceIDI, ok := entry.Data["device_id"]
	if !ok {
		return 0, fmt.Errorf("missing device_id field")
	}
	deviceID, ok := deviceIDI.(uint64)
	if !ok {
		return 0, fmt.Errorf("invalid device_id type")
	}
	return deviceID, nil
}

func getTopic(entry *logrus.Entry) (string, error) {
	topicI, ok := entry.Data["topic"]
	if !ok {
		return "", fmt.Errorf("missing topic field")
	}
	topic, ok := topicI.(string)
	if !ok {
		return "", fmt.Errorf("invalid topic type")
	}
	return topic, nil
}

func (k *kafkaDeviceLogKeyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	deviceID, err := getDeviceID(entry)
	if err != nil {
		return nil, err
	}
	return stream.StringToBytes(strconv.FormatUint(deviceID, 10)), nil
}

var _ logrus.Formatter = (*kafkaDeviceLogValueFormatter)(nil)

type kafkaDeviceLogValueFormatter struct{}

func (k *kafkaDeviceLogValueFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	deviceID, err := getDeviceID(entry)
	if err != nil {
		return nil, err
	}
	topic, err := getTopic(entry)
	if err != nil {
		return nil, err
	}
	dl := logApi.DeviceLogData{
		DeviceId:  deviceID,
		Topic:     topic,
		Timestamp: entry.Time.UnixMilli(),
		Message:   entry.Message,
		Level:     uint32(entry.Level),
	}
	return proto.Marshal(&dl)
}
