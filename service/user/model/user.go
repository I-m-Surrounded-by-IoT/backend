package model

import (
	"database/sql"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"gorm.io/gorm"
)

type User struct {
	ID              string `gorm:"primaryKey;type:char(32);index:,type:hash" json:"id"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Username        string          `gorm:"not null;uniqueIndex;type:varchar(32)"`
	HashedPassword  []byte          `gorm:"not null"`
	Role            user.Role       `gorm:"not null;default:0"`
	Status          user.Status     `gorm:"not null;default:0"`
	Email           string          `gorm:"type:varchar(64)"`
	FollowDevices   []*FollowDevice `gorm:"foreignKey:UserID;references:ID"`
	FollowAllDevice sql.NullBool    `gorm:"default:false"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = utils.SortUUID()
	}
	return nil
}
