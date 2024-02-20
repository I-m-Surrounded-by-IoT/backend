package model

import (
	"time"

	"gorm.io/gorm"
)

type Device struct {
	ID        uint64         `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Mac       string         `gorm:"not null;index:,type:hash;type:char(17)"`
	Comment   string         `gorm:"type:varchar(255)"`
}
