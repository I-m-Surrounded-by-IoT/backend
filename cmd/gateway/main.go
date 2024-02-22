package gateway

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

func newApp(logger log.Logger, s *utils.TcpServer, r registry.Registrar) *kratos.App {
	e, err := s.Endpoint()
	if err != nil {
		panic(err)
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name("gateway"),
		kratos.Version(flags.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			s,
		),
		kratos.Registrar(r),
		kratos.Endpoint(e),
	)
}

var GatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Start backend gateway",
	Run:   Server,
}

func Server(cmd *cobra.Command, args []string) {
	bc := conf.GatewayServer{
		Server:   conf.DefaultTcpServer(),
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
			panic(err)
		}
		if err := c.Scan(&bc); err != nil {
			panic(err)
		}
	}

	if err := env.Parse(&bc); err != nil {
		panic(err)
	}

	id = fmt.Sprintf("%s-%s", id, bc.Server.Addr)

	logger := utils.TransLogrus(logrus.StandardLogger())

	app, cleanup, err := wireApp(bc.Server, bc.Registry, bc.Config, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func init() {
	GatewayCmd.PersistentFlags().StringVarP(&flagconf, "conf", "c", "", "config path, eg: -c config.yaml")
}
