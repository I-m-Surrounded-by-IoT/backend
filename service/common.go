package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/api/email"
	"github.com/I-m-Surrounded-by-IoT/backend/api/log"
	"github.com/I-m-Surrounded-by-IoT/backend/api/message"
	"github.com/IBM/sarama"
	"google.golang.org/protobuf/proto"
)

const (
	KafkaTopicDeviceReport = "device-report"
)

func KafkaTopicDeviceReportUnmarshal(data []byte) (*collection.CreateCollectionRecordReq, error) {
	v := &collection.CreateCollectionRecordReq{}
	err := proto.Unmarshal(data, v)
	return v, err
}

func KafkaTopicDeviceReportUnmarshalTo(data []byte, v *collection.CreateCollectionRecordReq) error {
	return proto.Unmarshal(data, v)
}

func KafkaTopicDeviceReportSend(kc sarama.AsyncProducer, deviceID uint64, data *collection.CreateCollectionRecordReq) error {
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
			Topic:     topic,
			Key:       sarama.StringEncoder(strconv.FormatUint(deviceID, 10)),
			Value:     sarama.ByteEncoder(bytes),
			Timestamp: time.UnixMilli(data.Data.Timestamp),
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
	KafkaTopicEmail = "email"
)

func KafkaTopicEmailUnmarshal(data []byte) (*email.SendEmailReq, error) {
	v := &email.SendEmailReq{}
	err := proto.Unmarshal(data, v)
	return v, err
}

func KafkaTopicEmailUnmarshalTo(data []byte, v *email.SendEmailReq) error {
	return proto.Unmarshal(data, v)
}

func KafkaTopicEmailSend(kc sarama.AsyncProducer, data *email.SendEmailReq) error {
	bytes, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	kc.Input() <- &sarama.ProducerMessage{
		Topic: KafkaTopicEmail,
		Value: sarama.ByteEncoder(bytes),
	}
	return nil
}

func UnmarshalCollectionData(data []byte) (*collection.CollectionData, error) {
	v := &collection.CollectionData{}
	err := proto.Unmarshal(data, v)
	return v, err
}

func UnmarshalCollectionDataTo(data []byte, v *collection.CollectionData) error {
	return proto.Unmarshal(data, v)
}

const (
	KafkaTopicMessage = "message"
)

func KafkaTopicMessageUnmarshal(data []byte) (*message.SendMessageReq, error) {
	v := &message.SendMessageReq{}
	err := proto.Unmarshal(data, v)
	return v, err
}

func KafkaTopicMessageUnmarshalTo(data []byte, v *message.SendMessageReq) error {
	return proto.Unmarshal(data, v)
}

func KafkaTopicMessageSend(kc sarama.AsyncProducer, data *message.SendMessageReq) error {
	bytes, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	kc.Input() <- &sarama.ProducerMessage{
		Topic: KafkaTopicMessage,
		Value: sarama.ByteEncoder(bytes),
	}
	return nil
}
