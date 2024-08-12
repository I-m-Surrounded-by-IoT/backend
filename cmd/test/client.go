package test

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	addr     string
	interval int
)

var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  `https://github.com/I-m-Surrounded-by-IoT/backend`,
	Run:   ClientRun,
}

func ClientRun(cmd *cobra.Command, args []string) {
	if addr == "" {
		log.Fatal("mtqq address is required, please use -a to specify it.")
	}
	opt := mqtt.NewClientOptions().
		AddBroker(addr).
		SetUsername("client").
		SetClientID("test").
		SetPassword("test").
		SetAutoReconnect(true)
	cli := mqtt.NewClient(opt)
	if token := cli.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("failed to connect mqtt server: %v", token.Error())
	}
	log.Info("connected to mqtt server")
	for {
		if token := cli.Publish(fmt.Sprintf("device/%d/control", 2), 2, false, "report"); !token.WaitTimeout(time.Second * 5) {
			log.Errorf("failed to publish data: %v", token.Error())
		}
		time.Sleep(time.Second * time.Duration(interval))
	}
}

func init() {
	TestCmd.PersistentFlags().StringVarP(&addr, "addr", "a", "", "mqtt address")
	TestCmd.PersistentFlags().IntVarP(&interval, "interval", "i", 15, "interval to publish data")
}
