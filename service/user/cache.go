package user

import (
	"context"
	"fmt"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/rcache"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type UserRcache struct {
	rcache *rcache.Rcache
	db     *dbUtils
}

func NewUserRcache(rcache *rcache.Rcache, client *dbUtils) *UserRcache {
	return &UserRcache{
		rcache: rcache,
		db:     client,
	}
}

func (uc *UserRcache) GetUserInfoFromCache(ctx context.Context, id string, fields ...string) (*user.UserInfo, error) {
	info := new(user.UserInfo)
	if len(fields) == 0 {
		resp := uc.rcache.HGetAll(ctx, fmt.Sprintf("user:info:%s", id))
		if resp.Err() != nil {
			return nil, resp.Err()
		}
		if len(resp.Val()) == 0 {
			return nil, redis.Nil
		}
		return info, resp.Scan(info)
	} else {
		resp := uc.rcache.HMGet(ctx, fmt.Sprintf("user:info:%s", id), fields...)
		if resp.Err() != nil {
			return nil, resp.Err()
		}
		if len(resp.Val()) == 0 {
			return nil, redis.Nil
		}
		return info, resp.Scan(info)
	}
}

func (uc *UserRcache) SetUserInfoToCache(ctx context.Context, id string, info *user.UserInfo) error {
	return uc.rcache.HSet(ctx, fmt.Sprintf("user:info:%s", id), info).Err()
}

func (uc *UserRcache) DelUserInfoCache(ctx context.Context, id string) error {
	return uc.rcache.Del(ctx, fmt.Sprintf("user:info:%s", id)).Err()
}

func (uc *UserRcache) GetUserInfo(ctx context.Context, id string, fields ...string) (*user.UserInfo, error) {
	u, err := uc.GetUserInfoFromCache(ctx, id, fields...)
	if err == nil {
		return u, nil
	}
	if err != redis.Nil {
		log.Errorf("failed to get user info from cache: %v", err)
	}

	lock := uc.rcache.NewMutex(fmt.Sprintf("mutex:user:info:%s", id))
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

	u, err = uc.GetUserInfoFromCache(ctx, id, fields...)
	if err == nil {
		return u, nil
	}
	if err != redis.Nil {
		log.Errorf("failed to get user info from cache: %v", err)
	}

	dbLock := uc.rcache.NewMutex(fmt.Sprintf("mutex:db:user:info:%s", id))
	err = dbLock.LockContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to lock db mutex: %w", err)
	}
	defer func() {
		_, err := dbLock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock db mutex: %v", err)
		}
	}()

	info, err := uc.db.GetUser(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from db: %w", err)
	}
	pbInfo := user2Proto(info)
	err = uc.SetUserInfoToCache(ctx, id, pbInfo)
	if err != nil {
		log.Errorf("failed to set user info to cache: %v", err)
	}
	return pbInfo, nil
}

func (uc *UserRcache) GetUserPasswordVersionFromCache(ctx context.Context, id string) (uint64, error) {
	resp := uc.rcache.HGet(ctx, "user_password_version", id)
	return resp.Uint64()
}

func (uc *UserRcache) SetUserPasswordVersionToCache(ctx context.Context, id string, version uint32) error {
	return uc.rcache.HSet(ctx, "user_password_version", id, version).Err()
}

func (uc *UserRcache) DelUserPasswordVersionCache(ctx context.Context, id string) error {
	return uc.rcache.HDel(ctx, "user_password_version", id).Err()
}

func (uc *UserRcache) GetUserPasswordVersion(ctx context.Context, id string) (uint32, error) {
	i, err := uc.GetUserPasswordVersionFromCache(ctx, id)
	if err == nil {
		return uint32(i), nil
	}

	lock := uc.rcache.NewMutex(fmt.Sprintf("mutex:user_password:%s", id))
	err = lock.LockContext(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to lock mutex: %w", err)
	}
	defer func() {
		_, err := lock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock mutex: %v", err)
		}
	}()

	i, err = uc.GetUserPasswordVersionFromCache(ctx, id)
	if err == nil {
		return uint32(i), nil
	}

	dbLock := uc.rcache.NewMutex(fmt.Sprintf("mutex:db:user_password:%s", id))
	err = dbLock.LockContext(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to lock db mutex: %w", err)
	}
	defer func() {
		_, err := dbLock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock db mutex: %v", err)
		}
	}()

	info, err := uc.db.GetUserPasswordVersion(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("failed to get user password version from db: %w", err)
	}
	err = uc.SetUserPasswordVersionToCache(ctx, id, info)
	if err != nil {
		log.Errorf("failed to set user password version to cache: %v", err)
	}
	return info, nil
}

func (uc *UserRcache) SetUserPassword(ctx context.Context, id string, password string) error {
	lock := uc.rcache.NewMutex(fmt.Sprintf("mutex:db:user_password:%s", id))
	err := lock.LockContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to lock db mutex: %w", err)
	}
	defer func() {
		_, err := lock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock db mutex: %v", err)
		}
	}()

	err = uc.db.SetUserPassword(ctx, id, password)
	if err != nil {
		return err
	}

	err = uc.DelUserPasswordVersionCache(ctx, id)
	if err != nil {
		log.Errorf("failed to del user password version cache: %v", err)
	}
	return nil
}

