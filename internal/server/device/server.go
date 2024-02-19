package device

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewDeviceServer,
)
