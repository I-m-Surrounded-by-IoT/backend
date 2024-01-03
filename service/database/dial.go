package database

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/cmd/flags"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(ctx context.Context, dbConf *conf.DatabaseConfig) (*gorm.DB, error) {
	dialector, err := createDialector(dbConf)
	if err != nil {
		return nil, err
	}

	var opts []gorm.Option
	opts = append(opts, &gorm.Config{
		TranslateError: true,
		Logger:         newDBLogger(),
		PrepareStmt:    true,
	})
	d, err := gorm.Open(dialector, opts...)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func createDialector(dbConf *conf.DatabaseConfig) (dialector gorm.Dialector, err error) {
	var dsn string
	switch dbConf.Type {
	case conf.DatabaseType_MYSQL:
		if dbConf.Port == 0 {
			dsn = fmt.Sprintf("%s:%s@unix(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=true&tls=%s",
				dbConf.User,
				dbConf.Password,
				dbConf.Host,
				dbConf.Name,
				dbConf.SslMode,
			)
			log.Infof("mysql database: %s", dbConf.Host)
		} else {
			dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&interpolateParams=true&tls=%s",
				dbConf.User,
				dbConf.Password,
				dbConf.Host,
				dbConf.Port,
				dbConf.Name,
				dbConf.SslMode,
			)
			log.Infof("mysql database tcp: %s:%d", dbConf.Host, dbConf.Port)
		}
		dialector = mysql.New(mysql.Config{
			DSN:                       dsn,
			DefaultStringSize:         256,
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		})
	case conf.DatabaseType_SQLITE:
		if dbConf.Name == "memory" || strings.HasPrefix(dbConf.Name, ":memory:") {
			dsn = "file::memory:?cache=shared&_journal_mode=WAL&_vacuum=incremental&_pragma=foreign_keys(1)"
			log.Infof("sqlite3 database memory")
		} else {
			if !strings.HasSuffix(dbConf.Name, ".db") {
				dbConf.Name = dbConf.Name + ".db"
			}
			dsn = fmt.Sprintf("%s?_journal_mode=WAL&_vacuum=incremental&_pragma=foreign_keys(1)", dbConf.Name)
			log.Infof("sqlite3 database file: %s", dbConf.Name)
		}
		dialector = sqlite.Open(dsn)
	case conf.DatabaseType_POSTGRESQL:
		if dbConf.Port == 0 {
			dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s",
				dbConf.Host,
				dbConf.User,
				dbConf.Password,
				dbConf.Name,
				dbConf.SslMode,
			)
			log.Infof("postgres database: %s", dbConf.Host)
		} else {
			dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
				dbConf.Host,
				dbConf.Port,
				dbConf.User,
				dbConf.Password,
				dbConf.Name,
				dbConf.SslMode,
			)
			log.Infof("postgres database tcp: %s:%d", dbConf.Host, dbConf.Port)
		}
		dialector = postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		})
	default:
		log.Fatalf("unknown database type: %s", dbConf.Type)
	}
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
