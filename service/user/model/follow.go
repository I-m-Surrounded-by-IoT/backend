package model

import "time"

type FollowDevice struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeviceID  uint64 `gorm:"not null;uniqueIndex:idx_device_user" json:"deviceId"`
	UserID    string `gorm:"not null;type:char(32);uniqueIndex:idx_device_user" json:"userId"`
}
