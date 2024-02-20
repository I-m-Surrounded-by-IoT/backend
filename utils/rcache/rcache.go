package rcache

import (
	"math/rand"
	"time"

	redsync "github.com/go-redsync/redsync/v4"
	goredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

type Rcache struct {
	*redis.Client
	*redsync.Redsync
}

func NewRcache(rdb *redis.Client) *Rcache {
	return &Rcache{
		Client:  rdb,
		Redsync: redsync.New(goredis.NewPool(rdb)),
	}
}

func NewRcacheWithRsync(rdb *redis.Client, rsync *redsync.Redsync) *Rcache {
	return &Rcache{
		Client:  rdb,
		Redsync: rsync,
	}
}

func randExpireDuration(t time.Duration, maxRand time.Duration) time.Duration {
	if maxRand == 0 {
		return t
	}
	return t + time.Duration(rand.Int63n(int64(maxRand)))
}
