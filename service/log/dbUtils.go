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
	err := u.db.Scopes(scopes...).Where("device_id = ?", deviceID).Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}
