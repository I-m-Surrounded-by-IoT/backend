package model

import (
	"gorm.io/gorm"
)

func WithLevelRange(min, max uint32) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("level >= ? AND level <= ?", min, max)
	}
}

func WithLevelFilter(filter string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, c := range filter {
			db = db.Where("level != ?", c)
		}
		return db
	}
}
