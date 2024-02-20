package model

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"gorm.io/gorm"
)

type User struct {
	ID             string `gorm:"primaryKey;type:char(32);index:,type:hash" json:"id"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Username       string      `gorm:"not null;uniqueIndex;type:varchar(32)"`
	HashedPassword []byte      `gorm:"not null"`
	Role           user.Role   `gorm:"not null;default:0"`
	Status         user.Status `gorm:"not null;default:1"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var existingUser User
	err := tx.Where("username = ?", u.Username).First(&existingUser).Error
	if err == nil {
		u.Username = fmt.Sprintf("%s#%d", u.Username, rand.Intn(9999))
	}
	if u.ID == "" {
		u.ID = utils.SortUUID()
	}
	return nil
}
