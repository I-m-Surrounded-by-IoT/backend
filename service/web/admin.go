package web

import (
	"net/http"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/service/web/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (ws *WebService) CreateUser(ctx *gin.Context) {
	req := model.CreateUserReq{}
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

	info, err := ws.userClient.CreateUser(ctx, (*user.CreateUserReq)(&req))
	if err != nil {
		log.Errorf("create user error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewApiDataResp(info))
}

func (ws *WebService) ListUser(ctx *gin.Context) {
	req := model.ListUserReq{}
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		log.Errorf("bind query error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	if err := req.Validate(); err != nil {
		log.Errorf("validate error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	list, err := ws.userClient.ListUser(ctx, (*user.ListUserReq)(&req))
	if err != nil {
		log.Errorf("list user error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewApiDataResp(list))
}
