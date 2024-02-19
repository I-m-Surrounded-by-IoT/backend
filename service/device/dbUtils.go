package device

import (
	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/service/device/model"
	"gorm.io/gorm"
)

type dbUtils struct {
	db *gorm.DB
}

func newDBUtils(db *gorm.DB) *dbUtils {
	return &dbUtils{db: db}
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

func Device2Proto(d *model.Device) *device.DeviceRecord {
	return &device.DeviceRecord{
		Id:        d.DeviceID,
		Mac:       d.Mac,
		CreatedAt: d.CreatedAt.UnixMicro(),
		UpdatedAt: d.UpdatedAt.UnixMicro(),
	}
}
