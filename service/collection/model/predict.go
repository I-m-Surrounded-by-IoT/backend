package model

import (
	"time"

	"gorm.io/gorm"
)

type PredictAndGuess struct {
	ID                 uint              `gorm:"primarykey"`
	CollectionRecordID uint              `gorm:"uniqueIndex;not null"`
	DeviceID           uint64            `gorm:"not null"`
	CreatedAt          time.Time         `gorm:"autoCreateTime"`
	UpdatedAt          time.Time         `gorm:"autoUpdateTime"`
	Level              int64             `gorm:"default:-1"`
	Predicts           []*CollectionData `gorm:"serializer:fastjson"`
	Levles             []int64           `gorm:"serializer:fastjson"`
}

// 更新CollectionRecord的UpdatedAt
func (p *PredictAndGuess) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&CollectionRecord{}).Where("id = ?", p.CollectionRecordID).Update("updated_at", p.CreatedAt).Error
}

func (p *PredictAndGuess) AfterUpdate(tx *gorm.DB) (err error) {
	return tx.Model(&CollectionRecord{}).Where("id = ?", p.CollectionRecordID).Update("updated_at", p.UpdatedAt).Error
}
