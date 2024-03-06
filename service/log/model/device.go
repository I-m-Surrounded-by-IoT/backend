package model

import (
	"time"

	"github.com/sirupsen/logrus"
)

type DeviceLog struct {
	ID        uint64       `gorm:"primaryKey"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
	DeviceID  uint64       `gorm:"not null;index:idx_device_id_timestamp"`
	Topic     string       `gorm:"index"`
	Timestamp time.Time    `gorm:"not null;index:idx_device_id_timestamp"`
	Level     logrus.Level `gorm:"not null;index"`
	Message   string       `gorm:"not null;type:text"`
}
