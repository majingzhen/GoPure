package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto.com/GoPure/src/global"
	"time"
)

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 处理请求
		c.Next()
		// 计算请求处理时间
		latency := time.Since(start)
		// 获取相关信息
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		if method == "OPTIONS" { // 跨域请求会先发送一个OPTIONS请求，这里不做处理
			return
		}
		// 将日志输出到Zap
		global.Logger.Info("Gin request",
			zap.Int("status", statusCode),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("ip", clientIP),
			zap.Duration("latency", latency),
		)
	}
}
