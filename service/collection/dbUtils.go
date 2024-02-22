package collection

import (
	"github.com/I-m-Surrounded-by-IoT/backend/service/collection/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"gorm.io/gorm"
)

type dbUtils struct {
	db *gorm.DB
}

func newDBUtils(db *gorm.DB) *dbUtils {
	return &dbUtils{db: db}
}

func (u *dbUtils) GetCollectionRecord(deviceID uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]*model.CollectionRecord, error) {
	return u.ListCollectionRecord(append(scopes, utils.WithDeviceIDEq(deviceID))...)
}

func (u *dbUtils) ListCollectionRecord(scopes ...func(*gorm.DB) *gorm.DB) ([]*model.CollectionRecord, error) {
	var collections []*model.CollectionRecord
	err := u.db.Scopes(scopes...).Find(&collections).Error
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (u *dbUtils) CountCollectionRecord(scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := u.db.Model(&model.CollectionRecord{}).Scopes(scopes...).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *dbUtils) CreateCollectionRecord(collection *model.CollectionRecord) error {
	return u.db.Create(collection).Error
}
