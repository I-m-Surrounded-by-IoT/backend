package notify

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
	"github.com/go-kratos/kratos/v2/registry"

	_ "go.uber.org/automaxprocs"
)

var (
	flagconf string

	id, _ = os.Hostname()
)

func newApp(logger log.Logger, gs *utils.GrpcGatewayServer, r registry.Registrar) *kratos.App {
	es, err := gs.Endpoints()
	if err != nil {
		logrus.Fatalf("failed to get endpoints: %v", err)
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name("notify"),
		kratos.Version(flags.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(r),
		kratos.Endpoint(es...),
	)
}

var NotifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "Start backend notify",
	Run:   Server,
}

func Server(cmd *cobra.Command, args []string) {
	bc := conf.NotifyServer{
		Server:   conf.DefaultGrpcServer(),
		Registry: conf.DefaultRegistry(),
		Config:   &conf.NotifyConfig{},
		Kafka:    conf.DefaultKafka(),
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
			logrus.Fatalf("error parsing config: %v", err)
		}
	}

	if err := env.Parse(&bc); err != nil {
		logrus.Fatalf("error parsing config: %v", err)
	}
	if err := env.ParseWithOptions(&bc, env.Options{
		Prefix: "NOTIFY_",
	}); err != nil {
		logrus.Fatalf("error parsing config: %v", err)
	}

	id = fmt.Sprintf("%s-%s", id, bc.Server.Addr)

	err := utils.InitTracer(bc.Server.TracingEndpoint, "notify")
	if err != nil {
		logrus.Fatalf("failed to init tracer: %v", err)
	}

	logger := utils.TransLogrus(logrus.StandardLogger())

	app, cleanup, err := wireApp(bc.Server, bc.Registry, bc.Config, bc.Kafka, logger)
	if err != nil {
		logrus.Fatalf("wire app error: %v", err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		logrus.Fatalf("app run error: %v", err)
	}
}

func init() {
	NotifyCmd.PersistentFlags().StringVarP(&flagconf, "conf", "c", "", "config path, eg: -c config.yaml")
}
