package collection_database

import (
	"context"
	"time"

	collection_database "github.com/I-m-Surrounded-by-IoT/backend/api/collection-database"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/collection-database/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CollectionDatabaseService struct {
	db *dbUtils
	collection_database.UnimplementedCollectionDatabaseServer
}

func NewCollectionDatabase(c *conf.DatabaseConfig) *CollectionDatabaseService {
	d, err := dbdial.NewDatabase(context.Background(), c)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	if c.AutoMigrate {
		log.Infof("auto migrate database...")
		err = d.AutoMigrate(
			new(model.CollectionRecord),
		)
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
		}
	}

	db := &CollectionDatabaseService{
		db: newDBUtils(d),
	}
	return db
}

func (s *CollectionDatabaseService) CreateCollectionRecord(ctx context.Context, req *collection_database.CollectionRecord) (*collection_database.Empty, error) {
	err := s.db.CreateCollection(&model.CollectionRecord{
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
	return &collection_database.Empty{}, nil
}

func (s *CollectionDatabaseService) ListCollectionRecord(ctx context.Context, req *collection_database.ListCollectionRecordReq) (*collection_database.ListCollectionRecordResp, error) {
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
	resp := collection_database.ListCollectionRecordResp{
		CollectionInfos: make([]*collection_database.ListCollectionRecordResp_CollectionRecord, len(c)),
	}
	for i, info := range c {
		resp.CollectionInfos[i] = &collection_database.ListCollectionRecordResp_CollectionRecord{
			CreatedAt: info.CreatedAt.UnixMicro(),
			DeviceId:  info.DeviceID,
			Timestamp: info.Timestamp.UnixMicro(),
			GeoPoint: &collection_database.GeoPoint{
				Lat: info.GeoPoint.Lat,
				Lng: info.GeoPoint.Lng,
			},
		}
	}
	return &resp, nil
}
