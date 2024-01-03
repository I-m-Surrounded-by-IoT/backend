package model

import (
	"time"
)

type Device struct {
	Mac         string       `gorm:"primarykey"`
	DeviceID    uint64       `gorm:"not null;uniqueIndex"`
	Collections []Collection `gorm:"foreignKey:DeviceID;references:DeviceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime"`
}
