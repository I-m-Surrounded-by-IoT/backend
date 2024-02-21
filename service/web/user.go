package web

import (
	"fmt"
	"net/http"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/service/web/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (ws *WebService) Login(ctx *gin.Context) {
	req := model.LoginUserReq{}

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

	ui, err := ws.userClient.GetUserInfoByName(ctx, &user.GetUserInfoByNameReq{
		Name: req.Username,
	})
	if err != nil {
		log.Errorf("get user info by name error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(fmt.Errorf("get user info by name error: %v", err)))
		return
	}

	if ui.Status.IsInActive() {
		ctx.AbortWithStatusJSON(http.StatusForbidden, model.NewApiErrorResp(fmt.Errorf("user is inactive")))
		return
	}

	resp, err := ws.userClient.ValidateUserPassword(ctx, &user.ValidateUserPasswordReq{
		Id:       ui.Id,
		Password: req.Password,
	})
	if err != nil {
		log.Errorf("validate user password error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}
	if !resp.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewApiErrorResp(fmt.Errorf("invalid password")))
		return
	}

	token, err := ws.NewUserAuthToken(ctx, ui.Id)
	if err != nil {
		log.Errorf("new user auth token error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(fmt.Errorf("new user auth token error: %v", err)))
		return
	}

	ctx.JSON(http.StatusOK, model.NewApiDataResp(gin.H{
		"token": token,
	}))
}

func (ws *WebService) Me(ctx *gin.Context) {
	uinfo := ctx.MustGet("user").(*user.UserInfo)

	ctx.JSON(http.StatusOK, model.NewApiDataResp(uinfo))
}
