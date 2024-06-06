package collection

import (
	"context"
	"fmt"
	"strconv"

	"github.com/I-m-Surrounded-by-IoT/backend/service/collection/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/rcache"
	json "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"github.com/zijiren233/stream"
)

type CollectionRcache struct {
	rcache *rcache.Rcache
	db     *dbUtils
}

func NewCollectionRcache(rcache *rcache.Rcache, client *dbUtils) *CollectionRcache {
	return &CollectionRcache{
		rcache: rcache,
		db:     client,
	}
}

func (rc *CollectionRcache) UpdateLastPredictQuality(ctx context.Context, deviceID uint64, data *model.PredictAndGuess) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return rc.rcache.Set(
		ctx,
		fmt.Sprintf("predict:quality:%d", deviceID),
		b,
		redis.KeepTTL,
	).Err()
}

func (rc *CollectionRcache) GetLastPredictQualityFromCache(ctx context.Context, deviceID uint64) (*model.PredictAndGuess, error) {
	b, err := rc.rcache.Get(
		ctx,
		fmt.Sprintf("predict:quality:%d", deviceID),
	).Bytes()
	if err != nil {
		return nil, err
	}
	data := &model.PredictAndGuess{}
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (rc *CollectionRcache) GetLastPredictQuality(ctx context.Context, deviceID uint64) (*model.PredictAndGuess, error) {
	resp, err := rc.GetLastPredictQualityFromCache(ctx, deviceID)
	if err == nil {
		return resp, nil
	}
	if err != redis.Nil {
		log.Errorf("failed to get predict quality from cache: %v", err)
	}

	lock := rc.rcache.NewMutex(fmt.Sprintf("mutex:predict:quality:%d", deviceID))
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

	resp, err = rc.GetLastPredictQualityFromCache(ctx, deviceID)
	if err == nil {
		return resp, nil
	}
	if err != redis.Nil {
		log.Errorf("failed to get predict quality from cache: %v", err)
	}

	data, err := rc.db.GetDeviceLastPredictAndGuess(deviceID)
	if err != nil {
		return nil, err
	}
	err = rc.UpdateLastPredictQuality(ctx, deviceID, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

var updateDeviceLastReportScript = redis.NewScript(`
local key = KEYS[1]
local at = ARGV[1]
local data = ARGV[2]
local last = redis.call('HGET', key, 'at')
if last == false or last < at then
	redis.call('HMSET', key, 'at', at, 'data', data)
end
return 0
`)

func (rc *CollectionRcache) UpdateDeviceLastReport(ctx context.Context, id uint64, lastlocal *model.CollectionRecord) error {
	b, err := json.Marshal(lastlocal.CollectionData)
	if err != nil {
		return err
	}
	return updateDeviceLastReportScript.Run(
		ctx,
		rc.rcache,
		[]string{fmt.Sprintf("device:last:report:%d", id)},
		strconv.FormatInt(lastlocal.ReceivedAt.UnixMilli(), 10),
		b,
	).Err()
}

func (rc *CollectionRcache) GetDeviceLastReportFromCache(ctx context.Context, id uint64) (*model.CollectionRecord, error) {
	resp, err := rc.rcache.HMGet(
		ctx,
		fmt.Sprintf("device:last:report:%d", id),
		"at",
		"data",
	).Result()
	if err == nil {
		if len(resp) != 2 {
			return nil, redis.Nil
		}
		lastlocal := &model.CollectionData{}
		data, ok := resp[1].(string)
		if !ok {
			return nil, fmt.Errorf("failed to convert resp proto data to string")
		}
		err = json.Unmarshal(stream.StringToBytes(data), lastlocal)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal device last report: %w", err)
		}
		return &model.CollectionRecord{
			DeviceID:       id,
			CollectionData: lastlocal,
		}, nil
	}
	return nil, err
}

func (rc *CollectionRcache) GetDeviceLastReport(ctx context.Context, id uint64) (*model.CollectionRecord, error) {
	resp, err := rc.GetDeviceLastReportFromCache(ctx, id)
	if err == nil {
		return resp, nil
	}
	if err != redis.Nil {
		log.Errorf("failed to get device last report from cache: %v", err)
	}

	lock := rc.rcache.NewMutex(fmt.Sprintf("mutex:device:last:report:%d", id))
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
	resp, err = rc.GetDeviceLastReportFromCache(ctx, id)
	if err == nil {
		return resp, nil
	}
	if err != redis.Nil {
		log.Errorf("failed to get device last report from cache: %v", err)
	}

	resp, err = rc.db.GetDeviceLastReport(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get device last report from db: %w", err)
	}

	err = rc.UpdateDeviceLastReport(ctx, id, resp)
	if err != nil {
		log.Errorf("failed to update device last report to cache: %v", err)
	}
	return resp, nil
}
