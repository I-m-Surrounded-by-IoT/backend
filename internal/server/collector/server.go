package collector

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewCollectorGrpcServer, NewTcpServer)
