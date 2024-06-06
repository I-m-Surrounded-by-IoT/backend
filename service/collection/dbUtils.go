package collection

import (
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/service/collection/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"gorm.io/gorm"
)

type dbUtils struct {
	db *gorm.DB
}

func newDBUtils(db *gorm.DB) *dbUtils {
	return &dbUtils{db: db}
}

func (u *dbUtils) GetCollectionRecord(deviceID uint64, scopes ...func(*gorm.DB) *gorm.DB) ([]*model.CollectionRecord, error) {
	return u.ListCollectionRecord(append(scopes, utils.WithDeviceIDEq(deviceID))...)
}

func (u *dbUtils) ListCollectionRecord(scopes ...func(*gorm.DB) *gorm.DB) ([]*model.CollectionRecord, error) {
	var collections []*model.CollectionRecord
	err := u.db.Scopes(scopes...).Find(&collections).Error
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (u *dbUtils) CountCollectionRecord(scopes ...func(*gorm.DB) *gorm.DB) (int64, error) {
	var count int64
	err := u.db.Model(&model.CollectionRecord{}).Scopes(scopes...).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *dbUtils) CreateCollectionRecord(collection *model.CollectionRecord) error {
	return u.db.Create(collection).Error
}

func (u *dbUtils) UpdateCollectionRecordLevel(deviceID uint64, level int64) error {
	return u.db.Model(&model.CollectionRecord{}).Where("device_id = ?", deviceID).Update("level", level).Error
}

func (u *dbUtils) GetDeviceIDsWithinRange(centerLat, centerLon, radiusMeters float64, before, after time.Time) ([]uint64, error) {
	var deviceIDs []uint64
	err := u.db.Table("collection_records").
		Select("DISTINCT ON (device_id) device_id").
		Joins("JOIN (SELECT device_id, MAX(timestamp) as latest_timestamp FROM collection_records WHERE (created_at > ? AND created_at < ?) OR (updated_at > ? AND updated_at < ?) GROUP BY device_id) latest ON collection_records.device_id = latest.device_id AND collection_records.timestamp = latest.latest_timestamp", after, before, after, before).
		Where("ST_DWithin(geo_point, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography, ?)", centerLon, centerLat, radiusMeters).
		Scan(&deviceIDs).Error
	if err != nil {
		return nil, err
	}
	return deviceIDs, nil
}

func (u *dbUtils) GetLatestRecordsWithinRange(centerLat, centerLon, radiusMeters float64, before, after time.Time) ([]*model.CollectionRecord, error) {
	var records []*model.CollectionRecord
	err := u.db.Table("collection_records").
		Select("collection_records.*").
		Joins("JOIN (SELECT device_id, MAX(timestamp) as latest_timestamp FROM collection_records WHERE (created_at > ? AND created_at < ?) OR (updated_at > ? AND updated_at < ?) GROUP BY device_id) latest ON collection_records.device_id = latest.device_id AND collection_records.timestamp = latest.latest_timestamp", after, before, after, before).
		Where("ST_DWithin(geo_point, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography, ?)", centerLon, centerLat, radiusMeters).
		Preload("PredictAndGuess").
		Scan(&records).Error
	if err != nil {
		return nil, err
	}
	return records, err
}

func (u *dbUtils) GetIDsNotWithinRange(ids []uint64, centerLat, centerLon, radiusMeters float64, after time.Time) ([]uint64, error) {
	var result []uint64
	err := u.db.Table("collection_records").
		Select("collection_records.device_id").
		Joins("JOIN (SELECT device_id, MAX(timestamp) as latest_timestamp FROM collection_records WHERE (created_at > ? OR updated_at > ?) AND device_id IN ? GROUP BY device_id) latest ON collection_records.device_id = latest.device_id AND collection_records.timestamp = latest.latest_timestamp", after, after, ids).
		Where("NOT ST_DWithin(geo_point, ST_SetSRID(ST_MakePoint(?, ?), 4326)::geography, ?)", centerLon, centerLat, radiusMeters).
		Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (u *dbUtils) CreateOrUpdatePredictAndGuess(predictAndGuess *model.PredictAndGuess) error {
	return u.db.Save(predictAndGuess).Error
}

func (u *dbUtils) GetPredictAndGuess(recordID uint64) (*model.PredictAndGuess, error) {
	var predictAndGuess model.PredictAndGuess
	err := u.db.Where("collection_record_id = ?", recordID).First(&predictAndGuess).Error
	if err != nil {
		return nil, err
	}
	return &predictAndGuess, nil
}

func (u *dbUtils) GetDeviceLastPredictAndGuess(deviceID uint64) (*model.PredictAndGuess, error) {
	var predictAndGuess model.PredictAndGuess
	err := u.db.Where("device_id = ?", deviceID).Order("created_at DESC").First(&predictAndGuess).Error
	if err != nil {
		return nil, err
	}
	return &predictAndGuess, nil
}

func (u *dbUtils) GetDeviceLastReport(deviceID uint64) (*model.CollectionRecord, error) {
	var record model.CollectionRecord
	err := u.db.
		Where("device_id = ?", deviceID).
		Order("timestamp DESC").
		Preload("PredictAndGuess").
		First(&record).
		Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}
