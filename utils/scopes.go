package utils

import (
	"time"

	"golang.org/x/exp/constraints"
	"gorm.io/gorm"
)

type Scope = func(*gorm.DB) *gorm.DB

func WithPageAndPageSize(page, pageSize int) Scope {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	} else if pageSize > 100 {
		pageSize = 100
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}

func WithTimestampBefore(t int64) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("timestamp < ?", time.UnixMilli(t))
	}
}

func WithTimestampAfter(t int64) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("timestamp > ?", time.UnixMilli(t))
	}
}

func WithOrder(order string) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}

func WithDeviceIDEq[T constraints.Integer](id T) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("device_id = ?", id)
	}
}

func WithDeviceIDRange[T constraints.Integer](start, end T) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("device_id >= ? AND device_id <= ?", start, end)
	}
}

func WithIDEq[T constraints.Ordered](id T) Scope {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id = ?", id)
	}
}

func WithIDRange[T constraints.Ordered](start, end T) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id >= ? AND id <= ?", start, end)
	}
}
