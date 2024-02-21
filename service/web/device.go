package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/service/web/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (ws *WebService) ListDevice(ctx *gin.Context) {
	req := device.ListDeviceReq{}

	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		log.Errorf("bind query error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	list, err := ws.deviceClient.ListDevice(ctx, &req)
	if err != nil {
		log.Errorf("list device error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewApiDataResp(list))
}

func (ws *WebService) CreateDevice(ctx *gin.Context) {
	req := model.CreateDeviceReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("bind json error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	if err := req.Validate(); err != nil {
		log.Errorf("validate error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	info, err := ws.deviceClient.CreateDevice(ctx, (*device.CreateDeviceReq)(&req))
	if err != nil {
		log.Errorf("create device error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewApiDataResp(info))
}

func (ws *WebService) GenDeviceDetail(ctx context.Context, id uint64) (*model.GetDeviceDetailResp, error) {
	deviceInfo, err := ws.deviceClient.GetDeviceInfo(ctx, &device.GetDeviceInfoReq{
		Id: id,
	})
	if err != nil {
		return nil, fmt.Errorf("get device info error: %v", err)
	}
	resp := &model.GetDeviceDetailResp{
		DeviceInfo: deviceInfo,
	}
	lastSeen, err := ws.deviceClient.GetDeviceLastSeen(ctx, &device.GetDeviceLastSeenReq{
		Id: id,
	})
	if err != nil {
		return nil, fmt.Errorf("get device last seen error: %v", err)
	}
	resp.LastSeen = lastSeen.LastSeen

	return resp, nil
}

func (ws *WebService) GetDeviceDetail(ctx *gin.Context) {
	req := model.GetDeviceDetailReq{}
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		log.Errorf("bind query error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	info, err := ws.GenDeviceDetail(ctx, req.ID)
	if err != nil {
		log.Errorf("get device detail error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewApiDataResp(info))
}
