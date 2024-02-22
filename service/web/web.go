package web

import (
	"context"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	logApi "github.com/I-m-Surrounded-by-IoT/backend/api/log"
	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	registryClient "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/service/web/middlewares"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type jwtConfig struct {
	secret []byte
	expire time.Duration
}

type WebService struct {
	config           *conf.WebConfig
	jwt              *jwtConfig
	rdb              *redis.Client
	etcd             *registryClient.EtcdRegistry
	userClient       user.UserClient
	deviceClient     device.DeviceClient
	logClient        logApi.LogClient
	collectionClient collection.CollectionClient
}

func NewWebServer(c *conf.WebConfig, reg registry.Registrar, rc *conf.RedisConfig) *WebService {
	etcd := reg.(*registryClient.EtcdRegistry)
	discoveryUserConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///user",
		TimeOut:  "10s",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}
	userClient := user.NewUserClient(discoveryUserConn)

	discoveryDeviceConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///device",
		TimeOut:  "10s",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}
	deviceClient := device.NewDeviceClient(discoveryDeviceConn)

	discoveryLogConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///log",
		TimeOut:  "10s",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}
	logClient := logApi.NewLogClient(discoveryLogConn)

	discoveryCollectionConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///collection",
		TimeOut:  "10s",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}
	collectionClient := collection.NewCollectionClient(discoveryCollectionConn)

	jwtExpire, err := time.ParseDuration(c.Jwt.Expire)
	if err != nil {
		log.Fatalf("failed to parse jwt expire: %v", err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Username: rc.Username,
		Password: rc.Password,
		DB:       int(rc.Db),
	})

	return &WebService{
		config: c,
		jwt: &jwtConfig{
			secret: []byte(c.Jwt.Secret),
			expire: jwtExpire,
		},
		rdb:              rdb,
		etcd:             etcd,
		userClient:       userClient,
		deviceClient:     deviceClient,
		logClient:        logClient,
		collectionClient: collectionClient,
	}
}

func (ws *WebService) Init(eng *gin.Engine) {
	middlewares.Init(eng)
	ws.RegisterRouter(eng)
}
