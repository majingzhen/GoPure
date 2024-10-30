package routers

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api"
	"matuto.com/GoPure/src/framework/middleware"
)

// InitSystemRouter 初始化系统路由
func InitSystemRouter(r *gin.Engine) {
	// 登录相关
	r.GET("/login", api.System.JumpLoginView)
	r.GET("/captcha", api.System.CaptchaImage)
	r.POST("/doLogin", api.System.Login)
	r.GET("/logout", api.System.Logout)

	// 分组
	loginGroup := r.Group("/").Use(middleware.AuthMiddleware())
	{
		loginGroup.GET("/getLoginUser", api.System.GetLoginUser)

		// 首页
		loginGroup.GET("/", api.System.JumpHomeView)
		loginGroup.GET("/welcome", api.System.JumpWelcomeView)
	}
}
