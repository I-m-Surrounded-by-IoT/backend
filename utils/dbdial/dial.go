package dbdial

import (
	"context"
	"fmt"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/cmd/flags"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig interface {
	GetAutoMigrate() bool
	GetHost() string
	GetName() string
	GetPassword() string
	GetPort() uint32
	GetSslMode() string
	GetUser() string
}

func Dial(ctx context.Context, dbConf DatabaseConfig) (*gorm.DB, error) {
	dialector, err := createDialector(dbConf)
	if err != nil {
		return nil, err
	}

	var opts []gorm.Option
	opts = append(opts, &gorm.Config{
		TranslateError:                           true,
		Logger:                                   newDBLogger(),
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: false,
		IgnoreRelationshipsWhenMigrating:         false,
	})
	d, err := gorm.Open(dialector, opts...)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func createDialector(dbConf DatabaseConfig) (dialector gorm.Dialector, err error) {
	var dsn string
	if dbConf.GetPort() == 0 {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
			dbConf.GetHost(),
			dbConf.GetUser(),
			dbConf.GetPassword(),
			dbConf.GetName(),
			dbConf.GetSslMode(),
		)
		log.Infof("postgres database: %s", dbConf.GetHost())
	} else {
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			dbConf.GetHost(),
			dbConf.GetPort(),
			dbConf.GetUser(),
			dbConf.GetPassword(),
			dbConf.GetName(),
			dbConf.GetSslMode(),
		)
		log.Infof("postgres database tcp: %s:%d", dbConf.GetHost(), dbConf.GetPort())
	}
	dialector = postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})
	return
}

func newDBLogger() logger.Interface {
	var logLevel logger.LogLevel
	if flags.Dev {
		logLevel = logger.Info
	} else {
		logLevel = logger.Warn
	}
	return logger.New(
		log.StandardLogger(),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      !flags.Dev,
			Colorful:                  utils.ForceColor(),
		},
	)
}
