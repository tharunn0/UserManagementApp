package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/tharunn0/gin-server-gorm/pkg/logger"
	"go.uber.org/zap"
)

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		log.Info("REQUEST",
			zap.String("path", path),
			zap.String("method", method),
			zap.Int("status", c.Writer.Status()),
		)
		c.Next()
	}
}
