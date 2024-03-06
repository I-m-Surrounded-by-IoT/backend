package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/service/web/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (ws *WebService) ListDevice(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)

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
	resp.DeviceLastSeen = lastSeen

	lastReport, err := ws.deviceClient.GetDeviceLastReport(ctx, &device.GetDeviceLastReportReq{
		Id: id,
	})
	if err != nil {
		return nil, fmt.Errorf("get device last location error: %v", err)
	}
	resp.DeviceLastReport = lastReport

	return resp, nil
}

func (ws *WebService) GetDeviceDetail(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)

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

func (ws *WebService) GetDeviceStreamReport(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)

	req := collection.GetDeviceStreamReportReq{}
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		log.Errorf("bind query error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}
	c, err := ws.collectionClient.GetDeviceStreamReport(ctx, &req)
	if err != nil {
		log.Errorf("get device stream log error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}
	defer func() { _ = c.CloseSend() }()
	for {
		select {
		case <-ctx.Request.Context().Done():
			return
		case <-c.Context().Done():
			ctx.SSEvent("close", nil)
			return
		default:
			resp, err := c.Recv()
			if err != nil {
				log.Errorf("get device stream repoty error: %v", err)
				ctx.SSEvent("error", err)
				return
			}
			ctx.SSEvent("report", resp)
		}
	}
}
