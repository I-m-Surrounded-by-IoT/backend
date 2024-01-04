package collector

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

func newApp(logger log.Logger, gs *utils.GrpcGatewayServer, s *utils.TcpServer, r registry.Registrar) *kratos.App {
	es, err := gs.Endpoints()
	if err != nil {
		panic(err)
	}
	e, err := s.Endpoint()
	if err != nil {
		panic(err)
	}
	es = append(es, e)
	return kratos.New(
		kratos.ID(id),
		kratos.Name("collector"),
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

var CollectorCmd = &cobra.Command{
	Use:   "collector",
	Short: "Start backend collector",
	Run:   Server,
}

func Server(cmd *cobra.Command, args []string) {
	bc := conf.CollectorServer{
		TcpServer:  conf.DefaultTcpServer(),
		GrpcServer: conf.DefaultGrpcServer(),
		Registry:   conf.DefaultRegistry(),
		Config:     &conf.CollectorConfig{},
		Kafka:      conf.DefaultKafka(),
	}

	if flagconf != "" {
		c := config.New(
			config.WithSource(
				file.NewSource(flagconf),
			),
		)
		defer c.Close()

		if err := c.Load(); err != nil {
			panic(err)
		}
		if err := c.Scan(&bc); err != nil {
			panic(err)
		}
	}

	if err := env.Parse(&bc); err != nil {
		panic(err)
	}

	id = fmt.Sprintf("%s-%s", id, bc.GrpcServer.Addr)

	logger := log.With(log.NewStdLogger(logrus.StandardLogger().Writer()),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", "collector",
		"service.version", flags.Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	app, cleanup, err := wireApp(bc.GrpcServer, bc.TcpServer, bc.Registry, bc.Config, bc.Kafka, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func init() {
	CollectorCmd.PersistentFlags().StringVarP(&flagconf, "conf", "c", "", "config path, eg: -c config.yaml")
}
