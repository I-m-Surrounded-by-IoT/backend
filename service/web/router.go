package web

import "github.com/gin-gonic/gin"

func (ws *WebService) RegisterRouter(e *gin.Engine) {
	api := e.Group("/api")

	needAuthUserApi := api.Group("", ws.AuthUserMiddleware)

	// needAuthAdminApi := api.Group("", ws.AuthAdminMiddleware)

	{
		userApi := api.Group("/user")

		needAuthUserApi := needAuthUserApi.Group("/user")

		ws.registerUser(userApi, needAuthUserApi)
	}
}

func (ws *WebService) registerUser(api, needAuthUserApi *gin.RouterGroup) {
	api.POST("/login", ws.Login)

	needAuthUserApi.GET("/me", ws.Me)
}
