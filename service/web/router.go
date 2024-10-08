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

	{
		logApi := needAuthUserApi.Group("/log")

		ws.registerLog(logApi)
	}

	{
		collectionApi := needAuthUserApi.Group("/collection")

		ws.registerCollection(collectionApi)
	}
}

func (ws *WebService) registerUser(api, needAuthUserApi *gin.RouterGroup) {
	api.POST("/login", ws.Login)

	needAuthUserApi.GET("/me", ws.Me)

	needAuthUserApi.POST("/username", ws.SetUsername)

	needAuthUserApi.POST("/password", ws.SetUserPassword)

	needAuthUserApi.POST("/email/captcha", ws.SendBindEmailCaptcha)

	needAuthUserApi.POST("/email/bind", ws.BindEmail)

	needAuthUserApi.POST("/email/unbind", ws.UnbindEmail)

	needAuthUserApi.POST("/email/test", ws.SendTestEmail)

	needAuthUserApi.POST("/follow/device", ws.FollowDevice)

	needAuthUserApi.POST("/follow/device/all", ws.FollowAllDevice)

	needAuthUserApi.POST("/unfollow/device", ws.UnfollowDevice)

	needAuthUserApi.POST("/unfollow/device/all", ws.UnfollowAllDevice)

	needAuthUserApi.GET("/follow/device", ws.ListFollowedDevice)
}

func (ws *WebService) registerAdmin(adminApi *gin.RouterGroup) {
	{
		userApi := adminApi.Group("/user")

		userApi.POST("/create", ws.CreateUser)

		userApi.GET("/list", ws.ListUser)

		userApi.POST("/status", ws.SetUserStatus)

		userApi.POST("/role", ws.SetUserRole)

		userApi.POST("/username", ws.AdminSetUsername)

		userApi.POST("/password", ws.AdminSetUserPassword)
	}

	{
		deviceApi := adminApi.Group("/device")

		deviceApi.POST("/register", ws.RegisterDevice)

		deviceApi.POST("/password", ws.SetDevicePassword)

		deviceApi.GET("/event/stream", ws.GetDeviceStreamEvent)
	}
}

func (ws *WebService) registerDevice(deviceApi *gin.RouterGroup) {
	deviceApi.GET("/list", ws.ListDevice)

	deviceApi.GET("/detail", ws.GetDeviceDetail)

	// deviceApi.GET("/log/stream", ws.GetDeviceStreamLog)

	deviceApi.GET("/log", ws.ListDeviceLog)
}

func (ws *WebService) registerLog(logApi *gin.RouterGroup) {
	logApi.GET("/device", ws.ListDeviceLog)
}

func (ws *WebService) registerCollection(collectionApi *gin.RouterGroup) {
	collectionApi.GET("/records/device", ws.ListCollectionRecord)

	collectionApi.GET("/records/device/stream", ws.GetDeviceStreamReport)

	collectionApi.GET("/records/range/stream", ws.GetStreamLatestRecordsWithinRange)

	collectionApi.GET("/report_now", ws.ReportNow)

	collectionApi.GET("/boat_control", ws.BoatControl)
}
