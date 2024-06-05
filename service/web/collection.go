package web

import (
	"net/http"

	"github.com/I-m-Surrounded-by-IoT/backend/api/collection"
	"github.com/I-m-Surrounded-by-IoT/backend/service/web/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (ws *WebService) ListCollectionRecord(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)

	req := collection.ListCollectionRecordReq{}
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		log.Errorf("bind query error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	list, err := ws.collectionClient.ListCollectionRecord(ctx, &req)
	if err != nil {
		log.Errorf("list collection error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewApiDataResp(list))
}

func (ws *WebService) GetStreamLatestRecordsWithinRange(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)

	req := collection.GetStreamLatestWithinRangeReq{}
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		log.Errorf("bind query error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	stream, err := ws.collectionClient.GetStreamLatestRecordsWithinRange(ctx, &req)
	if err != nil {
		log.Errorf("get stream latest records within range error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	defer func() { _ = stream.CloseSend() }()

	for {
		select {
		case <-stream.Context().Done():
			ctx.SSEvent("stop", "finish")
			return
		default:
			resp, err := stream.Recv()
			if err != nil {
				log.Errorf("get stream latest records within range error: %v", err)
				ctx.SSEvent("stop", err)
				return
			}
			ctx.SSEvent("message", resp)
			if err := ctx.Errors.Last(); err != nil {
				log.Errorf("get stream latest records within range error: %v", err)
				ctx.SSEvent("stop", err)
				return
			}
			ctx.Writer.Flush()
		}
	}
}
