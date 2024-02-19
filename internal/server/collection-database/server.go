package collection_database

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewCollectionDatabase,
	NewLogConsumer,
	NewCollectionConsumerServer,
)
