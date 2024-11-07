package routers

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api"
	"matuto.com/GoPure/src/framework/middleware"
)

var Option = new(OptionRouter)

type OptionRouter struct{}

func (r *OptionRouter) InitOptionRoutes(e *gin.Engine) {
	options := e.Group("option").Use(middleware.AuthMiddleware()).Use(middleware.PermissionMiddleware())
	{
		options.GET("/page", api.Option.Page)
		options.GET("/list", api.Option.List)
		options.POST("/add", api.Option.Add)
		options.POST("/update", api.Option.Update)
		options.DELETE("/delete/:id", api.Option.Delete)
		options.GET("/get/:id", api.Option.Get)

		// 页面路由
		options.GET("/index", func(c *gin.Context) {
			c.HTML(200, "option/index.html", nil)
		})
		options.GET("/add", func(c *gin.Context) {
			c.HTML(200, "option/add.html", nil)
		})
		options.GET("/edit", func(c *gin.Context) {
			c.HTML(200, "option/edit.html", nil)
		})
	}

}
