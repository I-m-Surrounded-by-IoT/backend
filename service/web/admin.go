package web

import (
	"net/http"

	"github.com/I-m-Surrounded-by-IoT/backend/api/device"
	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/service/web/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (ws *WebService) CreateUser(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)

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
	log := ctx.MustGet("log").(*log.Entry)

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

func (ws *WebService) RegisterDevice(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)

	req := model.RegisterDeviceReq{}
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

	info, err := ws.deviceClient.RegisterDevice(ctx, (*device.RegisterDeviceReq)(&req))
	if err != nil {
		log.Errorf("create device error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.JSON(http.StatusOK, model.NewApiDataResp(info))
}

func (ws *WebService) SetUserStatus(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)

	req := user.SetUserStatusReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("bind json error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	userinfo, err := ws.userClient.GetUserInfo(ctx, &user.GetUserInfoReq{
		Id: req.Id,
		Fields: []string{
			"status",
			"role",
		},
	})
	if err != nil {
		log.Errorf("get user info error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	if userinfo.Role.IsAdmin() {
		log.Errorf("admin status can't be changed")
		ctx.AbortWithStatusJSON(http.StatusForbidden, model.NewApiErrorStringResp("admin status can't be changed"))
		return
	}

	if userinfo.Status == req.Status {
		ctx.Status(http.StatusNoContent)
		return
	}

	_, err = ws.userClient.SetUserStatus(ctx, &req)
	if err != nil {
		log.Errorf("set user status error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (ws *WebService) SetUserRole(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)

	req := user.SetUserRoleReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("bind json error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	userinfo, err := ws.userClient.GetUserInfo(ctx, &user.GetUserInfoReq{
		Id: req.Id,
		Fields: []string{
			"status",
			"role",
		},
	})
	if err != nil {
		log.Errorf("get user info error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	if userinfo.Role.IsAdmin() {
		log.Errorf("admin role can't be changed")
		ctx.AbortWithStatusJSON(http.StatusForbidden, model.NewApiErrorStringResp("admin role can't be changed"))
		return
	}

	if userinfo.Role == req.Role {
		ctx.Status(http.StatusNoContent)
		return
	}

	_, err = ws.userClient.SetUserRole(ctx, &req)
	if err != nil {
		log.Errorf("set user role error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (ws *WebService) AdminSetUsername(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)
	me := ctx.MustGet("user").(*user.UserInfo)

	req := user.SetUsernameReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("bind json error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	if me.Id != req.Id {
		userinfo, err := ws.userClient.GetUserInfo(ctx, &user.GetUserInfoReq{
			Id: req.Id,
			Fields: []string{
				"username",
				"role",
			},
		})
		if err != nil {
			log.Errorf("get user info error: %v", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
			return
		}

		if userinfo.Role.IsAdmin() {
			log.Errorf("admin role can't be changed")
			ctx.AbortWithStatusJSON(http.StatusForbidden, model.NewApiErrorStringResp("admin role can't be changed"))
			return
		}

		if userinfo.Username == req.Username {
			ctx.Status(http.StatusNoContent)
			return
		}
	} else {
		if me.Username == req.Username {
			ctx.Status(http.StatusNoContent)
			return
		}
	}

	_, err = ws.userClient.SetUsername(ctx, &req)
	if err != nil {
		log.Errorf("set user name error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (ws *WebService) AdminSetUserPassword(ctx *gin.Context) {
	log := ctx.MustGet("log").(*log.Entry)
	me := ctx.MustGet("user").(*user.UserInfo)

	req := user.SetUserPasswordReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Errorf("bind json error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorResp(err))
		return
	}

	if me.Id != req.Id {
		userinfo, err := ws.userClient.GetUserInfo(ctx, &user.GetUserInfoReq{
			Id: req.Id,
			Fields: []string{
				"role",
			},
		})
		if err != nil {
			log.Errorf("get user info error: %v", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
			return
		}

		if userinfo.Role.IsAdmin() {
			log.Errorf("admin role can't be changed")
			ctx.AbortWithStatusJSON(http.StatusForbidden, model.NewApiErrorStringResp("admin role can't be changed"))
			return
		}
	}

	_, err = ws.userClient.SetUserPassword(ctx, &req)
	if err != nil {
		log.Errorf("set user password error: %v", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.NewApiErrorResp(err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
