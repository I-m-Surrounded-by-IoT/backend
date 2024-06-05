package collection

import (
	"context"
	"fmt"
	"time"

	collectionApi "github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/api/waterquality"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	registryClient "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/service"
	"github.com/I-m-Surrounded-by-IoT/backend/service/collection/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/rcache"
	"github.com/IBM/sarama"
	"github.com/go-kratos/kratos/v2/registry"
	redsync "github.com/go-redsync/redsync/v4"
	goredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/zijiren233/gencontainer/set"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

type CollectionService struct {
	kc                 sarama.Client
	db                 *dbUtils
	ccache             *CollectionRcache
	waterQualityClient waterquality.WaterQualityServiceClient
	collectionApi.UnimplementedCollectionServer
}

func NewCollectionDatabase(dc *conf.DatabaseServerConfig, cc *conf.CollectionConfig, kc sarama.Client, rc *conf.RedisConfig, reg registry.Registrar) *CollectionService {
	etcd := reg.(*registryClient.EtcdRegistry)
	discoveryWaterQualityConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///water-quality",
		TimeOut:  "10s",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}
	waterQualityClient := waterquality.NewWaterQualityServiceClient(discoveryWaterQualityConn)

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

	rdb := redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Username: rc.Username,
		Password: rc.Password,
		DB:       int(rc.Db),
	})

	dbu := newDBUtils(d)

	db := &CollectionService{
		db: dbu,
		kc: kc,
		ccache: NewCollectionRcache(
			rcache.NewRcacheWithRsync(
				rdb,
				redsync.New(goredis.NewPool(rdb)),
			),
			dbu,
		),
		waterQualityClient: waterQualityClient,
	}
	return db
}

func collectionRecords2Qualities(records []*collectionApi.CollectionRecord) []*waterquality.Quality {
	qualities := make([]*waterquality.Quality, len(records))
	for i, r := range records {
		qualities[i] = r.Data
	}
	return qualities
}

func (s *CollectionService) CreateCollectionRecord(ctx context.Context, req *collectionApi.CreateCollectionRecordReq) (*collectionApi.Empty, error) {
	guess, err := s.waterQualityClient.GuessLevel(
		ctx,
		req.Data,
	)
	var level int64
	if err != nil {
		log.Errorf("failed to guess level: %v", err)
		level = -1
	} else {
		level = guess.Level
	}
	err = s.db.CreateCollectionRecord(&model.CollectionRecord{
		DeviceID:       req.DeviceId,
		ReceivedAt:     time.UnixMilli(req.ReceivedAt),
		CollectionData: proto2Data(req.Data),
		Level:          level,
	})
	if err != nil {
		return nil, err
	}
	err = s.ccache.UpdateDeviceLastReport(context.Background(), req.DeviceId, &collectionApi.DeviceLastReport{
		ReceivedAt: req.ReceivedAt,
		Data:       req.Data,
		Level:      level,
	})
	if err != nil {
		log.Errorf("failed to update device last report: %v", err)
	}
	resp, err := s.ListCollectionRecord(
		ctx,
		&collectionApi.ListCollectionRecordReq{
			DeviceId: req.DeviceId,
			Page:     1,
			Size:     24,
			Order:    collectionApi.CollectionRecordOrder_TIMESTAMP,
			Sort:     collectionApi.Sort_DESC,
		},
	)
	if err != nil {
		log.Errorf("failed to list collection record: %v", err)
	} else {
		pg, err := s.waterQualityClient.PredictAndGuess(
			ctx,
			&waterquality.PredictAndGuessReq{
				Qualities: collectionRecords2Qualities(resp.Records),
				LookBack:  3,
				Horizon:   24,
			},
		)
		if err != nil {
			log.Errorf("failed to predict and guess: %v", err)
		}
		err = s.ccache.UpdatePredictQuality(
			ctx,
			req.DeviceId,
			pg,
		)
		if err != nil {
			log.Errorf("failed to update predict quality: %v", err)
		}
	}
	return &collectionApi.Empty{}, nil
}

