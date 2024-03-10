package client

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var (
	addr     string
	clientid string
	password string
)

var ClientCmd = &cobra.Command{
	Use:   "client",
	Short: "client",
	Long:  `https://github.com/I-m-Surrounded-by-IoT/backend`,
	Run:   ClientRun,
}

func ClientRun(cmd *cobra.Command, args []string) {
	if addr == "" {
		log.Fatal("mtqq address is required, please use -a to specify it.")
	}
	if clientid == "" {
		log.Fatal("mtqq client id is required, please use -c to specify it.")
	}
	if password == "" {
		log.Fatal("mtqq password is required, please use -p to specify it.")
	}
	_, deviceID, found := strings.Cut(clientid, "-")
	if !found {
		log.Fatalf("failed to get device id: %v", clientid)
	}
	opt := mqtt.NewClientOptions().
		AddBroker(addr).
		SetUsername("client").
		SetClientID(clientid).
		SetPassword(password).
		SetAutoReconnect(true)
	cli := mqtt.NewClient(opt)
	if token := cli.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("failed to connect mqtt server: %v", token.Error())
	}
	log.Info("connected to mqtt server")
	timer := time.NewTicker(time.Second * 5)
	defer timer.Stop()
	for range timer.C {
		data := &collection.CollectionData{
			Timestamp: time.Now().UnixMilli(),
			GeoPoint: &collection.GeoPoint{
				Lat: rand.Float64() * 100,
				Lon: rand.Float64() * 100,
			},
			Temperature: rand.Float32() * 40,
		}
		log.Infof("publish data: %+v", data)
		bytes, err := proto.Marshal(data)
		if err != nil {
			log.Errorf("failed to marshal data: %v", err)
			continue
		}
		if token := cli.Publish(fmt.Sprintf("device/%s/report", deviceID), 2, false, bytes); !token.WaitTimeout(time.Second * 5) {
			log.Errorf("failed to publish data: %v", token.Error())
		}
	}
}

func init() {
	ClientCmd.PersistentFlags().StringVarP(&addr, "addr", "a", "", "mqtt address")
	ClientCmd.PersistentFlags().StringVarP(&clientid, "clientid", "c", "", "mqtt client id")
	ClientCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "mqtt password")
}
