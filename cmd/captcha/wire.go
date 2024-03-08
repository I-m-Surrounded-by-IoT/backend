//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package captcha

import (
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	reg "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	server "github.com/I-m-Surrounded-by-IoT/backend/internal/server/captcha"
	service "github.com/I-m-Surrounded-by-IoT/backend/service/captcha"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(*conf.GrpcServerConfig, *conf.Registry, *conf.CaptchaConfig, *conf.RedisConfig, *conf.KafkaConfig, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, reg.ProviderSet, newApp))
}
