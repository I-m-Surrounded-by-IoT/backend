package web

import (
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/I-m-Surrounded-by-IoT/backend/service/web"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewWebServer(
	config *conf.WebServerConfig,
	web *web.WebService,
) *http.Server {
	eng := gin.New()
	s := http.NewServer(
		http.Address(
			config.Addr,
		),
	)
	s.HandlePrefix("/", eng)
	return s
}
