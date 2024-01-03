package model

import (
	"database/sql/driver"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Collection struct {
	DeviceID  uint64    `gorm:"primarykey"`
	Timestamp time.Time `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	GeoPoint  GeoPoint  `gorm:"not null;type:geometry"`
}

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (p *GeoPoint) String() string {
	return fmt.Sprintf("POINT(%v %v)", p.Lng, p.Lat)
}

func (p GeoPoint) Value() (driver.Value, error) {
	return p.String(), nil
}

var pointRegex = regexp.MustCompile(`POINT\(([-0-9.]+) ([-0-9.]+)\)`)

func (p *GeoPoint) Scan(v any) error {
	if v == nil {
		return nil
	}
	switch v := v.(type) {
	case []byte:
		matches := pointRegex.FindSubmatch(v)
		if len(matches) != 3 {
			return fmt.Errorf("failed to unmarshal GeoPoint value: %#v", v)
		}
		f, err := strconv.ParseFloat(string(matches[1]), 64)
		if err != nil {
			return err
		}
		p.Lng = f
		f, err = strconv.ParseFloat(string(matches[2]), 64)
		if err != nil {
			return err
		}
		p.Lat = f
	case string:
		matches := pointRegex.FindSubmatch([]byte(v))
		if len(matches) != 3 {
			return fmt.Errorf("failed to unmarshal GeoPoint value: %#v", v)
		}
		f, err := strconv.ParseFloat(string(matches[1]), 64)
		if err != nil {
			return err
		}
		p.Lng = f
		f, err = strconv.ParseFloat(string(matches[2]), 64)
		if err != nil {
			return err
		}
		p.Lat = f
	default:
		return fmt.Errorf("failed to unmarshal GeoPoint value: %#v", v)
	}
	return nil
}

func WithDeviceID(deviceID uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("device_id = ?", deviceID)
	}
}

func WithStartTime(startTime time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("timestamp >= ?", startTime)
	}
}

func WithEndTime(endTime time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("timestamp <= ?", endTime)
	}
}

func WithPageAndPageSize(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}

func WithOrder(order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}
