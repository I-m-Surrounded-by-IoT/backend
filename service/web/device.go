package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/I-m-Surrounded-by-IoT/backend/api/collector"
	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/service/web/model"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
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
	resp.LastSeen = lastSeen.LastSeen

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

func (ws *WebService) GetDeviceStreamLog(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)

	req := collector.GetDeviceStreamLogReq{}
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		log.Errorf("bind query error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}
	si, err := ws.etcd.GetService(ctx, fmt.Sprintf("device-%v", req.Id))
	if err != nil || len(si) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorStringResp("device is not online"))
		return
	}

	discoveryDeviceConn, err := utils.NewDiscoveryGrpcConn(context.Background(), &utils.Backend{
		Endpoint: fmt.Sprintf("discovery:///device-%d", req.Id),
	}, ws.etcd)
	if err != nil {
		log.Fatalf("failed to create grpc conn: %v", err)
	}
	cli := collector.NewCollectorClient(discoveryDeviceConn)
	c, err := cli.GetDeviceStreamLog(ctx, &req)
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
				log.Errorf("get device stream log error: %v", err)
				ctx.SSEvent("error", model.NewApiErrorResp(err))
				return
			}
			ctx.SSEvent("log", resp)
		}
	}
}
