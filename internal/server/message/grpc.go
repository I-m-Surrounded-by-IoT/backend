package message

import (
	messageApi "github.com/I-m-Surrounded-by-IoT/backend/api/message"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/message"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewMessageGrpcServer(
	config *conf.GrpcServerConfig,
	message *message.MessageService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	messageApi.RegisterMessageServer(ggs.GrpcRegistrar(), message)
	messageApi.RegisterMessageHTTPServer(ggs.HttpRegistrar(), message)
	return ggs
}
