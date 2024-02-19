package collector

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var _ logrus.Formatter = (*kafkaDeviceLogKeyFormatter)(nil)

type kafkaDeviceLogKeyFormatter struct{}

func (k *kafkaDeviceLogKeyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	deviceIDI, ok := entry.Data["device_id"]
	if !ok {
		return nil, fmt.Errorf("missing device_id field")
	}
	deviceID, ok := deviceIDI.(uint64)
	if !ok {
		return nil, fmt.Errorf("invalid device_id type")
	}
	return []byte(fmt.Sprintf("%v", deviceID)), nil
}
