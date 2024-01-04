//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package collector

import (
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	reg "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/internal/server/collector"
	service "github.com/I-m-Surrounded-by-IoT/backend/service/collector"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(*conf.GrpcServer, *conf.TcpServer, *conf.Registry, *conf.CollectorConfig, *conf.KafkaConfig, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(collector.ProviderSet, service.ProviderSet, reg.ProviderSet, newApp))
}
