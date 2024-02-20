package collection

import (
	collectionApi "github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	collection "github.com/I-m-Surrounded-by-IoT/backend/service/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewCollectionDatabase(
	config *conf.GrpcServerConfig,
	db *collection.CollectionService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	collectionApi.RegisterCollectionServer(ggs.GrpcRegistrar(), db)
	collectionApi.RegisterCollectionHTTPServer(ggs.HttpRegistrar(), db)
	return ggs
}
