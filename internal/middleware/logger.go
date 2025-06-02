package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		latency := time.Since(start)
		log.WithFields(log.Fields{
			"status":     ctx.Writer.Status(),
			"method":     ctx.Request.Method,
			"path":       ctx.Request.URL.Path,
			"latency":    latency,
			"client_ip":  ctx.ClientIP(),
			"user_agent": ctx.Request.UserAgent(),
		}).Info("HTTP request log")
	}
}
