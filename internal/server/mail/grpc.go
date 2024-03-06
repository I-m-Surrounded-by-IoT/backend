package mail

import (
	mailApi "github.com/I-m-Surrounded-by-IoT/backend/api/mail"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/mail"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewMailServer(
	config *conf.GrpcServerConfig,
	db *mail.MailService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	mailApi.RegisterMailServer(ggs.GrpcRegistrar(), db)
	mailApi.RegisterMailHTTPServer(ggs.HttpRegistrar(), db)
	return ggs
}
