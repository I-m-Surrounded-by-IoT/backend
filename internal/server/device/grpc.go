package device

import (
	deviceApi "github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/device"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewDeviceServer(
	config *conf.GrpcServerConfig,
	db *device.DeviceService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	deviceApi.RegisterDeviceServer(ggs.GrpcRegistrar(), db)
	deviceApi.RegisterDeviceHTTPServer(ggs.HttpRegistrar(), db)
	return ggs
}
