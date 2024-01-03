package gateway

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"

	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	registryClient "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/proto/gateway"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	tcpconn "github.com/I-m-Surrounded-by-IoT/backend/utils/tcpConn"
	"github.com/go-kratos/kratos/v2/registry"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type Servers struct {
	servers []string
	lock    sync.RWMutex
}

func (s *Servers) Set(servers []string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.servers = servers
}

func (s *Servers) Get() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.servers
}

type GatewayService struct {
	servers *Servers
}

func (g *GatewayService) ServeTcp(ctx context.Context, conn net.Conn) error {
	Conn := tcpconn.NewConn(conn)
	defer Conn.Close()
	err := Conn.SayHello()
	if err != nil {
		return fmt.Errorf("say hello to collector failed: %w", err)
	}
	for {
		b, err := Conn.NextMessage()
		if err != nil {
			return fmt.Errorf("receive message from collector failed: %w", err)
		}
		msg := gateway.GetServerReq{}
		err = proto.Unmarshal(b, &msg)
		if err != nil {
			return fmt.Errorf("unmarshal message from collector failed: %w", err)
		}
		log.Infof("receive message from collector: %s", msg.String())
		s := g.servers.Get()
		if len(s) == 0 {
			log.Errorf("no collector available")
			continue
		}
		addr := utils.GetRand(s)
		resp := gateway.GetServerResp{
			ServerAddr: strings.TrimPrefix(addr, "tcp://"),
		}
		b, err = proto.Marshal(&resp)
		if err != nil {
			return fmt.Errorf("marshal message to collector failed: %w", err)
		}
		err = Conn.Send(b)
		if err != nil {
			return fmt.Errorf("send message to collector failed: %w", err)
		}
	}
}

func NewGatewayService(c *conf.GatewayConfig, reg registry.Registrar) utils.TCPHandler {
	ss := &Servers{}
	switch reg := reg.(type) {
	case *registryClient.EtcdRegistry:
		w, err := reg.Watch(context.Background(), "collector")
		if err != nil {
			panic(err)
		}
		go func() {
			for {
				si, err := w.Next()
				if err != nil {
					log.Errorf("watch collector failed: %v", err)
					continue
				}
				servers := make([]string, 0, len(si))
				for _, si2 := range si {
					for _, v := range si2.Endpoints {
						if strings.HasPrefix(v, "tcp://") {
							servers = append(servers, v)
						}
					}
				}
				log.Infof("collector servers: %v", servers)
				ss.Set(servers)
			}
		}()
	case *registryClient.ConsulRegistry:
		w, err := reg.Watch(context.Background(), "collector")
		if err != nil {
			panic(err)
		}
		go func() {
			for {
				si, err := w.Next()
				if err != nil {
					log.Errorf("watch collector failed: %v", err)
					continue
				}
				servers := make([]string, 0, len(si))
				for _, si2 := range si {
					if len(si2.Endpoints) > 0 {
						servers = append(servers, si2.Endpoints[0])
					}
				}
				ss.Set(servers)
			}
		}()
	default:
		panic("invalid registry")
	}
	return &GatewayService{
		servers: ss,
	}
}
