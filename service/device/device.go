package device

import (
	"context"

	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/device/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	log "github.com/sirupsen/logrus"
)

type DeviceService struct {
	db *dbUtils
	device.UnimplementedDeviceServer
}

func NewDeviceService(dc *conf.DatabaseServerConfig, deviceConfig *conf.DeviceConfig) *DeviceService {
	d, err := dbdial.Dial(context.Background(), dc)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	if dc.AutoMigrate {
		log.Infof("auto migrate database...")
		err = d.AutoMigrate(
			new(model.Device),
		)
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
		}
	}

	db := &DeviceService{
		db: newDBUtils(d),
	}
	return db
}

func (s *DeviceService) GetDevice(ctx context.Context, req *device.GetDeviceReq) (*device.DeviceRecord, error) {
	d, err := s.db.GetDevice(req.Id)
	if err != nil {
		return nil, err
	}
	return &device.DeviceRecord{
		Id:        d.ID,
		CreatedAt: d.CreatedAt.UnixMicro(),
		UpdatedAt: d.UpdatedAt.UnixMicro(),
		Mac:       d.Mac,
		Comment:   d.Comment,
	}, nil
}

func (s *DeviceService) GetDeviceByMac(ctx context.Context, req *device.GetDeviceByMacReq) (*device.DeviceRecord, error) {
	d, err := s.db.GetDeviceWithMac(req.Mac)
	if err != nil {
		return nil, err
	}
	return Device2Proto(d), nil
}

func (s *DeviceService) CreateDevice(ctx context.Context, req *device.CreateDeviceReq) (*device.DeviceRecord, error) {
	d := &model.Device{
		Mac: req.Mac,
	}
	err := s.db.CreateDevice(d)
	if err != nil {
		return nil, err
	}
	return Device2Proto(d), nil
}

func (s *DeviceService) GetOrCreateDevice(ctx context.Context, req *device.GetOrCreateDeviceReq) (*device.DeviceRecord, error) {
	d := &model.Device{
		Mac: req.Mac,
	}
	err := s.db.FirstOrCreateDevice(d)
	if err != nil {
		return nil, err
	}
	return Device2Proto(d), nil
}
