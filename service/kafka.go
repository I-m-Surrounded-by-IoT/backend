package service

import (
	"github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/api/log"
	"github.com/I-m-Surrounded-by-IoT/backend/api/mail"
	"google.golang.org/protobuf/proto"
)

const (
	KafkaTopicDeviceReport = "device-report"
)

func KafkaTopicDeviceReportUnmarshal(data []byte) (*collection.CollectionData, error) {
	v := &collection.CollectionData{}
	err := proto.Unmarshal(data, v)
	return v, err
}

func KafkaTopicDeviceReportUnmarshalTo(data []byte, v *collection.CollectionData) error {
	return proto.Unmarshal(data, v)
}

const (
	KafkaTopicDeviceLog = "device-log"
)

func KafkaTopicDeviceLogUnmarshal(data []byte) (*log.DeviceLogData, error) {
	v := &log.DeviceLogData{}
	err := proto.Unmarshal(data, v)
	return v, err
}

func KafkaTopicDeviceLogUnmarshalTo(data []byte, v *log.DeviceLogData) error {
	return proto.Unmarshal(data, v)
}

const (
	KafkaTopicMail = "mail"
)

func KafkaTopicMailUnmarshal(data []byte) (*mail.SendMailReq, error) {
	v := &mail.SendMailReq{}
	err := proto.Unmarshal(data, v)
	return v, err
}

func KafkaTopicMailUnmarshalTo(data []byte, v *mail.SendMailReq) error {
	return proto.Unmarshal(data, v)
}
