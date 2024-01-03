package collector

import (
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewTcpServer(
	config *conf.TcpServer,
	collectorService *collector.CollectorService,
) *utils.TcpServer {
	ggs := utils.NewTcpServer(config, collectorService)
	return ggs
}
