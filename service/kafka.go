package service

import (
	"fmt"
	"strconv"

	"github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/api/log"
	"github.com/I-m-Surrounded-by-IoT/backend/api/mail"
	"github.com/IBM/sarama"
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

func KafkaTopicDeviceReportSend(kc sarama.AsyncProducer, deviceID uint64, data *collection.CollectionData) error {
	bytes, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	topics := []string{
		KafkaTopicDeviceReport,
		fmt.Sprintf("%s-%d", KafkaTopicDeviceReport, deviceID),
	}
	for _, topic := range topics {
		kc.Input() <- &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(strconv.FormatUint(deviceID, 10)),
			Value: sarama.ByteEncoder(bytes),
		}
	}
	return nil
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

func KafkaTopicMailSend(kc sarama.AsyncProducer, data *mail.SendMailReq) error {
	bytes, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	kc.Input() <- &sarama.ProducerMessage{
		Topic: KafkaTopicMail,
		Value: sarama.ByteEncoder(bytes),
	}
	return nil
}
