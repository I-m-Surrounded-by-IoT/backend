package database

import (
	databaseApi "github.com/I-m-Surrounded-by-IoT/backend/api/database"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/database"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
)

func NewDatabaseServer(
	config *conf.GrpcServer,
	db *database.DatabaseService,
) *utils.GrpcGatewayServer {
	ggs := utils.NewGrpcGatewayServer(config)
	databaseApi.RegisterDatabaseServer(ggs.GrpcRegistrar(), db)
	databaseApi.RegisterDatabaseHTTPServer(ggs.HttpRegistrar(), db)
	return ggs
}
