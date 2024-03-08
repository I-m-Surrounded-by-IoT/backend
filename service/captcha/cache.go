package captcha

import (
	"context"
	"fmt"
	"time"

	"github.com/I-m-Surrounded-by-IoT/backend/utils"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type CaptchaRcache struct {
	rcli *redis.Client
}

func NewCaptchaRcache(rcli *redis.Client) *CaptchaRcache {
	return &CaptchaRcache{
		rcli: rcli,
	}
}

func (c *CaptchaRcache) NewMailCaptcha(ctx context.Context, userid, mail string) (string, error) {
	captcha := utils.GetRandString(6)
	resp := c.rcli.Set(ctx, fmt.Sprintf("captcha:mail:%s:%s:%s", userid, mail, captcha), "1", time.Minute*5)
	if resp.Err() != nil {
		return "", resp.Err()
	}
	return captcha, nil
}

var ErrCaptchaNotMatch = fmt.Errorf("captcha not match")

func (c *CaptchaRcache) VerifyEmailCaptcha(ctx context.Context, userid, mail, captcha string) error {
	resp, err := c.rcli.GetDel(ctx, fmt.Sprintf("captcha:mail:%s:%s:%s", userid, mail, captcha)).Result()
	if err != nil {
		if err != redis.Nil {
			log.Errorf("get email captcha error: %v", err)
		}
		return ErrCaptchaNotMatch
	}
	if resp != "1" {
		return ErrCaptchaNotMatch
	}
	return nil
}
