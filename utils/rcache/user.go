package rcache

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type UserRcache struct {
	rdb    *Rcache
	client user.UserClient
}

func NewUserRcache(rdb *Rcache, client user.UserClient) *UserRcache {
	return &UserRcache{
		rdb:    rdb,
		client: client,
	}
}

func (uc *UserRcache) GetUserInfoFromCache(ctx context.Context, id string) (*user.UserInfo, error) {
	info := new(user.UserInfo)
	resp := uc.rdb.HGetAll(ctx, fmt.Sprintf("userinfo:%s", id))
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	if len(resp.Val()) == 0 {
		return nil, redis.Nil
	}
	err := resp.Scan(info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (uc *UserRcache) SetUserInfoToCache(ctx context.Context, id string, info *user.UserInfo) error {
	return uc.rdb.HSet(ctx, fmt.Sprintf("userinfo:%s", id), info).Err()
}

func (uc *UserRcache) DelUserInfoCache(ctx context.Context, id string) error {
	return uc.rdb.Del(ctx, fmt.Sprintf("userinfo:%s", id)).Err()
}

func (uc *UserRcache) GetUserInfo(ctx context.Context, id string) (*user.UserInfo, error) {
	u, err := uc.GetUserInfoFromCache(ctx, id)
	if err == nil {
		return u, nil
	}

	lock := uc.rdb.NewMutex(fmt.Sprintf("mutex:userinfo:%s", id))
	err = lock.LockContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to lock mutex: %w", err)
	}
	defer func() {
		_, err := lock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock mutex: %v", err)
		}
	}()

	u, err = uc.GetUserInfoFromCache(ctx, id)
	if err == nil {
		return u, nil
	}
	log.Warnf("failed to get user info from cache: %v", err)

	info, err := uc.client.GetUser(ctx, &user.GetUserReq{
		Id: id,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get user from db: %w", err)
	}
	err = uc.SetUserInfoToCache(ctx, id, info)
	if err != nil {
		return nil, fmt.Errorf("failed to set user info to cache: %w", err)
	}
	return info, nil
}

func (uc *UserRcache) GetUserPasswordVersion(ctx context.Context, id string) (int64, error) {
	resp := uc.rdb.HGet(ctx, "user_passwold_version", id)
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

func (uc *UserRcache) IncrUserPasswordVersion(ctx context.Context, id string) error {
	return uc.rdb.HIncrBy(ctx, "user_passwold_version", id, 1).Err()
}

func (uc *UserRcache) GetUserIDFromCache(ctx context.Context, name string) (string, error) {
	resp := uc.rdb.HGet(ctx, "username_to_id", name)
	return resp.Val(), resp.Err()
}

func (uc *UserRcache) SetUserIDToCache(ctx context.Context, name, id string) error {
	return uc.rdb.HSet(ctx, "username_to_id", name, id).Err()
}

func (uc *UserRcache) DelUserIDCache(ctx context.Context, name string) error {
	return uc.rdb.HDel(ctx, "username_to_id", name).Err()
}

func (uc *UserRcache) GetUserID(ctx context.Context, name string) (string, error) {
	id, err := uc.GetUserIDFromCache(ctx, name)
	if err == nil {
		return id, nil
	}

	lock := uc.rdb.NewMutex(fmt.Sprintf("mutex:username:%s", name))
	err = lock.LockContext(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to lock mutex: %w", err)
	}
	defer func() {
		_, err := lock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock mutex: %v", err)
		}
	}()

	id, err = uc.GetUserIDFromCache(ctx, name)
	if err == nil {
		return id, nil
	}
	log.Warnf("failed to get user id from cache: %v", err)

	info, err := uc.client.GetUserByName(ctx, &user.GetUserByNameReq{
		Name: name,
		Fields: []string{
			"id",
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to get user id from db: %w", err)
	}
	err = uc.SetUserIDToCache(ctx, name, info.Id)
	if err != nil {
		return "", fmt.Errorf("failed to set user id to cache: %w", err)
	}
	return info.Id, nil
}

func (uc *UserRcache) GetUserInfoByName(ctx context.Context, name string) (*user.UserInfo, error) {
	id, err := uc.GetUserID(ctx, name)
	if err != nil {
		return nil, err
	}
	return uc.GetUserInfo(ctx, id)
}
