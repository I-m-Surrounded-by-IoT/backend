package web

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/service/web/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/panjf2000/ants/v2"
	log "github.com/sirupsen/logrus"
)

var (
	ErrAuthFailed  = errors.New("auth failed")
	ErrAuthExpired = errors.New("auth expired")
)

type AuthClaims struct {
	UserId      string `json:"u"`
	UserVersion uint32 `json:"v"`
	jwt.RegisteredClaims
}

func (ws *WebService) authUser(Authorization string) (*AuthClaims, error) {
	t, err := jwt.ParseWithClaims(strings.TrimPrefix(Authorization, `Bearer `), &AuthClaims{}, func(token *jwt.Token) (any, error) {
		return ws.jwt.secret, nil
	})
	if err != nil {
		return nil, ErrAuthFailed
	}
	claims, ok := t.Claims.(*AuthClaims)
	if !ok || !t.Valid {
		return nil, ErrAuthFailed
	}
	return claims, nil
}

func (h *WebService) AuthUser(ctx context.Context, Authorization string) (*user.UserInfo, error) {
	claims, err := h.authUser(Authorization)
	if err != nil {
		return nil, err
	}

	if len(claims.UserId) != 32 {
		return nil, ErrAuthFailed
	}

	i, err := h.userClient.GetUserPasswordVersion(ctx, &user.GetUserPasswordVersionReq{
		Id: claims.UserId,
	})
	if err != nil {
		return nil, err
	}

	if i.Version != claims.UserVersion {
		return nil, ErrAuthExpired
	}

	return h.userClient.GetUserInfo(ctx, &user.GetUserInfoReq{
		Id: claims.UserId,
	})
}

func (ws *WebService) NewUserAuthToken(ctx context.Context, ID string) (string, error) {
	version, err := ws.userClient.GetUserPasswordVersion(ctx, &user.GetUserPasswordVersionReq{
		Id: ID,
	})
	if err != nil {
		return "", err
	}
	claims := &AuthClaims{
		UserId:      ID,
		UserVersion: version.Version,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ws.jwt.expire)),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(ws.jwt.secret)
}

func (h *WebService) AuthUserMiddleware(ctx *gin.Context) {
	token, err := GetAuthorizationTokenFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewApiErrorResp(err))
		return
	}
	userInfo, err := h.AuthUser(ctx, token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.NewApiErrorResp(err))
		return
	}

	if userInfo.Status.IsInActive() {
		ctx.AbortWithStatusJSON(http.StatusForbidden, model.NewApiErrorStringResp("user is inactive"))
		return
	}

	_ = ants.Submit(func() {
		_, _ = h.userClient.UpdateUserLastSeen(ctx, &user.UpdateUserLastSeenReq{
			Id: userInfo.Id,
			LastSeen: &user.UserLastSeen{
				LastSeenAt: time.Now().UnixMilli(),
				LastSeenIp: ctx.ClientIP(),
			},
		})
	})

	ctx.Set("user", userInfo)
	le := ctx.MustGet("log").(*log.Entry)
	if le.Data == nil {
		le.Data = make(log.Fields, 3)
	}
	le.Data["uid"] = userInfo.Id
	le.Data["unm"] = userInfo.Username
	le.Data["uro"] = userInfo.Role.String()
}

func (h *WebService) AuthAdminMiddleware(ctx *gin.Context) {
	h.AuthUserMiddleware(ctx)
	if ctx.IsAborted() {
		return
	}

	user := ctx.MustGet("user").(*user.UserInfo)
	if !user.Role.IsAdmin() {
		ctx.AbortWithStatusJSON(http.StatusForbidden, model.NewApiErrorStringResp("user is not admin"))
		return
	}
}

func GetAuthorizationTokenFromContext(ctx *gin.Context) (string, error) {
	Authorization := ctx.GetHeader("Authorization")
	if Authorization != "" {
		return Authorization, nil
	}
	Authorization = ctx.Query("token")
	if Authorization != "" {
		return Authorization, nil
	}
	return "", errors.New("token is empty")
}
