package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// GinLogger 定义日志中间件
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 请求路径
		path := c.Request.URL.Path
		// 请求参数
		c.Next()

		cost := time.Since(start)

		if c.Writer.Status() != http.StatusOK {
			zap.L().Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.Duration("cost", cost),
				zap.String("path", path),
			)
		}
	}
}
