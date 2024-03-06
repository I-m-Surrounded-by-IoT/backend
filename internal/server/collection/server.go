package collection

import (
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewCollectionDatabase,
	NewCollectionConsumerServer,
	utils.ForceNewKafkaClient,
)
