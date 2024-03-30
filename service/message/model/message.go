package model

import (
	"database/sql"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/message"
)

type Message struct {
	ID          uint64       `gorm:"primaryKey"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdateAt    time.Time    `gorm:"autoUpdateTime"`
	Unread      sql.NullBool `gorm:"default:true"`
	UserID      string       `gorm:"not null;index"`
	Timestamp   time.Time    `gorm:"not null;index"`
	MessageType message.MessageType
	Title       string `gorm:"type:VARCHAR(255)"`
	Content     string `gorm:"type:TEXT"`
}
