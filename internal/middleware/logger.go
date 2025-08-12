package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ZapLogger(lgr *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		lgr.Info("INFO",
			zap.String("path", path),
			zap.String("method", method),
			zap.Int("status", c.Writer.Status()),
		)
	}
}
