package routers

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api"
	"matuto.com/GoPure/src/framework/middleware"
)

var User = new(UserRouter)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(e *gin.Engine) {
	userGroup := e.Group("user").Use(middleware.AuthMiddleware())
	{
		// 页面路由
		userGroup.GET("/", api.User.JumpUserView)
		userGroup.GET("/add", api.User.JumpUserAddView)
		userGroup.GET("/edit", api.User.JumpUserEditView)

		// 数据接口
		userGroup.GET("/page", api.User.Page)
		userGroup.GET("/get", api.User.Get)
		userGroup.POST("/add", api.User.Add)
		userGroup.POST("/update", api.User.Update)
		userGroup.POST("/delete", api.User.Delete)
		userGroup.POST("/updateStatus", api.User.UpdateStatus)
		userGroup.POST("/resetPassword", api.User.ResetPassword)
	}
}
