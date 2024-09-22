package routers

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api"
)

type UserRouter struct {
	userApi api.UserAPI
}

func (r *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")
	{
		userRouter.GET("/:id", r.userApi.GetUserById)
	}
}
