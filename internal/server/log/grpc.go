package log

import (
	logApi "github.com/I-m-Surrounded-by-IoT/backend/api/log"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/log"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewLogServer(
	config *conf.GrpcServerConfig,
	lg *log.LogService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	logApi.RegisterLogServer(ggs.GrpcRegistrar(), lg)
	logApi.RegisterLogHTTPServer(ggs.HttpRegistrar(), lg)
	return ggs
}
