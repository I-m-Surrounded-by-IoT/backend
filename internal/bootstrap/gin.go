package bootstrap

import (
	"context"

	"github.com/I-m-Surrounded-by-IoT/backend/cmd/flags"
	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/gin-gonic/gin"
)

func InitGinMode(ctx context.Context) error {
	if flags.Dev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	if utils.ForceColor() {
		gin.ForceConsoleColor()
	} else {
		gin.DisableConsoleColor()
	}

	return nil
}
