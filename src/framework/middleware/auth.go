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
			// 检查是否是 AJAX 请求
			if c.GetHeader("X-Requested-With") == "XMLHttpRequest" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code":     401,
					"message":  "未登录",
					"redirect": "/login",
				})
			} else {
				// 不再检查 Sec-Fetch-Dest，统一返回 unauthorized.html
				c.HTML(http.StatusUnauthorized, "unauthorized.html", gin.H{
					"message":  "会话已过期，请重新登录",
					"loginUrl": "/login",
				})
			}
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Set("userId", session.Get("userId"))
		c.Next()
	}
}
