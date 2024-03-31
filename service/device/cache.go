package device

import (
	"context"
	"fmt"

	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/service/device/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/rcache"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type DeviceRcache struct {
	rcache *rcache.Rcache
	db     *dbUtils
}

func NewDeviceRcache(rcache *rcache.Rcache, client *dbUtils) *DeviceRcache {
	return &DeviceRcache{
		rcache: rcache,
		db:     client,
	}
}

func (dc *DeviceRcache) GetDeviceInfoFromCache(ctx context.Context, id uint64, fields ...string) (*device.DeviceInfo, error) {
	info := new(device.DeviceInfo)
	if len(fields) == 0 {
		resp := dc.rcache.HGetAll(ctx, fmt.Sprintf("device:info:%d", id))
		if resp.Err() != nil {
			return nil, resp.Err()
		}
		if len(resp.Val()) == 0 {
			return nil, redis.Nil
		}
		return info, resp.Scan(info)
	} else {
		resp := dc.rcache.HMGet(ctx, fmt.Sprintf("device:info:%d", id), fields...)
		if resp.Err() != nil {
			return nil, resp.Err()
		}
		if len(resp.Val()) == 0 {
			return nil, redis.Nil
		}
		return info, resp.Scan(info)
	}
}

func (dc *DeviceRcache) SetDeviceInfoToCache(ctx context.Context, id uint64, info *device.DeviceInfo) error {
	return dc.rcache.HSet(ctx, fmt.Sprintf("device:info:%d", id), info).Err()
}

func (dc *DeviceRcache) DelDeviceInfoCache(ctx context.Context, id uint64) error {
	return dc.rcache.Del(ctx, fmt.Sprintf("device:info:%d", id)).Err()
}

func (dc *DeviceRcache) GetDeviceInfo(ctx context.Context, id uint64, fields ...string) (*device.DeviceInfo, error) {
	u, err := dc.GetDeviceInfoFromCache(ctx, id, fields...)
	if err == nil {
		return u, nil
	}
	if err != redis.Nil {
		log.Errorf("failed to get device info from cache: %v", err)
	}

	lock := dc.rcache.NewMutex(fmt.Sprintf("mutex:device:info:%d", id))
	err = lock.LockContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to lock mutex: %v", err)
	}
	defer func() {
		_, err = lock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock mutex: %v", err)
		}
	}()

	u, err = dc.GetDeviceInfoFromCache(ctx, id, fields...)
	if err == nil {
		return u, nil
	}
	if err != redis.Nil {
		log.Errorf("failed to get device info from cache: %v", err)
	}

	dbLock := dc.rcache.NewMutex(fmt.Sprintf("mutex:db:device:info:%d", id))
	err = dbLock.LockContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to lock db mutex: %v", err)
	}
	defer func() {
		_, err = dbLock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock db mutex: %v", err)
		}
	}()

	info, err := dc.db.GetDeviceInfo(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get device info from database: %v", err)
	}

	deviceInfo := device2Proto(info)

	err = dc.SetDeviceInfoToCache(ctx, id, deviceInfo)
	if err != nil {
		log.Errorf("failed to set device info to cache: %v", err)
	}
	return deviceInfo, nil
}

func (dc *DeviceRcache) Transaction(fn func(dc *DeviceRcache) error) error {
	return dc.db.Transaction(func(db *dbUtils) error {
		return fn(NewDeviceRcache(dc.rcache, db))
	})
}

func (dc *DeviceRcache) DelDevice(ctx context.Context, id uint64) (*model.Device, error) {
	device, err := dc.db.DelDevice(ctx, id)
	if err != nil {
		return nil, err
	}
	err = dc.DelDeviceInfoCache(ctx, id)
	if err != nil {
		log.Errorf("failed to delete device info cache: %v", err)
	}
	err = dc.DelDeviceIDCache(ctx, device.Mac)
	if err != nil {
		log.Errorf("failed to delete device id cache: %v", err)
	}
	err = dc.DelDeviceExtraCache(ctx, id)
	if err != nil {
		log.Errorf("failed to delete device extra cache: %v", err)
	}
	return device, nil
}

func (dc *DeviceRcache) UndelDevice(ctx context.Context, id uint64) (*model.Device, error) {
	device, err := dc.db.UndelDevice(ctx, id)
	if err != nil {
		return nil, err
	}
	err = dc.DelDeviceInfoCache(ctx, id)
	if err != nil {
		log.Errorf("failed to delete device info cache: %v", err)
	}
	err = dc.DelDeviceIDCache(ctx, device.Mac)
	if err != nil {
		log.Errorf("failed to delete device id cache: %v", err)
	}
	return device, nil
}

