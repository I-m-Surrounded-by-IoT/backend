package web

import (
	"fmt"
	"net/http"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/service/web/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (ws *WebService) CreateUser(ctx *gin.Context) {
	req := model.CreateUserReq{}
	if err := model.Decode(ctx, &req); err != nil {
		log.Errorf("decode error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(fmt.Errorf("decode error: %v", err)))
		return
	}

	info, err := ws.uclient.CreateUser(ctx, (*user.CreateUserReq)(&req))
	if err != nil {
		log.Errorf("create user error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewApiDataResp(info))
}

func (ws *WebService) ListUser(ctx *gin.Context) {
	req := model.ListUserReq{}
	if err := model.Decode(ctx, &req); err != nil {
		log.Errorf("decode error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(fmt.Errorf("decode error: %v", err)))
		return
	}

	list, err := ws.uclient.ListUser(ctx, (*user.ListUserReq)(&req))
	if err != nil {
		log.Errorf("list user error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewApiDataResp(list))
}
