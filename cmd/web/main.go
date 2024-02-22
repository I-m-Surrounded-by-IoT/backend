package web

import (
	"fmt"
	"os"

	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/I-m-Surrounded-by-IoT/backend/cmd/flags"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/internal/bootstrap"
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

func newApp(logger log.Logger, s *http.Server, r registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name("web"),
		kratos.Version(flags.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			s,
		),
		kratos.Registrar(r),
	)
}

var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "Start backend web",
	PreRun: func(cmd *cobra.Command, args []string) {
		err := bootstrap.InitGinMode(cmd.Context())
		if err != nil {
			logrus.Fatalf("error init gin mode: %v", err)
		}
	},
	Run: Server,
}

const defaultJwtSecret = "jwt_secret"

func Server(cmd *cobra.Command, args []string) {
	uc := conf.WebServer{
		Server:   conf.DefaultWebServer(),
		Registry: conf.DefaultRegistry(),
		Config: &conf.WebConfig{
			Jwt: &conf.WebConfig_JWT{
				Secret: defaultJwtSecret,
				Expire: "24h",
			},
		},
		Redis: &conf.RedisConfig{},
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

	logger := utils.TransLogrus(logrus.StandardLogger())

	app, cleanup, err := wireApp(uc.Server, uc.Registry, uc.Config, uc.Redis, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func init() {
	WebCmd.PersistentFlags().StringVarP(&flagconf, "conf", "c", "", "config path, eg: -c config.yaml")
}
