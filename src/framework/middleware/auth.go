package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			// 重定向到登录页
			// 检查是否是 AJAX 请求
			if c.GetHeader("X-Requested-With") == "XMLHttpRequest" {
				c.JSON(401, gin.H{
					"code":     401,
					"message":  "未登录",
					"redirect": "/login",
				})
			} else {
				c.Redirect(http.StatusFound, "/login")
			}
			return
		}
		c.Next()
	}
}
