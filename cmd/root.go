package cmd

import (
	"fmt"
	"os"

	"github.com/I-m-Surrounded-by-IoT/backend/cmd/captcha"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/client"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/device"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/email"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/flags"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/log"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/message"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/notify"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/user"
	"github.com/I-m-Surrounded-by-IoT/backend/cmd/web"
	"github.com/I-m-Surrounded-by-IoT/backend/internal/bootstrap"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:               "backend",
	Short:             "backend",
	Long:              `https://github.com/I-m-Surrounded-by-IoT/backend`,
	PersistentPreRunE: PersistentPreRunE,
}

func PersistentPreRunE(cmd *cobra.Command, args []string) error {
	return bootstrap.InitLog()
}

func Execute() {
	_ = bootstrap.LoadEnvFromFile()
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(collector.CollectorCmd)
	RootCmd.AddCommand(collection.DatabaseCmd)
	RootCmd.AddCommand(device.DeviceCmd)
	RootCmd.AddCommand(client.ClientCmd)
	RootCmd.AddCommand(log.LogCmd)
	RootCmd.AddCommand(user.UserCmd)
	RootCmd.AddCommand(web.WebCmd)
	RootCmd.AddCommand(email.EmailCmd)
	RootCmd.AddCommand(captcha.CaptchaCmd)
	RootCmd.AddCommand(notify.NotifyCmd)
	RootCmd.AddCommand(message.MessageCmd)
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&flags.Dev, "dev", "d", false, "dev mode")
}
