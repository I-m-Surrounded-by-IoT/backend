package collection_database

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCollectionDatabase, NewCollectionConsumer)
