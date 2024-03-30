package message

import (
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewMessageGrpcServer,
	utils.ForceNewKafkaClient,
	NewConsumerGroup,
	NewMessageMQServerServer,
)
