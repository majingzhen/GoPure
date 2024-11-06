package routers

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api"
	"matuto.com/GoPure/src/framework/middleware"
)

var Dict = new(DictRouter)

type DictRouter struct{}

func (r *DictRouter) InitDictRouter(e *gin.Engine) {
	dictGroup := e.Group("dict").Use(middleware.AuthMiddleware())
	{
		// 页面路由
		dictGroup.GET("/", func(c *gin.Context) {
			c.HTML(200, "dict/index.html", nil)
		})
		dictGroup.GET("/add", func(c *gin.Context) {
			c.HTML(200, "dict/add.html", nil)
		})
		dictGroup.GET("/edit", func(c *gin.Context) {
			c.HTML(200, "dict/edit.html", nil)
		})

		// 字典接口
		dictGroup.GET("/page", api.Dict.Page)
		dictGroup.POST("/add", api.Dict.Add)
		dictGroup.POST("/update", api.Dict.Update)
		dictGroup.DELETE("/delete/:id", api.Dict.Delete)
		dictGroup.GET("/get/:id", api.Dict.Get)
		dictGroup.POST("/editStatus", api.Dict.EditStatus)
		// 字典数据接口
		dictGroup.GET("/data/list", api.DictData.List)
		dictGroup.GET("/data/page", api.DictData.Page)
	}
}
