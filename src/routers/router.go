package routers

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/app/routers"
	"matuto.com/GoPure/src/framework/middleware"
	"matuto.com/GoPure/src/global"
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

	// 使用中间件
	r.Use(middleware.GinLogger())
	r.Use(gin.Recovery())

	// 将User 类型 注册到gob中，允许在session中存储User类型
	gob.Register(model.User{})
	store := cookie.NewStore([]byte(global.Viper.GetString("session.secret")))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: global.Viper.GetInt("session.expire"),
	})
	r.Use(sessions.Sessions("GoPure-session", store))

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
