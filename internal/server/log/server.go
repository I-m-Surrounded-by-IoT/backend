package log

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewLogServer,
	NewLogConsumer,
	NewDeviceLogServer,
)
