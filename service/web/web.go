package web

import (
	"context"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	registryClient "github.com/I-m-Surrounded-by-IoT/backend/internal/registry"
	"github.com/I-m-Surrounded-by-IoT/backend/service/web/middlewares"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/rcache"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/registry"
	redsync "github.com/go-redsync/redsync/v4"
	goredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type jwtConfig struct {
	secret []byte
	expire time.Duration
}

type WebService struct {
	config  *conf.WebConfig
	jwt     *jwtConfig
	rdb     *redis.Client
	rsync   *redsync.Redsync
	uclient user.UserClient
	ucache  *rcache.UserRcache
}

func NewWebServer(c *conf.WebConfig, reg registry.Registrar, rc *conf.RedisConfig) *WebService {
	etcd := reg.(*registryClient.EtcdRegistry)
	conn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: "discovery:///user",
	}, etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}

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
	rsync := redsync.New(goredis.NewPool(rdb))
	cache := rcache.NewRcacheWithRsync(rdb, rsync)
	uclient := user.NewUserClient(conn)

	return &WebService{
		config: c,
		jwt: &jwtConfig{
			secret: []byte(c.Jwt.Secret),
			expire: jwtExpire,
		},
		rdb:     rdb,
		rsync:   rsync,
		uclient: uclient,
		ucache:  rcache.NewUserRcache(cache, uclient),
	}
}

func (ws *WebService) Init(eng *gin.Engine) {
	middlewares.Init(eng)
	ws.RegisterRouter(eng)
}
