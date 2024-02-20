package rcache

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/zijiren233/stream"
	"google.golang.org/protobuf/proto"
)

type UserRcache struct {
	rc     *Rcache
	client user.UserClient
}

func NewUserRcache(rc *Rcache, client user.UserClient) *UserRcache {
	return &UserRcache{
		rc:     rc,
		client: client,
	}
}

func (uc *UserRcache) GetUserFromCache(ctx context.Context, id string) (*user.UserInfo, error) {
	resp := uc.rc.Get(ctx, fmt.Sprintf("user:%s", id))
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	if resp.Val() != "" {
		info := new(user.UserInfo)
		return info, proto.Unmarshal(stream.StringToBytes(resp.Val()), info)
	}
	return nil, errors.New("found but empty")
}

func (uc *UserRcache) SetUserToCache(ctx context.Context, id string, info *user.UserInfo) error {
	data, err := proto.Marshal(info)
	if err != nil {
		return err
	}
	resp := uc.rc.Set(ctx, fmt.Sprintf("user:%s", id), stream.BytesToString(data), randExpireDuration(time.Hour, 10*time.Minute))
	return resp.Err()
}

func (uc *UserRcache) GetUser(ctx context.Context, id string) (*user.UserInfo, error) {
	u, err := uc.GetUserFromCache(ctx, id)
	if err == nil {
		return u, nil
	}

	lock := uc.rc.NewMutex(fmt.Sprintf("mutex:user:%s", id))
	err = lock.Lock()
	if err != nil {
		return nil, err
	}
	defer func() {
		_, err := lock.Unlock()
		if err != nil {
			log.Errorf("failed to unlock mutex: %v", err)
		}
	}()

	u, err = uc.GetUserFromCache(ctx, id)
	if err == nil {
		return u, nil
	}

	info, err := uc.client.GetUser(ctx, &user.GetUserReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	err = uc.SetUserToCache(ctx, id, info)
	if err != nil {
		log.Errorf("failed to set user to cache: %v", err)
	}
	return info, nil
}

func (uc *UserRcache) GetPasswordVersion(ctx context.Context, id string) (int64, error) {
	resp := uc.rc.Get(ctx, fmt.Sprintf("password_version:%s", id))
	switch {
	case resp.Err() == redis.Nil:
		return 0, nil
	case resp.Err() != nil:
		return 0, resp.Err()
	}
	if resp.Val() != "" {
		return strconv.ParseInt(resp.Val(), 10, 64)
	}
	return 0, errors.New("found but empty")
}

func (uc *UserRcache) IncrPasswordVersion(ctx context.Context, id string) error {
	resp := uc.rc.Incr(ctx, fmt.Sprintf("password_version:%s", id))
	return resp.Err()
}
