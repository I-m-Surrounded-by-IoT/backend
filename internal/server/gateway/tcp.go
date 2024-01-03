package gateway

import (
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewTCPServer(
	config *conf.TcpServer,
	gatewayService utils.TCPHandler,
) *utils.TcpServer {
	ggs := utils.NewTcpServer(config, gatewayService)
	return ggs
}
