package web

import (
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
	"github.com/gin-gonic/gin"
)

type WebService struct {
	eng *gin.Engine
}

func NewWebServer(c *conf.WebConfig) *WebService {
	return &WebService{}
}

func (ws *WebService) Init(eng *gin.Engine) {
	ws.eng = eng
}
