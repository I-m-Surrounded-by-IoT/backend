package model

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CollectionRecord struct {
	DeviceID    uint64    `gorm:"primarykey"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	Timestamp   time.Time `gorm:"primarykey"`
	GeoPoint    GeoPoint  `gorm:"not null;type:geography(POINT, 4326);index:,type:gist"`
	Temperature float32   `gorm:"not null"`
}

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (p *GeoPoint) String() string {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", p.Lon, p.Lat)
}

func (p GeoPoint) Value() (driver.Value, error) {
	return p.String(), nil
}

func (p *GeoPoint) Scan(v any) (err error) {
	var b []byte
	switch v := v.(type) {
	case string:
		b, err = hex.DecodeString(v)
		if err != nil {
			return err
		}
	case []byte:
		b = v
	default:
		return fmt.Errorf("invalid type %T", v)
	}

	r := bytes.NewReader(b)
	var wkbByteOrder uint8
	err = binary.Read(r, binary.LittleEndian, &wkbByteOrder)
	if err != nil {
		return err
	}

	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return fmt.Errorf("invalid byte order %d", wkbByteOrder)
	}

	var wkbGeometryType uint64
	err = binary.Read(r, byteOrder, &wkbGeometryType)
	if err != nil {
		return err
	}

	return binary.Read(r, byteOrder, p)
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

func WithOrder(order string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(order)
	}
}
