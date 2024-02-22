package model

import (
	"time"

	"gorm.io/gorm"
)

func WithTimestampBefore(t int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("timestamp < ?", time.UnixMilli(t))
	}
}

func WithTimestampAfter(t int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("timestamp > ?", time.UnixMilli(t))
	}
}

func WithDeviceIDEq(id uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("device_id = ?", id)
	}
}

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

func WithOrder(order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}
