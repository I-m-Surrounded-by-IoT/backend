package collection

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewCollectionDatabase,
	NewLogConsumer,
	NewCollectionConsumerServer,
)
