package web

import (
	"fmt"
	"net/http"

	"github.com/I-m-Surrounded-by-IoT/backend/api/captcha"
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

func (ws *WebService) SendBindEmailCaptcha(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)
	userinfo := ctx.MustGet("user").(*user.UserInfo)

	req := model.SendBindEmailCaptchaReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("bind json error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	_, err = ws.captchaClient.SendEmailCaptcha(ctx, &captcha.SendEmailCaptchaReq{
		UserId: userinfo.Id,
		Email:  req.Email,
	})
	if err != nil {
		log.Errorf("get bind email captcha error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (ws *WebService) BindEmail(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)
	userinfo := ctx.MustGet("user").(*user.UserInfo)

	req := model.BindEmailReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("bind json error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	_, err = ws.captchaClient.VerifyEmailCaptcha(ctx, &captcha.VerifyEmailCaptchaReq{
		UserId:  userinfo.Id,
		Email:   req.Email,
		Captcha: req.Captcha,
	})
	if err != nil {
		log.Errorf("verify email captcha error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	_, err = ws.userClient.BindEmail(ctx, &user.BindEmailReq{
		Id:    userinfo.Id,
		Email: req.Email,
	})
	if err != nil {
		log.Errorf("bind email error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (ws *WebService) UnbindEmail(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)
	userinfo := ctx.MustGet("user").(*user.UserInfo)

	_, err := ws.userClient.UnbindEmail(ctx, &user.UnbindEmailReq{
		Id: userinfo.Id,
	})
	if err != nil {
		log.Errorf("unbind email error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (ws *WebService) FollowDevice(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)
	userinfo := ctx.MustGet("user").(*user.UserInfo)

	req := model.DeviceIDReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("bind json error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	_, err = ws.userClient.FollowDevice(ctx, &user.FollowDeviceReq{
		UserId:   userinfo.Id,
		DeviceId: req.DeviceID,
	})
	if err != nil {
		log.Errorf("follow device error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (ws *WebService) UnfollowDevice(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)
	userinfo := ctx.MustGet("user").(*user.UserInfo)

	req := model.DeviceIDReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("bind json error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	_, err = ws.userClient.UnfollowDevice(ctx, &user.UnfollowDeviceReq{
		UserId:   userinfo.Id,
		DeviceId: req.DeviceID,
	})
	if err != nil {
		log.Errorf("unfollow device error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

// TODO: return device info, add query params
func (ws *WebService) ListFollowedDevice(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)
	userinfo := ctx.MustGet("user").(*user.UserInfo)

	resp, err := ws.userClient.ListFollowedDeviceIDs(ctx, &user.ListFollowedDeviceIDsReq{
		UserId: userinfo.Id,
	})
	if err != nil {
		log.Errorf("list followed device error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	// add has followed all device field
	ctx.JSON(http.StatusOK, model.NewApiDataResp(resp))
}

func (ws *WebService) FollowAllDevice(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)
	userinfo := ctx.MustGet("user").(*user.UserInfo)

	_, err := ws.userClient.FollowAllDevice(ctx, &user.FollowAllDeviceReq{
		UserId: userinfo.Id,
	})
	if err != nil {
		log.Errorf("follow all device error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (ws *WebService) UnfollowAllDevice(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)
	userinfo := ctx.MustGet("user").(*user.UserInfo)

	_, err := ws.userClient.UnfollowAllDevice(ctx, &user.UnfollowAllDeviceReq{
		UserId: userinfo.Id,
	})
	if err != nil {
		log.Errorf("unfollow all device error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
