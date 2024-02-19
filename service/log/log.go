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

func NewLogService(c *conf.DatabaseConfig) *LogService {
	d, err := dbdial.NewDatabase(context.Background(), c)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	if c.AutoMigrate {
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
