// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package collector

import (
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	collector2 "github.com/I-m-Surrounded-by-IoT/backend/internal/server/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/service/collector"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func wireApp(grpcServerConfig *conf.GrpcServerConfig, tcpServer *conf.TcpServer, confRegistry *conf.Registry, collectorConfig *conf.CollectorConfig, kafkaConfig *conf.KafkaConfig, logger log.Logger) (*kratos.App, func(), error) {
	registrar := registry.NewRegistry(confRegistry)
	collectorService := collector.NewCollectorService(collectorConfig, kafkaConfig, registrar)
	grpcGatewayServer := collector2.NewCollectorGrpcServer(grpcServerConfig, collectorService)
	utilsTcpServer := collector2.NewTcpServer(tcpServer, collectorService)
	app := newApp(logger, grpcGatewayServer, utilsTcpServer, registrar)
	return app, func() {
	}, nil
}
