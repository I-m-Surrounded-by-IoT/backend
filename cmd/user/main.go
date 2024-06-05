package user

import (
	"context"
	"fmt"
	"os"

	"github.com/I-m-Surrounded-by-IoT/backend/cmd/flags"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/internal/bootstrap"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/caarlos0/env/v9"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/I-m-Surrounded-by-IoT/backend/cmd/user/create"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
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
		logrus.Fatalf("failed to get endpoints: %v", err)
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name("user"),
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
	Use:              "user",
	Short:            "Start backend user",
	PersistentPreRun: PersistentPreRun,
	Run:              Server,
}

func Load(path string) (*conf.UserServer, error) {
	uc := &conf.UserServer{
		Server: conf.DefaultGrpcServer(),
		Database: &conf.DatabaseServerConfig{
			Name: "users",
		},
		Registry: conf.DefaultRegistry(),
		Config:   &conf.UserConfig{},
		Redis:    &conf.RedisConfig{},
	}

	if path != "" {
		c := config.New(
			config.WithSource(
				file.NewSource(path),
			),
		)
		defer c.Close()

		if err := c.Load(); err != nil {
			return nil, err
		}
		if err := c.Scan(uc); err != nil {
			return nil, err
		}
	}

	if err := env.Parse(uc); err != nil {
		return nil, err
	}
	if err := env.ParseWithOptions(uc, env.Options{
		Prefix: "USER_",
	}); err != nil {
		logrus.Fatalf("error parsing config: %v", err)
	}
	return uc, nil
}

func PersistentPreRun(cmd *cobra.Command, args []string) {
	_ = bootstrap.InitLog()
	uc, err := Load(flagconf)
	if err != nil {
		logrus.Fatalf("failed to load config: %v", err)
	}
	cmd.SetContext(context.WithValue(cmd.Context(), "config", uc))
	logrus.Info("log config success")
}

func Server(cmd *cobra.Command, args []string) {
	uc := cmd.Context().Value("config").(*conf.UserServer)

	id = fmt.Sprintf("%s-%s", id, uc.Server.Addr)

	err := utils.InitTracer(uc.Server.TracingEndpoint, "user")
	if err != nil {
		logrus.Fatalf("failed to init tracer: %v", err)
	}

	logger := utils.TransLogrus(logrus.StandardLogger())

	app, cleanup, err := wireApp(uc.Server, uc.Registry, uc.Database, uc.Config, uc.Redis, logger)
	if err != nil {
		logrus.Fatalf("failed to new app: %v", err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		logrus.Fatalf("failed to run app: %v", err)
	}
}

func init() {
	UserCmd.AddCommand(create.CreateCmd)
}

func init() {
	UserCmd.PersistentFlags().StringVarP(&flagconf, "conf", "c", "", "config path, eg: -c config.yaml")
}
