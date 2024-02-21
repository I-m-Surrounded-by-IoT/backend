package device

import (
	"context"
	"fmt"

	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/device/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/rcache"
	redsync "github.com/go-redsync/redsync/v4"
	goredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DeviceService struct {
	db      *dbUtils
	drcache *DeviceRcache
	device.UnimplementedDeviceServer
}

func NewDeviceService(dc *conf.DatabaseServerConfig, deviceConfig *conf.DeviceConfig, rc *conf.RedisConfig) *DeviceService {
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

	rdb := redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Username: rc.Username,
		Password: rc.Password,
		DB:       int(rc.Db),
	})
	db := newDBUtils(d)
	rsync := redsync.New(goredis.NewPool(rdb))

	return &DeviceService{
		db:      db,
		drcache: NewDeviceRcache(rcache.NewRcacheWithRsync(rdb, rsync), db),
	}
}

func (s *DeviceService) GetDeviceInfo(ctx context.Context, req *device.GetDeviceInfoReq) (*device.DeviceInfo, error) {
	return s.drcache.GetDeviceInfo(ctx, req.Id, req.Fields...)
}

func (s *DeviceService) GetDeviceInfoByMac(ctx context.Context, req *device.GetDeviceInfoByMacReq) (*device.DeviceInfo, error) {
	return s.drcache.GetDeviceInfoByMac(ctx, req.Mac, req.Fields...)
}

func (s *DeviceService) GetDeviceID(ctx context.Context, req *device.GetDeviceIDReq) (*device.DeviceInfo, error) {
	id, err := s.drcache.GetDeviceID(ctx, req.Mac)
	if err != nil {
		return nil, err
	}
	return &device.DeviceInfo{
		Id: id,
	}, nil
}

func (s *DeviceService) CreateDevice(ctx context.Context, req *device.CreateDeviceReq) (*device.DeviceInfo, error) {
	d := &model.Device{
		Mac: req.Mac,
	}
	err := s.db.CreateDevice(ctx, d)
	if err != nil {
		return nil, err
	}
	return device2Proto(d), nil
}

func (s *DeviceService) DeleteDevice(ctx context.Context, req *device.DeleteDeviceReq) (*device.Empty, error) {
	_, err := s.drcache.DelDevice(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &device.Empty{}, nil
}

func (s *DeviceService) GetOrCreateDevice(ctx context.Context, req *device.GetOrCreateDeviceReq) (*device.DeviceInfo, error) {
	d := &model.Device{
		Mac: req.Mac,
	}
	err := s.db.FirstOrCreateDevice(ctx, d)
	if err != nil {
		return nil, err
	}
	return device2Proto(d), nil
}

func (s *DeviceService) ListDeletedDeviceInfo(ctx context.Context, req *device.ListDeviceReq) (*device.ListDeviceResp, error) {
	opts := []func(*gorm.DB) *gorm.DB{}
	if req.Id != 0 {
		opts = append(opts, model.WithIDEq(req.Id))
	}
	if req.Mac != "" {
		opts = append(opts, model.WithMacEq(req.Mac))
	}
	count, err := s.db.CountDeletedDevice(ctx, opts...)
	if err != nil {
		return nil, err
	}
	opts = append(opts, utils.WithPageAndPageSize(int(req.Page), int(req.Size)))
	switch req.Order {
	case device.ListDeviceOrder_ID:
		opts = append(opts, model.WithOrder(fmt.Sprintf("id %s", req.Sort)))
	case device.ListDeviceOrder_MAC:
		opts = append(opts, model.WithOrder(fmt.Sprintf("mac %s", req.Sort)))
	case device.ListDeviceOrder_CREATED_AT:
		opts = append(opts, model.WithOrder(fmt.Sprintf("created_at %s", req.Sort)))
	case device.ListDeviceOrder_UPDATED_AT:
		opts = append(opts, model.WithOrder(fmt.Sprintf("updated_at %s", req.Sort)))
	}
	if len(req.Fields) != 0 {
		opts = append(opts, model.WithFields(req.Fields...))
	}
	d, err := s.db.ListDeletedDevice(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &device.ListDeviceResp{
		Devices: devices2Proto(d),
		Total:   int32(count),
	}, nil
}

func (s *DeviceService) ListDevice(ctx context.Context, req *device.ListDeviceReq) (*device.ListDeviceResp, error) {
	opts := []func(*gorm.DB) *gorm.DB{}
	if req.Id != 0 {
		opts = append(opts, model.WithIDEq(req.Id))
	}
	if req.Mac != "" {
		opts = append(opts, model.WithMacEq(req.Mac))
	}
	count, err := s.db.CountDevice(ctx, opts...)
	if err != nil {
		return nil, err
	}
	opts = append(opts, utils.WithPageAndPageSize(int(req.Page), int(req.Size)))
	switch req.Order {
	case device.ListDeviceOrder_ID:
		opts = append(opts, model.WithOrder(fmt.Sprintf("id %s", req.Sort)))
	case device.ListDeviceOrder_MAC:
		opts = append(opts, model.WithOrder(fmt.Sprintf("mac %s", req.Sort)))
	case device.ListDeviceOrder_CREATED_AT:
		opts = append(opts, model.WithOrder(fmt.Sprintf("created_at %s", req.Sort)))
	case device.ListDeviceOrder_UPDATED_AT:
		opts = append(opts, model.WithOrder(fmt.Sprintf("updated_at %s", req.Sort)))
	}
	if len(req.Fields) != 0 {
		opts = append(opts, model.WithFields(req.Fields...))
	}
	d, err := s.db.ListDevice(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &device.ListDeviceResp{
		Devices: devices2Proto(d),
		Total:   int32(count),
	}, nil
}

func (s *DeviceService) UpdateDeviceLastSeen(ctx context.Context, req *device.UpdateDeviceLastSeenReq) (*device.Empty, error) {
	return &device.Empty{}, s.drcache.UpdateDeviceLastSeen(ctx, req.Id, req.LastSeen)
}

func (s *DeviceService) GetDeviceLastSeen(ctx context.Context, req *device.GetDeviceLastSeenReq) (*device.GetDeviceLastSeenResp, error) {
	lastSeen, err := s.drcache.GetDeviceLastSeen(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &device.GetDeviceLastSeenResp{
		LastSeen: lastSeen,
	}, nil
}
