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

	userinfo, err := ws.userClient.GetUserInfoByUsername(ctx, &user.GetUserInfoByUsernameReq{
		Username: req.Username,
	})
	if err != nil {
		log.Errorf("get user info by name error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(fmt.Errorf("get user info by name error: %v", err)))
		return
	}

	if userinfo.Status.IsInActive() {
		ctx.AbortWithStatusJSON(http.StatusForbidden, model.NewApiErrorResp(fmt.Errorf("user is inactive")))
		return
	}

	resp, err := ws.userClient.ValidateUserPassword(ctx, &user.ValidateUserPasswordReq{
		Id:       userinfo.Id,
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

	token, err := ws.NewUserAuthToken(ctx, userinfo.Id)
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

func (ws *WebService) SetUserPassword(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)
	userinfo := ctx.MustGet("user").(*user.UserInfo)

	req := model.SetUserPasswordReq{}
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

	_, err = ws.userClient.SetUserPassword(ctx, &user.SetUserPasswordReq{
		Id:       userinfo.Id,
		Password: req.Password,
	})
	if err != nil {
		log.Errorf("set user password error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	token, err := ws.NewUserAuthToken(ctx, userinfo.Id)
	if err != nil {
		log.Errorf("new user auth token error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(fmt.Errorf("new user auth token error: %v", err)))
		return
	}

	ctx.JSON(http.StatusOK, model.NewApiDataResp(gin.H{
		"token": token,
	}))
}

func (ws *WebService) SetUsername(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)
	userinfo := ctx.MustGet("user").(*user.UserInfo)

	req := model.SetUsernameReq{}
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

	_, err = ws.userClient.SetUsername(ctx, &user.SetUsernameReq{
		Id:       userinfo.Id,
		Username: req.Username,
	})
	if err != nil {
		log.Errorf("set user name error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
