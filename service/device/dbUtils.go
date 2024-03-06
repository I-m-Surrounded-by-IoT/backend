package device

import (
	"context"

	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/service/device/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type dbUtils struct {
	db *gorm.DB
}

func newDBUtils(db *gorm.DB) *dbUtils {
	return &dbUtils{db: db}
}

func (u *dbUtils) GetDeviceInfo(ctx context.Context, id uint64, fields ...string) (*model.Device, error) {
	var device model.Device
	err := u.db.WithContext(ctx).Select(fields).First(&device, id).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (u *dbUtils) CreateDevice(ctx context.Context, device *model.Device) error {
	return u.db.WithContext(ctx).Create(device).Error
}

func (u *dbUtils) Transaction(fn func(db *dbUtils) error) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		return fn(newDBUtils(tx))
	})
}

func (u *dbUtils) DelDevice(ctx context.Context, id uint64, fields ...string) (*model.Device, error) {
	device := &model.Device{}
	err := u.db.WithContext(ctx).Clauses(clause.Returning{}).Select(fields).Delete(device, id).Error
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (u *dbUtils) UndelDevice(ctx context.Context, id uint64, fields ...string) (*model.Device, error) {
	device := &model.Device{}
	err := u.db.WithContext(ctx).Clauses(clause.Returning{}).Unscoped().Model(device).Select(fields).Where("id = ?", id).Update("deleted_at", nil).Error
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (u *dbUtils) ListDeletedDevice(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]*model.Device, error) {
	return u.ListDevice(ctx,
		append(scopes,
			func(d *gorm.DB) *gorm.DB {
				return d.Unscoped()
			},
		)...,
	)
}

func (u *dbUtils) ListDevice(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) ([]*model.Device, error) {
	var devices []*model.Device
	err := u.db.WithContext(ctx).Scopes(scopes...).Find(&devices).Error
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (u *dbUtils) CountDeletedDevice(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	return u.CountDevice(ctx,
		append(scopes,
			func(d *gorm.DB) *gorm.DB {
				return d.Unscoped()
			},
		)...,
	)
}

func (u *dbUtils) CountDevice(ctx context.Context, scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := u.db.WithContext(ctx).Scopes(scopes...).Model(&model.Device{}).Count(&count).Error
	return count, err
}

func (u *dbUtils) FirstOrCreateDevice(ctx context.Context, device *model.Device) error {
	return u.db.WithContext(ctx).Where("mac = ?", device.Mac).Attrs(device).FirstOrCreate(device).Error
}

func (u *dbUtils) GetDeviceInfoWithMac(ctx context.Context, mac string, fields ...string) (*model.Device, error) {
	var device model.Device
	err := u.db.WithContext(ctx).Select(fields).Where("mac = ?", mac).First(&device).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func device2Proto(d *model.Device) *device.DeviceInfo {
	return &device.DeviceInfo{
		Id:        d.ID,
		Mac:       d.Mac,
		CreatedAt: d.CreatedAt.UnixMilli(),
		UpdatedAt: d.UpdatedAt.UnixMilli(),
		Comment:   d.Comment,
	}
}

func devices2Proto(ds []*model.Device) []*device.DeviceInfo {
	var res []*device.DeviceInfo = make([]*device.DeviceInfo, len(ds))
	for i, d := range ds {
		res[i] = device2Proto(d)
	}
	return res
}
