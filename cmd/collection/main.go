package collection

import (
	"fmt"
	"os"

	"github.com/I-m-Surrounded-by-IoT/backend/cmd/flags"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	collection "github.com/I-m-Surrounded-by-IoT/backend/internal/server/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/caarlos0/env/v9"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	_ "go.uber.org/automaxprocs"
)

var (
	flagconf string

	id, _ = os.Hostname()
)

func newApp(logger log.Logger, s *utils.GrpcGatewayServer, c *collection.CollectionConsumerServer, r registry.Registrar) *kratos.App {
	es, err := s.Endpoints()
	if err != nil {
		logrus.Fatalf("failed to get endpoints: %v", err)
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name("collection"),
		kratos.Version(flags.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			s,
			c,
		),
		kratos.Registrar(r),
		kratos.Endpoint(es...),
	)
}

var DatabaseCmd = &cobra.Command{
	Use:   "collection",
	Short: "Start backend collection database",
	Run:   Server,
}

func Server(cmd *cobra.Command, args []string) {
	bc := conf.CollectionServer{
		Server: conf.DefaultGrpcServer(),
		Database: &conf.DatabaseServerConfig{
			Name: "collection",
		},
		Registry: conf.DefaultRegistry(),
		Config:   &conf.CollectionConfig{},
		Kafka:    conf.DefaultKafka(),
		Redis:    &conf.RedisConfig{},
	}

	if flagconf != "" {
		c := config.New(
			config.WithSource(
				file.NewSource(flagconf),
			),
		)
		defer c.Close()

		if err := c.Load(); err != nil {
			logrus.Fatalf("error loading config: %v", err)
		}
		if err := c.Scan(&bc); err != nil {
			logrus.Fatalf("error scanning config: %v", err)
		}
	}

	if err := env.Parse(&bc); err != nil {
		logrus.Fatalf("error parsing config: %v", err)
	}
	if err := env.ParseWithOptions(&bc, env.Options{
		Prefix: "COLLECTION_",
	}); err != nil {
		logrus.Fatalf("error parsing config: %v", err)
	}

	id = fmt.Sprintf("%s-%s", id, bc.Server.Addr)

	logger := utils.TransLogrus(logrus.StandardLogger())

	app, cleanup, err := wireApp(bc.Server, bc.Registry, bc.Database, bc.Kafka, bc.Config, bc.Redis, logger)
	if err != nil {
		logrus.Fatalf("failed to new app: %v", err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		logrus.Fatalf("failed to run app: %v", err)
	}
}

func init() {
	DatabaseCmd.PersistentFlags().StringVarP(&flagconf, "conf", "c", "", "config path, eg: -c config.yaml")
}
