package log

import (
	"context"
	"fmt"
	"time"

	logApi "github.com/I-m-Surrounded-by-IoT/backend/api/log"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/log/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type LogService struct {
	db *dbUtils
	logApi.UnimplementedLogServer
}

func NewLogService(dc *conf.DatabaseServerConfig, lc *conf.LogConfig) *LogService {
	d, err := dbdial.Dial(context.Background(), dc)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	if dc.AutoMigrate {
		log.Infof("auto migrate database...")
		err = d.AutoMigrate(
			new(model.DeviceLog),
		)
		if err != nil {
			log.Fatalf("failed to migrate database: %v", err)
		}
	}

	return &LogService{
		db: newDBUtils(d),
	}
}

func (ls *LogService) CreateDeviceLog(ctx context.Context, req *logApi.DeviceLog) (*logApi.Empty, error) {
	err := ls.db.CreateDeviceLog(&model.DeviceLog{
		DeviceID:  req.DeviceId,
		Timestamp: time.UnixMilli(req.Timestamp),
		Message:   req.Message,
		Level:     log.Level(req.Level),
	})
	if err != nil {
		return nil, err
	}
	return &logApi.Empty{}, nil
}

func deviceLog2Proto(log *model.DeviceLog) *logApi.DeviceLog {
	return &logApi.DeviceLog{
		Id:        log.ID,
		DeviceId:  log.DeviceID,
		Timestamp: log.Timestamp.UnixMilli(),
		Message:   log.Message,
		Level:     uint32(log.Level),
	}
}

func deviceLogs2Proto(logs []*model.DeviceLog) []*logApi.DeviceLog {
	var res []*logApi.DeviceLog = make([]*logApi.DeviceLog, len(logs))
	for i, log := range logs {
		res[i] = deviceLog2Proto(log)
	}
	return res
}

func (ls *LogService) ListDeviceLog(ctx context.Context, req *logApi.ListDeviceLogReq) (*logApi.ListDeviceLogResp, error) {
	opts := []func(*gorm.DB) *gorm.DB{}
	if req.DeviceId != 0 {
		opts = append(opts, utils.WithDeviceIDEq(req.DeviceId))
	}
	if req.LevelFilter != "" {
		opts = append(opts, model.WithLevelFilter(req.LevelFilter))
	}
	if req.Before != 0 {
		opts = append(opts, utils.WithTimestampBefore(req.Before))
	}
	if req.After != 0 {
		opts = append(opts, utils.WithTimestampAfter(req.After))
	}

	count, err := ls.db.CountDeviceLog(opts...)
	if err != nil {
		return nil, err
	}

	opts = append(opts, utils.WithPageAndPageSize(int(req.Page), int(req.Size)))
	switch req.Order {
	case logApi.DeviceLogOrder_CREATED_AT:
		opts = append(opts, utils.WithOrder(fmt.Sprintf("created_at %s", req.Sort)))
	default: // logApi.DeviceLogOrder_TIMESTAMP
		opts = append(opts, utils.WithOrder(fmt.Sprintf("timestamp %s", req.Sort)))
	}

	logs, err := ls.db.ListDeviceLog(opts...)
	if err != nil {
		return nil, err
	}
	return &logApi.ListDeviceLogResp{
		Logs:  deviceLogs2Proto(logs),
		Total: count,
	}, nil
}
