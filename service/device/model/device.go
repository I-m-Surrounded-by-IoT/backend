package model

import (
	"time"
)

type Device struct {
	Mac       string    `gorm:"primarykey;index:,type:hash;type:char(12)"`
	DeviceID  uint64    `gorm:"not null;uniqueIndex"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Comment   string    `gorm:"type:varchar(255)"`
}
