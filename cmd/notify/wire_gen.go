// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package notify

import (
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	notify2 "github.com/I-m-Surrounded-by-IoT/backend/internal/server/notify"
	"github.com/I-m-Surrounded-by-IoT/backend/service/notify"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func wireApp(grpcServerConfig *conf.GrpcServerConfig, confRegistry *conf.Registry, notifyConfig *conf.NotifyConfig, kafkaConfig *conf.KafkaConfig, logger log.Logger) (*kratos.App, func(), error) {
	registrar := registry.NewRegistry(confRegistry)
	notifyService := notify.NewNotifyService(notifyConfig, kafkaConfig, registrar)
	grpcGatewayServer := notify2.NewNotofyServer(grpcServerConfig, notifyService)
	app := newApp(logger, grpcGatewayServer, registrar)
	return app, func() {
	}, nil
}
