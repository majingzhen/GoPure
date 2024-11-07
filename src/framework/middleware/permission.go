package middleware

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/service"
	"net/http"
)

// PermissionMiddleware 权限中间件
func PermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取当前请求路径
		path := c.Request.URL.Path
		// 从上下文获取用户信息
		_, exists := c.Get("user")
		if !exists {
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
		// 获取用户id
		userId := c.GetInt("userId")
		// 检查用户是否有权限
		if !checkPermission(userId, path) {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    403,
				"message": "无权限访问",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// checkPermission 检查用户是否有权限
func checkPermission(userId int, path string) bool {
	return service.User.CheckPermission(userId, path)
}
