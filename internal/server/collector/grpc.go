package collector

import (
	collectorApi "github.com/I-m-Surrounded-by-IoT/backend/api/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewCollectorGrpcServer(
	config *conf.GrpcServerConfig,
	collector *collector.CollectorService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	collectorApi.RegisterCollectorServer(ggs.GrpcRegistrar(), collector)
	collectorApi.RegisterCollectorHTTPServer(ggs.HttpRegistrar(), collector)
	return ggs
}
