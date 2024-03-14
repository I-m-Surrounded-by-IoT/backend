package notify

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewNotofyServer,
)