func proto2Data(data *waterquality.Quality) *model.CollectionData {
	return &model.CollectionData{
		Timestamp:   time.UnixMilli(data.Timestamp),
		GeoPoint:    model.GeoPoint{Lat: data.GeoPoint.Lat, Lon: data.GeoPoint.Lon},
		Temperature: data.Temperature,
		Ph:          data.Ph,
		Tsw:         data.Tsw,
		Tds:         data.Tds,
		Oxygen:      data.Oxygen,
	}
}

func proto2Record(record *collectionApi.CollectionRecord) *model.CollectionRecord {
	return &model.CollectionRecord{
		DeviceID:       record.DeviceId,
		CreatedAt:      time.UnixMilli(record.CreatedAt),
		ReceivedAt:     time.UnixMilli(record.ReceivedAt),
		CollectionData: proto2Data(record.Data),
		Level:          record.Level,
	}
}

func data2Proto(data *model.CollectionData) *waterquality.Quality {
	return &waterquality.Quality{
		Timestamp:   data.Timestamp.UnixMilli(),
		GeoPoint:    &waterquality.GeoPoint{Lat: data.GeoPoint.Lat, Lon: data.GeoPoint.Lon},
		Temperature: data.Temperature,
		Ph:          data.Ph,
	}
}

func record2Proto(record *model.CollectionRecord) *collectionApi.CollectionRecord {
	return &collectionApi.CollectionRecord{
		DeviceId:  record.DeviceID,
		CreatedAt: record.CreatedAt.UnixMilli(),
		Data:      data2Proto(record.CollectionData),
		Level:     record.Level,
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

func (s *CollectionService) GetPredictQuality(ctx context.Context, req *collectionApi.GetPredictQualityReq) (*waterquality.PredictAndGuessResp, error) {
	return s.ccache.GetPredictQuality(ctx, req.DeviceId)
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
	var ch = make(chan *collectionApi.CreateCollectionRecordReq)
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

func (s *CollectionService) GetDeviceLastReport(ctx context.Context, req *collectionApi.GetDeviceLastReportReq) (*collectionApi.DeviceLastReport, error) {
	return s.ccache.GetDeviceLastReport(ctx, req.Id)
}

func (s *CollectionService) GetLatestIdWithinRange(ctx context.Context, req *collectionApi.GetLatestWithinRangeReq) (*collectionApi.GetLatestIdWithinRangeResp, error) {
	ids, err := s.db.GetDeviceIDsWithinRange(req.CenterLat, req.CenterLng, req.RadiusMeters, time.Now(), time.Time{})
	if err != nil {
		return nil, err
	}
	return &collectionApi.GetLatestIdWithinRangeResp{
		Ids: ids,
	}, nil
}

func (s *CollectionService) GetLatestRecordsWithinRange(ctx context.Context, req *collectionApi.GetLatestWithinRangeReq) (*collectionApi.GetLatestRecordsWithinRangeResp, error) {
	records, err := s.db.GetLatestRecordsWithinRange(req.CenterLat, req.CenterLng, req.RadiusMeters, time.Now(), time.Time{})
	if err != nil {
		return nil, err
	}
	return &collectionApi.GetLatestRecordsWithinRangeResp{
		Records: records2Proto(records),
	}, nil
}

func (s *CollectionService) GetStreamLatestIdWithinRange(req *collectionApi.GetStreamLatestWithinRangeReq, resp collectionApi.Collection_GetStreamLatestIdWithinRangeServer) error {
	if req.Interval == 0 {
		req.Interval = 3
	}
	after := time.Time{}
	ticker := time.NewTicker(time.Second * time.Duration(req.Interval))
	defer ticker.Stop()
	currentIds := set.New[uint64]()
	for {
		select {
		case <-resp.Context().Done():
			return nil
		case <-ticker.C:
			before := time.Now()
			ids, err := s.db.GetDeviceIDsWithinRange(req.CenterLat, req.CenterLng, req.RadiusMeters, before, after)
			if err != nil {
				log.Errorf("get ids within range error: %v", err)
				return err
			}
			if after.IsZero() {
				currentIds.Push(ids...)
				err = resp.Send(&collectionApi.GetStreamLatestIdWithinRangeResp{
					Ids:  ids,
					Type: collectionApi.GetStreamLatestWithinRangeRespType_ADD,
				})
				if err != nil {
					log.Errorf("send add ids error: %v", err)
					return err
				}
			} else {
				idset := set.New[uint64]()
				idset.Push(ids...)
				newIds := idset.Difference(currentIds).Slice()
				nids, err := s.db.GetIDsNotWithinRange(currentIds.Difference(idset).Slice(), req.CenterLat, req.CenterLng, req.RadiusMeters, after)
				if err != nil {
					log.Errorf("get ids not within range error: %v", err)
					return err
				}
				if len(nids) != 0 {
					for _, id := range nids {
						currentIds.Remove(id)
					}
					err = resp.Send(&collectionApi.GetStreamLatestIdWithinRangeResp{
						Ids:  nids,
						Type: collectionApi.GetStreamLatestWithinRangeRespType_REMOVE,
					})
					if err != nil {
						log.Errorf("send remove ids error: %v", err)
						return err
					}
				}
				if len(newIds) != 0 {
					currentIds.Push(newIds...)
					err = resp.Send(&collectionApi.GetStreamLatestIdWithinRangeResp{
						Ids:  newIds,
						Type: collectionApi.GetStreamLatestWithinRangeRespType_ADD,
					})
					if err != nil {
						log.Errorf("send add ids error: %v", err)
						return err
					}
				}
			}
			after = before
		}
	}
}

func (s *CollectionService) GetStreamLatestRecordsWithinRange(req *collectionApi.GetStreamLatestWithinRangeReq, resp collectionApi.Collection_GetStreamLatestRecordsWithinRangeServer) error {
	if req.Interval == 0 {
		req.Interval = 3
	}
	after := time.Time{}
	ticker := time.NewTicker(time.Second * time.Duration(req.Interval))
	defer ticker.Stop()
	currentIds := set.New[uint64]()
	for {
		select {
		case <-resp.Context().Done():
			return nil
		case <-ticker.C:
			before := time.Now()
			records, err := s.db.GetLatestRecordsWithinRange(req.CenterLat, req.CenterLng, req.RadiusMeters, before, after)
			if err != nil {
				log.Errorf("get latest records within range error: %v", err)
				return err
			}
			idset := set.New[uint64]()
			if len(records) != 0 {
				for _, r := range records {
					currentIds.Push(r.DeviceID)
					idset.Push(r.DeviceID)
				}
				err = resp.Send(&collectionApi.GetStreamLatestRecordsWithinRangeResp{
					Records: records2Proto(records),
					Type:    collectionApi.GetStreamLatestWithinRangeRespType_ADD,
				})
				if err != nil {
					log.Errorf("send add records error: %v", err)
					return err
				}
			}
			if !after.IsZero() {
				nids, err := s.db.GetIDsNotWithinRange(currentIds.Difference(idset).Slice(), req.CenterLat, req.CenterLng, req.RadiusMeters, after)
				if err != nil {
					log.Errorf("get ids not within range error: %v", err)
					return err
				}
				if len(nids) != 0 {
					for _, id := range nids {
						currentIds.Remove(id)
					}
					err = resp.Send(&collectionApi.GetStreamLatestRecordsWithinRangeResp{
						Ids:  nids,
						Type: collectionApi.GetStreamLatestWithinRangeRespType_REMOVE,
					})
					if err != nil {
						log.Errorf("send remove ids error: %v", err)
						return err
					}
				}
			}
			after = before
		}
	}
}
