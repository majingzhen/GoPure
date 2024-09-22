package routers

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/routers"
	"matuto.com/GoPure/src/framework/middleware"
)

type Routers struct {
	userRouter routers.UserRouter
}

func (routers *Routers) InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.GinLogger())
	r.Use(gin.Recovery())

	// 跨域处理
	// 使用Cors中间件处理跨域请求
	r.Use(middleware.Cors())
	api := r.Group("/api")
	{
		routers.userRouter.InitUserRouter(api)
		//routers.baseRouter.InitBaseRouter(api)
		//routers.genRouter.InitGenRouter(api)
	}

	return r
}
