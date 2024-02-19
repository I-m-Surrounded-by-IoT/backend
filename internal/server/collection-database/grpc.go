package collection_database

import (
	collection_databaseApi "github.com/I-m-Surrounded-by-IoT/backend/api/collection-database"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	collection_database "github.com/I-m-Surrounded-by-IoT/backend/service/collection-database"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewCollectionDatabase(
	config *conf.GrpcServer,
	db *collection_database.CollectionDatabaseService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	collection_databaseApi.RegisterCollectionDatabaseServer(ggs.GrpcRegistrar(), db)
	collection_databaseApi.RegisterCollectionDatabaseHTTPServer(ggs.HttpRegistrar(), db)
	return ggs
}
