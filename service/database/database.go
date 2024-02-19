package database

import (
	"context"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/database"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/database/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type DatabaseService struct {
	db *dbUtils
	database.UnimplementedDatabaseServer
}

func NewDatabaseService(c *conf.DatabaseConfig) *DatabaseService {
	d, err := dbdial.NewDatabase(context.Background(), c)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	if c.AutoMigrate {
		log.Infof("auto migrate database...")
		err = d.AutoMigrate(
			new(model.Device),
			new(model.Collection),
		)
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
		}
	}

	db := &DatabaseService{
		db: newDBUtils(d),
	}
	return db
}

func (s *DatabaseService) CreateCollectionInfo(ctx context.Context, req *database.CollectionInfo) (*database.Empty, error) {
	err := s.db.CreateCollection(&model.Collection{
		DeviceID:  req.DeviceId,
		Timestamp: time.UnixMicro(int64(req.Timestamp)),
		GeoPoint: model.GeoPoint{
			Lat: req.GeoPoint.Lat,
			Lng: req.GeoPoint.Lng,
		},
		Temperature: req.Temperature,
	})
	if err != nil {
		return nil, err
	}
	return &database.Empty{}, nil
}

func (s *DatabaseService) ListCollectionInfo(ctx context.Context, req *database.ListCollectionInfoReq) (*database.ListCollectionInfoResp, error) {
	scopes := make([]func(*gorm.DB) *gorm.DB, 0)

	if req.StartTimestamp != 0 {

		scopes = append(scopes, model.WithStartTime(time.UnixMicro(int64(req.StartTimestamp))))
	}
	if req.EndTimestamp != 0 {
		scopes = append(scopes, model.WithEndTime(time.UnixMicro(int64(req.EndTimestamp))))
	}
	if req.Page == 0 {
		req.Page = 1
	}
	if req.PageSize == 0 {
		req.PageSize = 10
	}
	scopes = append(scopes, model.WithPageAndPageSize(int(req.Page), int(req.PageSize)))
	c, err := s.db.ListCollectionInfo(req.DeviceId, scopes...)
	if err != nil {
		return nil, err
	}
	resp := database.ListCollectionInfoResp{
		CollectionInfos: make([]*database.ListCollectionInfoResp_CollectionInfo, len(c)),
	}
	for i, info := range c {
		resp.CollectionInfos[i] = &database.ListCollectionInfoResp_CollectionInfo{
			CreatedAt: info.CreatedAt.UnixMicro(),
			DeviceId:  info.DeviceID,
			Timestamp: info.Timestamp.UnixMicro(),
			GeoPoint: &database.GeoPoint{
				Lat: info.GeoPoint.Lat,
				Lng: info.GeoPoint.Lng,
			},
		}
	}
	return &resp, nil
}

func (s *DatabaseService) GetDevice(ctx context.Context, req *database.GetDeviceReq) (*database.Device, error) {
	d, err := s.db.GetDevice(req.DeviceId)
	if err != nil {
		return nil, err
	}
	return &database.Device{
		DeviceId:  d.DeviceID,
		CreatedAt: d.CreatedAt.UnixMicro(),
		UpdatedAt: d.UpdatedAt.UnixMicro(),
		Mac:       d.Mac,
	}, nil
}

func (s *DatabaseService) GetDeviceWithMac(ctx context.Context, req *database.GetDeviceWithMacReq) (*database.Device, error) {
	d, err := s.db.GetDeviceWithMac(req.Mac)
	if err != nil {
		return nil, err
	}
	return Device2Proto(d), nil
}

func (s *DatabaseService) CreateDevice(ctx context.Context, req *database.CreateDeviceReq) (*database.Device, error) {
	d := &model.Device{
		Mac: req.Mac,
	}
	err := s.db.CreateDevice(d)
	if err != nil {
		return nil, err
	}
	return Device2Proto(d), nil
}

func (s *DatabaseService) FirstOrCreateDevice(ctx context.Context, req *database.CreateDeviceReq) (*database.Device, error) {
	d := &model.Device{
		Mac: req.Mac,
	}
	err := s.db.FirstOrCreateDevice(d)
	if err != nil {
		return nil, err
	}
	return Device2Proto(d), nil
}
