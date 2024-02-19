package database

import (
	"github.com/I-m-Surrounded-by-IoT/backend/api/database"
	"github.com/I-m-Surrounded-by-IoT/backend/service/database/model"
	"gorm.io/gorm"
)

type dbUtils struct {
	db *gorm.DB
}

func newDBUtils(db *gorm.DB) *dbUtils {
	return &dbUtils{db: db}
}

func (u *dbUtils) ListCollectionInfo(deviceID uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]*model.Collection, error) {
	var collections []*model.Collection
	err := u.db.Scopes(scopes...).Where("device_id = ?", deviceID).Find(&collections).Error
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (u *dbUtils) CreateCollection(collection *model.Collection) error {
	return u.db.Create(collection).Error
}

func (u *dbUtils) GetDevice(id uint64) (*model.Device, error) {
	var device model.Device
	err := u.db.First(&device, id).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (u *dbUtils) CreateDevice(device *model.Device) error {
	return u.db.Create(device).Error
}

func (u *dbUtils) FirstOrCreateDevice(device *model.Device) error {
	return u.db.Where("mac = ?", device.Mac).Attrs(device).FirstOrCreate(device).Error
}

func (u *dbUtils) GetDeviceWithMac(mac string) (*model.Device, error) {
	var device model.Device
	err := u.db.Where("mac = ?", mac).First(&device).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func Device2Proto(device *model.Device) *database.Device {
	return &database.Device{
		DeviceId:  device.DeviceID,
		Mac:       device.Mac,
		CreatedAt: device.CreatedAt.UnixMicro(),
		UpdatedAt: device.UpdatedAt.UnixMicro(),
	}
}
