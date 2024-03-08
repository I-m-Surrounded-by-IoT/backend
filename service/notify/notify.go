package notify

import (
	"github.com/I-m-Surrounded-by-IoT/backend/api/user"
	"github.com/I-m-Surrounded-by-IoT/backend/conf"
)

type NotifyService struct {
	userClient user.UserClient
}

func NewNotifyService(dc *conf.DatabaseServerConfig, lc *conf.LogConfig) *NotifyService {
	return &NotifyService{}
}
