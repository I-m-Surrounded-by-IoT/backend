package model

import (
	"time"

	"github.com/sirupsen/logrus"
)

type DeviceLog struct {
	DeviceID  uint64    `gorm:"not null;index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Timestamp time.Time `gorm:"not null;index"`
	Level     logrus.Level
	Message   string
}
