package web

import (
	"github.com/gin-gonic/gin"
)

func (ws *WebService) RegisterRouter(e *gin.Engine) {
	api := e.Group("/api")

	needAuthUserApi := api.Group("", ws.AuthUserMiddleware)

	// needAuthAdminApi := api.Group("", ws.AuthAdminMiddleware)

	{
		userApi := api.Group("/user")

		needAuthUserApi := needAuthUserApi.Group("/user")

		ws.registerUser(userApi, needAuthUserApi)
	}

	{
		adminApi := api.Group("/admin", ws.AuthAdminMiddleware)

		ws.registerAdmin(adminApi)
	}

	{
		deviceApi := needAuthUserApi.Group("/device")

		ws.registerDevice(deviceApi)
	}
}

func (ws *WebService) registerUser(api, needAuthUserApi *gin.RouterGroup) {
	api.POST("/login", ws.Login)

	needAuthUserApi.GET("/me", ws.Me)
}

func (ws *WebService) registerAdmin(adminApi *gin.RouterGroup) {
	{
		userApi := adminApi.Group("/user")

		userApi.POST("/create", ws.CreateUser)

		userApi.GET("/list", ws.ListUser)
	}

	{
		deviceApi := adminApi.Group("/device")

		deviceApi.POST("/register", ws.RegisterDevice)
	}
}

func (ws *WebService) registerDevice(deviceApi *gin.RouterGroup) {
	deviceApi.GET("/list", ws.ListDevice)

	deviceApi.GET("/detail", ws.GetDeviceDetail)
}
