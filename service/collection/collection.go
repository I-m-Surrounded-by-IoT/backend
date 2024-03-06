package collection

import (
	"context"
	"fmt"
	"time"

	collectionApi "github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/I-m-Surrounded-by-IoT/backend/service/collection/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type CollectionService struct {
	kc sarama.Client
	db *dbUtils
	collectionApi.UnimplementedCollectionServer
}

func NewCollectionDatabase(dc *conf.DatabaseServerConfig, cc *conf.CollectionConfig, kc sarama.Client) *CollectionService {
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
		kc: kc,
	}
	return db
}

func (s *CollectionService) CreateCollectionRecord(ctx context.Context, req *collectionApi.CollectionRecord) (*collectionApi.Empty, error) {
	err := s.db.CreateCollectionRecord(proto2Record(req))
	if err != nil {
		return nil, err
	}
	return &collectionApi.Empty{}, nil
}

func proto2Record(record *collectionApi.CollectionRecord) *model.CollectionRecord {
	return &model.CollectionRecord{
		DeviceID:    record.DeviceId,
		CreatedAt:   time.UnixMilli(record.CreatedAt),
		Timestamp:   time.UnixMilli(record.Data.Timestamp),
		GeoPoint:    model.GeoPoint{Lat: record.Data.GeoPoint.Lat, Lon: record.Data.GeoPoint.Lon},
		Temperature: record.Data.Temperature,
	}
}

func record2Proto(record *model.CollectionRecord) *collectionApi.CollectionRecord {
	return &collectionApi.CollectionRecord{
		DeviceId: record.DeviceID,
		Data: &collectionApi.CollectionData{
			Timestamp: record.Timestamp.UnixMilli(),
			GeoPoint: &collectionApi.GeoPoint{
				Lat: record.GeoPoint.Lat,
				Lon: record.GeoPoint.Lon,
			},
			Temperature: record.Temperature,
		},
	}
}

func records2Proto(records []*model.CollectionRecord) []*collectionApi.CollectionRecord {
	resp := make([]*collectionApi.CollectionRecord, len(records))
	for i, r := range records {
		resp[i] = record2Proto(r)
	}
	return resp
}

func (s *CollectionService) ListCollectionRecord(ctx context.Context, req *collectionApi.ListCollectionRecordReq) (*collectionApi.ListCollectionRecordResp, error) {
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

	opts = append(opts, utils.WithPageAndPageSize(int(req.Page), int(req.Size)))
	switch req.Order {
	case collectionApi.CollectionRecordOrder_CREATED_AT:
		opts = append(opts, utils.WithOrder(fmt.Sprintf("created_at %s", req.Sort)))
	default: // collection.CollectionRecordOrder_TIMESTAMP
		opts = append(opts, utils.WithOrder(fmt.Sprintf("timestamp %s", req.Sort)))
	}

	c, err := s.db.ListCollectionRecord(opts...)
	if err != nil {
		return nil, err
	}

	return &collectionApi.ListCollectionRecordResp{
		Records: records2Proto(c),
		Total:   count,
	}, nil
}

func (s *CollectionService) GetDeviceStreamReport(req *collectionApi.GetDeviceStreamReportReq, resp collectionApi.Collection_GetDeviceStreamReportServer) error {
	cg, err := sarama.NewConsumerFromClient(s.kc)
	if err != nil {
		return err
	}
	defer cg.Close()
	var topic string
	if req.Id == 0 {
		topic = service.KafkaTopicDeviceReport
	} else {
		topic = fmt.Sprintf("%s-%d", service.KafkaTopicDeviceReport, req.Id)
	}
	ps, err := cg.Partitions(topic)
	if err != nil {
		return err
	}

	if len(ps) == 0 {
		log.Errorf("no partition found")
		return nil
	}

	wg, ctx := errgroup.WithContext(resp.Context())
	var ch = make(chan *collectionApi.CollectionData)
	for _, p := range ps {
		c, err := cg.ConsumePartition(topic, p, sarama.OffsetNewest)
		if err != nil {
			return err
		}
		wg.Go(func() error {
			defer c.Close()
			for {
				select {
				case <-ctx.Done():
					return nil
				case msg := <-c.Messages():
					data, err := service.KafkaTopicDeviceReportUnmarshal(msg.Value)
					if err != nil {
						log.Errorf("failed to unmarshal device report (%s): %v", msg.Value, err)
						return err
					}
					select {
					case <-ctx.Done():
						return ctx.Err()
					case ch <- data:
					}
				}
			}
		})
	}

	defer func() {
		_ = wg.Wait()
		close(ch)
	}()

	for v := range ch {
		err = resp.Send(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *CollectionService) GetDeviceStreamEvent(req *collectionApi.GetDeviceStreamEventReq, resp collectionApi.Collection_GetDeviceStreamEventServer) error {
	cg, err := sarama.NewConsumerFromClient(s.kc)
	if err != nil {
		log.Errorf("failed to create consumer group: %v", err)
		return err
	}
	defer cg.Close()
	var topic string
	if req.Id == 0 {
		topic = service.KafkaTopicDeviceLog
	} else {
		topic = fmt.Sprintf("%s-%d", service.KafkaTopicDeviceLog, req.Id)
	}
	ps, err := cg.Partitions(topic)
	if err != nil {
		log.Errorf("failed to get partitions: %v", err)
		return err
	}

	if len(ps) == 0 {
		log.Errorf("no partition found")
		return nil
	}

	wg, ctx := errgroup.WithContext(resp.Context())
	var ch = make(chan *collectionApi.GetDeviceStreamEventResp)
	for _, p := range ps {
		c, err := cg.ConsumePartition(topic, p, sarama.OffsetNewest)
		if err != nil {
			return err
		}
		wg.Go(func() error {
			defer c.Close()
			for {
				select {
				case <-ctx.Done():
					return nil
				case msg := <-c.Messages():
					data, err := service.KafkaTopicDeviceLogUnmarshal(msg.Value)
					if err != nil {
						log.Errorf("failed to unmarshal device log (%s): %v", msg.Value, err)
						return err
					}
					select {
					case ch <- &collectionApi.GetDeviceStreamEventResp{
						Topic:     data.Topic,
						Message:   data.Message,
						Timestamp: data.Timestamp,
					}:
					case <-ctx.Done():
						return ctx.Err()
					}
				}
			}
		})
	}

	defer func() {
		_ = wg.Wait()
		close(ch)
	}()

	for v := range ch {
		err = resp.Send(v)
		if err != nil {
			log.Errorf("failed to send event: %v", err)
			return err
		}
	}
	return nil
}
