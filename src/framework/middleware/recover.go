package middleware

import "github.com/gin-gonic/gin"

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.JSON(500, gin.H{
					"code": 500,
					"msg":  "服务器内部错误",
				})
				return
			}
		}()
		c.Next()
	}
}
