package model

import (
	"github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
)

type GetDeviceDetailReq struct {
	ID uint64 `form:"id" binding:"required"`
}

type GetDeviceDetailResp struct {
	*device.DeviceInfo
	*device.DeviceLastSeen
	*collection.CollectionRecord
}

type DeviceIDReq struct {
	DeviceID uint64 `json:"deviceId" binding:"required"`
}
