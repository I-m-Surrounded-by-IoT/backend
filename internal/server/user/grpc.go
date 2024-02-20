package user

import (
	userApi "github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/user"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewUserServer(
	config *conf.GrpcServerConfig,
	db *user.UserService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	userApi.RegisterUserServer(ggs.GrpcRegistrar(), db)
	userApi.RegisterUserHTTPServer(ggs.HttpRegistrar(), db)
	return ggs
}
