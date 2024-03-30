package model

import (
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"gorm.io/gorm"
)

func WithUnread() utils.Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("unread = ?", true)
	}
}
