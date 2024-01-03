//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package gateway

import (
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	reg "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/internal/server/gateway"
	service "github.com/I-m-Surrounded-by-IoT/backend/service/gateway"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(*conf.TcpServer, *conf.Registry, *conf.GatewayConfig, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(gateway.ProviderSet, service.ProviderSet, reg.ProviderSet, newApp))
}
