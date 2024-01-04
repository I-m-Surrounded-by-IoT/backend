//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package database

import (
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	reg "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/internal/server/database"
	service "github.com/I-m-Surrounded-by-IoT/backend/service/database"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(*conf.GrpcServer, *conf.Registry, *conf.DatabaseConfig, *conf.KafkaConfig, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(database.ProviderSet, service.ProviderSet, reg.ProviderSet, newApp))
}