func (uc *UserRcache) GetUserIDFromCache(ctx context.Context, name string) (string, error) {
	resp := uc.rcache.HGet(ctx, "username_to_id", name)
	return resp.Val(), resp.Err()
}

func (uc *UserRcache) SetUserIDToCache(ctx context.Context, name, id string) error {
	return uc.rcache.HSet(ctx, "username_to_id", name, id).Err()
}

func (uc *UserRcache) DelUserIDCache(ctx context.Context, name string) error {
	return uc.rcache.HDel(ctx, "username_to_id", name).Err()
}

func (uc *UserRcache) GetUserID(ctx context.Context, name string) (string, error) {
	id, err := uc.GetUserIDFromCache(ctx, name)
	if err == nil {
		return id, nil
	}
	if err != redis.Nil {
		log.Errorf("failed to get user info from cache: %v", err)
	}

	lock := uc.rcache.NewMutex(fmt.Sprintf("mutex:username:%s", name))
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
	if err != redis.Nil {
		log.Errorf("failed to get user info from cache: %v", err)
	}

	info, err := uc.db.GetUserByName(ctx, name, "id")
	if err != nil {
		return "", fmt.Errorf("failed to get user id from db: %w", err)
	}

	dbLock := uc.rcache.NewMutex(fmt.Sprintf("mutex:db:user:info:%s", info.ID))
	err = dbLock.LockContext(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to lock db mutex: %w", err)
	}
	defer func() {
		_, err := dbLock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock db mutex: %v", err)
		}
	}()

	userInfo, err := uc.GetUserInfoFromCache(ctx, info.ID)
	switch {
	case err == redis.Nil:
		u, err := uc.db.GetUser(ctx, info.ID, "username")
		if err != nil {
			return "", fmt.Errorf("failed to get user from db: %w", err)
		}
		userInfo = user2Proto(u)
	case err != nil:
		return "", fmt.Errorf("failed to get user info from cache: %w", err)
	}

	if userInfo.Username != name {
		return "", fmt.Errorf("user name changed, try again")
	}

	err = uc.SetUserIDToCache(ctx, name, info.ID)
	if err != nil {
		log.Errorf("failed to set user id to cache: %v", err)
	}
	return info.ID, nil
}

func (uc *UserRcache) GetUserInfoByUsername(ctx context.Context, name string, fields ...string) (*user.UserInfo, error) {
	id, err := uc.GetUserID(ctx, name)
	if err != nil {
		return nil, err
	}
	return uc.GetUserInfo(ctx, id, fields...)
}

func (uc *UserRcache) SetUsername(ctx context.Context, id, name string) (string, error) {
	dbLock := uc.rcache.NewMutex(fmt.Sprintf("mutex:db:user:info:%s", id))
	err := dbLock.LockContext(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to lock db mutex: %w", err)
	}
	defer func() {
		_, err := dbLock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock db mutex: %v", err)
		}
	}()

	old, err := uc.db.SetUsername(ctx, id, name)
	if err != nil {
		return "", err
	}

	err = uc.DelUserInfoCache(ctx, id)
	if err != nil {
		log.Errorf("failed to del user info cache: %v", err)
	}

	err = uc.DelUserIDCache(ctx, old)
	if err != nil {
		log.Errorf("failed to del user id cache: %v", err)
	}

	return old, nil
}

func (uc *UserRcache) SetUserRole(ctx context.Context, id string, role user.Role) error {
	lock := uc.rcache.NewMutex(fmt.Sprintf("mutex:db:user:info:%s", id))
	err := lock.LockContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to lock db mutex: %w", err)
	}
	defer func() {
		_, err := lock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock db mutex: %v", err)
		}
	}()

	err = uc.db.SetUserRole(ctx, id, role)
	if err != nil {
		return err
	}

	err = uc.DelUserInfoCache(ctx, id)
	if err != nil {
		log.Errorf("failed to del user info cache: %v", err)
	}
	return nil
}

func (uc *UserRcache) SetUserStatus(ctx context.Context, id string, status user.Status) error {
	lock := uc.rcache.NewMutex(fmt.Sprintf("mutex:db:user:info:%s", id))
	err := lock.LockContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to lock db mutex: %w", err)
	}
	defer func() {
		_, err := lock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock db mutex: %v", err)
		}
	}()

	err = uc.db.SetUserStatus(ctx, id, status)
	if err != nil {
		return err
	}

	err = uc.DelUserInfoCache(ctx, id)
	if err != nil {
		log.Errorf("failed to del user info cache: %v", err)
	}
	return nil
}

func (uc *UserRcache) UpdateUserLastSeen(ctx context.Context, id string, lastSeen *user.UserLastSeen) error {
	return uc.rcache.HSet(ctx, fmt.Sprintf("user:last:seen:%s", id), lastSeen).Err()
}

func (uc *UserRcache) GetUserLastSeen(ctx context.Context, id string) (*user.UserLastSeen, error) {
	resp := uc.rcache.HGetAll(ctx, fmt.Sprintf("user:last:seen:%s", id))
	if resp.Err() != nil {
		if resp.Err() == redis.Nil {
			return &user.UserLastSeen{}, nil
		}
		return nil, resp.Err()
	}
	var lastSeen user.UserLastSeen
	err := resp.Scan(&lastSeen)
	if err != nil {
		return nil, err
	}
	return &lastSeen, nil
}
