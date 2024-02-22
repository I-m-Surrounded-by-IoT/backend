package collection

import (
	"context"
	"fmt"
	"time"

	collection "github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/collection/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CollectionService struct {
	db *dbUtils
	collection.UnimplementedCollectionServer
}

func NewCollectionDatabase(dc *conf.DatabaseServerConfig, cc *conf.CollectionConfig) *CollectionService {
	d, err := dbdial.Dial(context.Background(), dc)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	if dc.AutoMigrate {
		log.Infof("auto migrate database...")
		err = d.AutoMigrate(
			new(model.CollectionRecord),
		)
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
		}
	}

	db := &CollectionService{
		db: newDBUtils(d),
	}
	return db
}

func (s *CollectionService) CreateCollectionRecord(ctx context.Context, req *collection.CollectionRecord) (*collection.Empty, error) {
	err := s.db.CreateCollectionRecord(proto2Record(req))
	if err != nil {
		return nil, err
	}
	return &collection.Empty{}, nil
}

func proto2Record(record *collection.CollectionRecord) *model.CollectionRecord {
	return &model.CollectionRecord{
		DeviceID:    record.DeviceId,
		CreatedAt:   time.UnixMilli(record.CreatedAt),
		Timestamp:   time.UnixMilli(record.Timestamp),
		GeoPoint:    model.GeoPoint{Lat: record.GeoPoint.Lat, Lng: record.GeoPoint.Lng},
		Temperature: record.Temperature,
	}
}

func record2Proto(record *model.CollectionRecord) *collection.CollectionRecord {
	return &collection.CollectionRecord{
		DeviceId:  record.DeviceID,
		Timestamp: record.Timestamp.UnixMilli(),
		GeoPoint: &collection.GeoPoint{
			Lat: record.GeoPoint.Lat,
			Lng: record.GeoPoint.Lng,
		},
		Temperature: record.Temperature,
	}
}

func records2Proto(records []*model.CollectionRecord) []*collection.CollectionRecord {
	resp := make([]*collection.CollectionRecord, len(records))
	for i, r := range records {
		resp[i] = record2Proto(r)
	}
	return resp
}

func (s *CollectionService) ListCollectionRecord(ctx context.Context, req *collection.ListCollectionRecordReq) (*collection.ListCollectionRecordResp, error) {
	opts := []func(*gorm.DB) *gorm.DB{}

	if req.Before != 0 {
		opts = append(opts, utils.WithTimestampBefore(req.Before))
	}
	if req.After != 0 {
		opts = append(opts, utils.WithTimestampAfter(req.After))
	}
	if req.DeviceId != 0 {
		opts = append(opts, utils.WithDeviceIDEq(req.DeviceId))
	}

	count, err := s.db.CountCollectionRecord(opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, model.WithPageAndPageSize(int(req.Page), int(req.Size)))
	switch req.Order {
	case collection.CollectionRecordOrder_CREATED_AT:
		opts = append(opts, utils.WithOrder(fmt.Sprintf("created_at %s", req.Sort)))
	default: // collection.CollectionRecordOrder_TIMESTAMP
		opts = append(opts, utils.WithOrder(fmt.Sprintf("timestamp %s", req.Sort)))
	}

	c, err := s.db.ListCollectionRecord(opts...)
	if err != nil {
		return nil, err
	}

	return &collection.ListCollectionRecordResp{
		Records: records2Proto(c),
		Total:   count,
	}, nil
}
