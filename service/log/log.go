package log

import (
	"context"

	logApi "github.com/I-m-Surrounded-by-IoT/backend/api/log"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/log/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/dbdial"
	log "github.com/sirupsen/logrus"
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