func (dc *DeviceRcache) GetDeviceIDFromCache(ctx context.Context, mac string) (uint64, error) {
	resp := dc.rcache.Get(ctx, fmt.Sprintf("deviceid:%s", mac))
	if resp.Err() != nil {
		return 0, resp.Err()
	}
	if resp.Val() == "" {
		return 0, fmt.Errorf("device id found in cache, but it's empty")
	}
	return resp.Uint64()
}

func (dc *DeviceRcache) SetDeviceIDToCache(ctx context.Context, mac string, id uint64) error {
	return dc.rcache.Set(ctx, fmt.Sprintf("deviceid:%s", mac), id, 0).Err()
}

func (dc *DeviceRcache) DelDeviceIDCache(ctx context.Context, mac string) error {
	return dc.rcache.Del(ctx, fmt.Sprintf("deviceid:%s", mac)).Err()
}

func (dc *DeviceRcache) GetDeviceID(ctx context.Context, mac string) (uint64, error) {
	id, err := dc.GetDeviceIDFromCache(ctx, mac)
	if err == nil {
		return id, nil
	}
	if err != redis.Nil {
		log.Errorf("failed to get device id from cache: %v", err)
	}

	lock := dc.rcache.NewMutex(fmt.Sprintf("mutex:deviceid:%s", mac))
	err = lock.LockContext(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to lock mutex: %v", err)
	}
	defer func() {
		_, err = lock.UnlockContext(ctx)
		if err != nil {
			log.Errorf("failed to unlock mutex: %v", err)
		}
	}()

	id, err = dc.GetDeviceIDFromCache(ctx, mac)
	if err == nil {
		return id, nil
	}
	if err != redis.Nil {
		log.Errorf("failed to get device id from cache: %v", err)
	}

	info, err := dc.db.GetDeviceInfoWithMac(ctx, mac, "id")
	if err != nil {
		return 0, fmt.Errorf("failed to get device id from database: %v", err)
	}

	err = dc.SetDeviceIDToCache(ctx, mac, info.ID)
	if err != nil {
		log.Errorf("failed to set device id to cache: %v", err)
	}
	return info.ID, nil
}

func (dc *DeviceRcache) GetDeviceInfoByMac(ctx context.Context, mac string, fields ...string) (*device.DeviceInfo, error) {
	id, err := dc.GetDeviceID(ctx, mac)
	if err != nil {
		return nil, err
	}
	return dc.GetDeviceInfo(ctx, id, fields...)
}

func (dc *DeviceRcache) DelDeviceExtraCache(ctx context.Context, id uint64) error {
	return dc.rcache.Del(ctx, fmt.Sprintf("device:last:seen:%d", id)).Err()
}

var updateDeviceLastReportScript = redis.NewScript(`
local key = KEYS[1]
local at = ARGV[1]
local ip = ARGV[2]

local last = redis.call('HGET', key, 'at')
if last == false or last < at then
	redis.call('HMSET', key, 'at', at, 'ip', ip)
end
return 0
`)

var updateDeviceLastReportWithoutIpScript = redis.NewScript(`
local key = KEYS[1]
local at = ARGV[1]

local last = redis.call('HGET', key, 'at')
if last == false or last < at then
	redis.call('HSET', key, 'at', at)
end
return 0
`)

func (dc *DeviceRcache) UpdateDeviceLastSeen(ctx context.Context, id uint64, lastSeen *device.DeviceLastSeen) error {
	if lastSeen.LastSeenIp == "" {
		return updateDeviceLastReportWithoutIpScript.Run(
			ctx,
			dc.rcache,
			[]string{fmt.Sprintf("device:last:seen:%d", id)},
			lastSeen.LastSeenAt,
		).Err()
	}

	return updateDeviceLastReportScript.Run(
		ctx,
		dc.rcache,
		[]string{fmt.Sprintf("device:last:seen:%d", id)},
		lastSeen.LastSeenAt,
		lastSeen.LastSeenIp,
	).Err()
}

func (dc *DeviceRcache) GetDeviceLastSeen(ctx context.Context, id uint64) (*device.DeviceLastSeen, error) {
	resp := dc.rcache.HGetAll(ctx, fmt.Sprintf("device:last:seen:%d", id))
	if resp.Err() != nil {
		if resp.Err() == redis.Nil {
			return &device.DeviceLastSeen{}, nil
		}
		return nil, resp.Err()
	}
	lastSeen := &device.DeviceLastSeen{}
	err := resp.Scan(lastSeen)
	if err != nil {
		return nil, err
	}
	return lastSeen, nil
}
