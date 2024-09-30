package routers

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api"
	"matuto.com/GoPure/src/framework/middleware"
)

type SystemRouter struct {
	systemApi api.SystemAPI
}

func (r *SystemRouter) InitSystemRouter(router *gin.RouterGroup) {
	systemRouter := router.Group("")
	{
		systemRouter.GET("/", r.systemApi.JumpHomeView)
		systemRouter.GET("/login", r.systemApi.JumpLoginView)
		systemRouter.POST("/doLogin", r.systemApi.Login)
		systemRouter.GET("/captcha", r.systemApi.CaptchaImage)
		systemRouter.GET("/logout", r.systemApi.Logout)
	}
	loginRouter := router.Group("").Use(middleware.AuthMiddleware())
	{
		loginRouter.GET("/getLoginUser", r.systemApi.GetLoginUser)
	}
}
