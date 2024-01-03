package cmd

import (
	"fmt"
	"os"

	"github.com/I-m-Surrounded-by-IoT/backend/cmd/client"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/database"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/flags"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/gateway"
	"github.com/I-m-Surrounded-by-IoT/backend/internal/bootstrap"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "backend",
	Short: "backend",
	Long:  `https://github.com/I-m-Surrounded-by-IoT/backend`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		s, err := utils.GetEnvFiles(".")
		if err != nil {
			logrus.Fatalf("error getting .env files: %v", err)
		}
		if len(s) != 0 {
			err = godotenv.Overload(s...)
			if err != nil {
				logrus.Fatalf("error loading .env files: %v", err)
			}
		}
		_ = bootstrap.InitLog()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(collector.CollectorCmd)
	RootCmd.AddCommand(gateway.GatewayCmd)
	RootCmd.AddCommand(database.DatabaseCmd)
	RootCmd.AddCommand(client.ClientCmd)
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&flags.Dev, "dev", "d", false, "dev mode")
}
