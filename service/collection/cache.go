package collection

import (
	"context"
	"fmt"
	"strconv"

	"github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/api/waterquality"
	"github.com/I-m-Surrounded-by-IoT/backend/utils/rcache"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
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

func (rc *CollectionRcache) UpdatePredictQuality(ctx context.Context, deviceID uint64, data *waterquality.PredictAndGuessResp) error {
	b, err := proto.Marshal(data)
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

func (rc *CollectionRcache) GetPredictQuality(ctx context.Context, deviceID uint64) (*waterquality.PredictAndGuessResp, error) {
	b, err := rc.rcache.Get(
		ctx,
		fmt.Sprintf("predict:quality:%d", deviceID),
	).Bytes()
	if err != nil {
		return nil, err
	}
	data := &waterquality.PredictAndGuessResp{}
	err = proto.Unmarshal(b, data)
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

func (rc *CollectionRcache) UpdateDeviceLastReport(ctx context.Context, id uint64, lastlocal *collection.DeviceLastReport) error {
	b, err := proto.Marshal(lastlocal)
	if err != nil {
		return err
	}
	return updateDeviceLastReportScript.Run(
		ctx,
		rc.rcache,
		[]string{fmt.Sprintf("device:last:report:%d", id)},
		strconv.FormatInt(lastlocal.ReceivedAt, 10),
		b,
	).Err()
}

func (rc *CollectionRcache) GetDeviceLastReport(ctx context.Context, id uint64) (*collection.DeviceLastReport, error) {
	resp, err := rc.rcache.HMGet(
		ctx,
		fmt.Sprintf("device:last:report:%d", id),
		"at",
		"data",
	).Result()
	if err != nil {
		return nil, err
	}
	if len(resp) != 2 {
		return nil, redis.Nil
	}
	lastlocal := &collection.DeviceLastReport{}
	lastlocal.ReceivedAt, err = strconv.ParseInt(resp[0].(string), 10, 64)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal([]byte(resp[1].(string)), lastlocal.Data)
	if err != nil {
		return nil, err
	}
	return lastlocal, nil
}
