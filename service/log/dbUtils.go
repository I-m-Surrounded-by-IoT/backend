package log

import (
	"github.com/I-m-Surrounded-by-IoT/backend/service/log/model"
	"gorm.io/gorm"
)

type dbUtils struct {
	db *gorm.DB
}

func newDBUtils(db *gorm.DB) *dbUtils {
	return &dbUtils{db: db}
}

func (u *dbUtils) CreateDeviceLog(log *model.DeviceLog) error {
	return u.db.Create(log).Error
}

func (u *dbUtils) ListDeviceLog(deviceID uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]*model.DeviceLog, error) {
	var logs []*model.DeviceLog
	err := u.db.Scopes(model.WithDeviceIDEq(deviceID)).Scopes(scopes...).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (u *dbUtils) CountDeviceLog(deviceID uint64, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := u.db.Model(&model.DeviceLog{}).Scopes(model.WithDeviceIDEq(deviceID)).Scopes(scopes...).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
