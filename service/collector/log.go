package collector

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var _ logrus.Formatter = (*kafkaLogKeyFormatter)(nil)

type kafkaLogKeyFormatter struct{}

func (k *kafkaLogKeyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
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
