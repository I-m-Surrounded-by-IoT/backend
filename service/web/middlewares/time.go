package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
)

func SetDateToHeader(ctx *gin.Context) {
	ctx.Header("Date", time.Now().Format(time.RFC1123Z))
	ctx.Next()
}
