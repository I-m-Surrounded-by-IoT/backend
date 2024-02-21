package model

import (
	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
)

type CreateUserReq user.CreateUserReq

func (r *CreateUserReq) Validate() error {
	return nil
}

func (r *CreateUserReq) Decode(ctx *gin.Context) error {
	return json.NewDecoder(ctx.Request.Body).Decode(r)
}

type ListUserReq user.ListUserReq

func (r *ListUserReq) Validate() error {
	return nil
}

func (r *ListUserReq) Decode(ctx *gin.Context) error {
	return json.NewDecoder(ctx.Request.Body).Decode(r)
}
