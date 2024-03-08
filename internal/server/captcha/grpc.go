package captcha

import (
	captchaApi "github.com/I-m-Surrounded-by-IoT/backend/api/captcha"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/captcha"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewCaptchaServer(
	config *conf.GrpcServerConfig,
	lg *captcha.CaptchaService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	captchaApi.RegisterCaptchaServer(ggs.GrpcRegistrar(), lg)
	captchaApi.RegisterCaptchaHTTPServer(ggs.HttpRegistrar(), lg)
	return ggs
}
