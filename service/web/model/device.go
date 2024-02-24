package model

import (
	"errors"

	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
)

var (
	ErrInvalidMac = errors.New("invalid mac")
)

type RegisterDeviceReq device.RegisterDeviceReq

func (c *RegisterDeviceReq) Validate() error {
	if len(c.Mac) != 17 {
		return ErrInvalidMac
	}
	return nil
}

type GetDeviceDetailReq struct {
	ID uint64 `form:"id" binding:"required"`
}

type GetDeviceDetailResp struct {
	*device.DeviceInfo
	*device.DeviceLastSeen
	*device.DeviceLastLocation
}
