package routers

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/routers"
	"matuto.com/GoPure/src/framework/middleware"
)

type Routers struct {
	systemRouter routers.SystemRouter
	userRouter   routers.UserRouter
}

func (routers *Routers) InitRouter() *gin.Engine {
	r := gin.New()

	// 前端文件
	r.LoadHTMLGlob("ui/templates/*")
	r.Static("/static", "ui/static")
	r.Use(middleware.GinLogger())
	r.Use(gin.Recovery())

	// 跨域处理
	// 使用Cors中间件处理跨域请求
	r.Use(middleware.Cors())
	api := r.Group("/")
	{
		routers.systemRouter.InitSystemRouter(api)
		routers.userRouter.InitUserRouter(api)
		//routers.baseRouter.InitBaseRouter(api)
		//routers.genRouter.InitGenRouter(api)
	}

	return r
}
