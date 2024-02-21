package model

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ApiResp struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func (ar *ApiResp) SetError(err error) {
	ar.Error = err.Error()
}

func (ar *ApiResp) SetDate(data any) {
	ar.Data = data
}

func NewApiErrorResp(err error) *ApiResp {
	return &ApiResp{
		Error: err.Error(),
	}
}

func NewApiErrorStringResp(err string) *ApiResp {
	return &ApiResp{
		Error: err,
	}
}

func NewApiDataResp(data any) *ApiResp {
	return &ApiResp{
		Data: data,
	}
}

func GetPageAndSize(ctx *gin.Context) (page int, size int, err error) {
	size, err = strconv.Atoi(ctx.DefaultQuery("size", "10"))
	if err != nil {
		return 0, 0, errors.New("max must be a number")
	}
	page, err = strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		return 0, 0, errors.New("page must be a number")
	}
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	} else if size > 100 {
		size = 100
	}
	return
}
