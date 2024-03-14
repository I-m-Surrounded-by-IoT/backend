package notify

import (
	notifyApi "github.com/I-m-Surrounded-by-IoT/backend/api/notify"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/notify"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewNotofyServer(
	config *conf.GrpcServerConfig,
	s *notify.NotifyService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	notifyApi.RegisterNotifyServer(ggs.GrpcRegistrar(), s)
	notifyApi.RegisterNotifyHTTPServer(ggs.HttpRegistrar(), s)
	return ggs
}
