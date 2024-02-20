package collection

import (
	"github.com/I-m-Surrounded-by-IoT/backend/service/collection/model"
	"gorm.io/gorm"
)

type dbUtils struct {
	db *gorm.DB
}

func newDBUtils(db *gorm.DB) *dbUtils {
	return &dbUtils{db: db}
}

func (u *dbUtils) ListCollectionInfo(deviceID uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]*model.CollectionRecord, error) {
	var collections []*model.CollectionRecord
	err := u.db.Scopes(scopes...).Where("device_id = ?", deviceID).Find(&collections).Error
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (u *dbUtils) CreateCollection(collection *model.CollectionRecord) error {
	return u.db.Create(collection).Error
}
