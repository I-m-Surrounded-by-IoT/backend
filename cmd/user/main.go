package user

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
		kratos.Name("database"),
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

var UserCmd = &cobra.Command{
	Use:   "user",
	Short: "Start backend user",
	Run:   Server,
}

func Server(cmd *cobra.Command, args []string) {
	uc := conf.UserServer{
		Server:   conf.DefaultGrpcServer(),
		Database: &conf.DatabaseConfig{},
		Registry: conf.DefaultRegistry(),
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
		if err := c.Scan(&uc); err != nil {
			logrus.Fatalf("error scanning config: %v", err)
		}
	}

	if err := env.Parse(&uc); err != nil {
		logrus.Fatalf("error parsing config: %v", err)
	}

	id = fmt.Sprintf("%s-%s", id, uc.Server.Addr)

	logger := log.With(log.NewStdLogger(logrus.StandardLogger().Writer()),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", "database",
		"service.version", flags.Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	app, cleanup, err := wireApp(uc.Server, uc.Registry, uc.Database, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func init() {
	UserCmd.PersistentFlags().StringVarP(&flagconf, "conf", "c", "", "config path, eg: -c config.yaml")
}