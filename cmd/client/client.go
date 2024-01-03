package client

import (
	"math/rand"
	"net"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/proto/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/proto/gateway"
	tcpconn "github.com/I-m-Surrounded-by-IoT/backend/utils/tcpConn"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	proto "google.golang.org/protobuf/proto"
)

var (
	gatewayAddr string
)

var ClientCmd = &cobra.Command{
	Use:   "client",
	Short: "client",
	Long:  `https://github.com/I-m-Surrounded-by-IoT/backend`,
	Run:   ClientRun,
}

func ClientRun(cmd *cobra.Command, args []string) {
	c, err := net.Dial("tcp", gatewayAddr)
	if err != nil {
		log.Fatalf("error dialing: %v", err)
	}
	conn := tcpconn.NewConn(c)
	defer conn.Close()
	err = conn.ClientSayHello()
	if err != nil {
		log.Fatalf("error saying hello: %v", err)
	}
	msg := gateway.GetServerReq{
		// TODO: mac not used
		Mac: "00:00:00:00:00:00",
	}
	b, err := proto.Marshal(&msg)
	if err != nil {
		log.Fatalf("error marshaling: %v", err)
	}
	err = conn.Send(b)
	if err != nil {
		log.Fatalf("error sending: %v", err)
	}
	b, err = conn.NextMessage()
	if err != nil {
		log.Fatalf("error receiving: %v", err)
	}
	resp := gateway.GetServerResp{}
	err = proto.Unmarshal(b, &resp)
	if err != nil {
		log.Fatalf("error unmarshaling: %v", err)
	}
	log.Infof("receive message: %v", resp.ServerAddr)
	conn.Close()
	c2, err2 := net.Dial("tcp", resp.ServerAddr)
	if err2 != nil {
		log.Fatalf("error dialing: %v", err2)
	}
	conn2 := tcpconn.NewConn(c2)
	defer conn2.Close()
	handlerCollector(conn2)
}

func handlerCollector(conn *tcpconn.Conn) {
	err := conn.ClientSayHello()
	if err != nil {
		log.Fatalf("error saying hello: %v", err)
	}
	log.Infof("client say hello success")
	msg := collector.Message{
		Type: collector.MessageType_ReportMac,
		Payload: &collector.Message_Mac{
			Mac: "00:00:00:00:00:00",
		},
	}
	b, err := proto.Marshal(&msg)
	if err != nil {
		log.Fatalf("error marshaling: %v", err)
	}
	err = conn.Send(b)
	if err != nil {
		log.Fatalf("error sending: %v", err)
	}
	log.Infof("report mac success")
	for {
		b, err := conn.NextMessage()
		if err != nil {
			log.Fatalf("error receiving: %v", err)
		}
		msg := collector.Message{}
		err = proto.Unmarshal(b, &msg)
		if err != nil {
			log.Fatalf("error unmarshaling: %v", err)
		}
		switch msg.Type {
		case collector.MessageType_ReportImmediately:
			msg.Type = collector.MessageType_Report
			msg.Payload = &collector.Message_ReportPayload{
				ReportPayload: &collector.ReportPayload{
					Timestamp: uint64(time.Now().UnixMicro()),
					GeoPoint: &collector.GeoPoint{
						Latitude:  rand.Float64(),
						Longitude: rand.Float64(),
					},
					Temperature: rand.Float64(),
				},
			}
			b, err = proto.Marshal(&msg)
			if err != nil {
				log.Fatalf("error marshaling: %v", err)
			}
			err = conn.Send(b)
			if err != nil {
				log.Fatalf("error sending: %v", err)
			}
		default:
			log.Errorf("invalid message type: %v", msg.Type)
		}
	}
}

func init() {
	ClientCmd.PersistentFlags().StringVar(&gatewayAddr, "gateway", "", "gateway address")
}
