package middlewares

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Init(e *gin.Engine) {
	w := log.StandardLogger().Writer()
	e.
		Use(gin.LoggerWithWriter(w), gin.RecoveryWithWriter(w)).
		Use(SetDateToHeader).
		Use(NewCors()).
		Use(NewLog(log.StandardLogger()))
	// if conf.Conf.RateLimit.Enable {
	// 	d, err := time.ParseDuration(conf.Conf.RateLimit.Period)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	options := []limiter.Option{
	// 		limiter.WithTrustForwardHeader(conf.Conf.RateLimit.TrustForwardHeader),
	// 	}
	// 	if conf.Conf.RateLimit.TrustedClientIPHeader != "" {
	// 		options = append(options, limiter.WithClientIPHeader(conf.Conf.RateLimit.TrustedClientIPHeader))
	// 	}
	// 	e.Use(NewLimiter(d, conf.Conf.RateLimit.Limit, options...))
	// }
}
