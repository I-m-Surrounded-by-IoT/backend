package device

import (
	"fmt"
	"os"

	"github.com/I-m-Surrounded-by-IoT/backend/cmd/flags"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/caarlos0/env/v9"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"

	_ "go.uber.org/automaxprocs"
)

var (
	flagconf string

	id, _ = os.Hostname()
)

func newApp(logger log.Logger, s *utils.GrpcGatewayServer, r registry.Registrar) *kratos.App {
	es, err := s.Endpoints()
	if err != nil {
		panic(err)
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name("device"),
		kratos.Version(flags.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			s,
		),
		kratos.Registrar(r),
		kratos.Endpoint(es...),
	)
}

var DeviceCmd = &cobra.Command{
	Use:   "device",
	Short: "Start backend device",
	Run:   Server,
}

func Server(cmd *cobra.Command, args []string) {
	bc := conf.DeviceServer{
		Server: conf.DefaultGrpcServer(),
		Database: &conf.DatabaseServerConfig{
			Name: "device",
		},
		Registry: conf.DefaultRegistry(),
		Config:   &conf.DeviceConfig{},
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

	id = fmt.Sprintf("%s-%s", id, bc.Server.Addr)

	logger := log.With(log.NewStdLogger(logrus.StandardLogger().Writer()),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", "device",
		"service.version", flags.Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	app, cleanup, err := wireApp(bc.Server, bc.Registry, bc.Database, bc.Config, bc.Redis, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func init() {
	DeviceCmd.PersistentFlags().StringVarP(&flagconf, "conf", "c", "", "config path, eg: -c config.yaml")
}
