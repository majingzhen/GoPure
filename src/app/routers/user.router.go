package routers

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api"
	"matuto.com/GoPure/src/framework/middleware"
)

type UserRouter struct {
	userApi api.UserAPI
}

func (r *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user").Use(middleware.AuthMiddleware())
	{
		userRouter.GET("/page", r.userApi.Page)
		userRouter.GET("/:id", r.userApi.GetUserById)
	}
}
