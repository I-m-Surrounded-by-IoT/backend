package email

import (
	emailApi "github.com/I-m-Surrounded-by-IoT/backend/api/email"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/email"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewEmailServer(
	config *conf.GrpcServerConfig,
	db *email.EmailService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	emailApi.RegisterEmailServer(ggs.GrpcRegistrar(), db)
	emailApi.RegisterEmailHTTPServer(ggs.HttpRegistrar(), db)
	return ggs
}
