package log

import (
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewLogServer,
	utils.ForceNewKafkaClient,
	NewConsumerGroup,
	NewDeviceLogServer,
)
