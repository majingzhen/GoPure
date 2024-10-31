package routers

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api"
	"matuto.com/GoPure/src/framework/middleware"
)

var Menu = new(MenuRouter)

type MenuRouter struct{}

// InitMenuRouter 初始化菜单路由
func (r *MenuRouter) InitMenuRouter(e *gin.Engine) {
	menuRouter := e.Group("/menu").Use(middleware.AuthMiddleware())
	{
		// 页面路由
		menuRouter.GET("/", api.Menu.JumpMenuView)
		menuRouter.GET("/add", api.Menu.JumpMenuAddView)
		menuRouter.GET("/edit", api.Menu.JumpMenuEditView)

		// 数据接口
		menuRouter.GET("/list", api.Menu.List)
		menuRouter.GET("/get/:id", api.Menu.Get)
		menuRouter.POST("/add", api.Menu.Add)
		menuRouter.POST("/update", api.Menu.Update)
		menuRouter.POST("/delete", api.Menu.Delete)
	}
}
