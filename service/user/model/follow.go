package model

import "time"

type FollowDevice struct {
	ID        uint64 `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string `gorm:"not null;type:char(32);uniqueIndex:idx_user_device,type:hash" json:"userId"`
	DeviceID  uint64 `gorm:"not null;uniqueIndex:idx_user_device" json:"deviceId"`
}
