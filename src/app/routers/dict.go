package routers

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api"
	"matuto.com/GoPure/src/framework/middleware"
)

var Dict = new(DictRouter)

type DictRouter struct{}

func (r *DictRouter) InitDictRouter(e *gin.Engine) {
	dictGroup := e.Group("dict").Use(middleware.AuthMiddleware()).Use(middleware.PermissionMiddleware())
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
		dictGroup.POST("/data/add", api.DictData.Add)
		dictGroup.POST("/data/edit", api.DictData.Edit)
		dictGroup.DELETE("/data/delete/:id", api.DictData.Delete)
		dictGroup.GET("/data/get/:id", api.DictData.Get)
		dictGroup.POST("/data/editStatus", api.DictData.EditStatus)
		// 字典数据页面
		dictGroup.GET("/data/add", func(c *gin.Context) {
			c.HTML(200, "dict/data/add.html", nil)
		})
		dictGroup.GET("/data/edit", func(c *gin.Context) {
			c.HTML(200, "dict/data/edit.html", nil)
		})
		dictGroup.GET("/data/index", func(c *gin.Context) {
			c.HTML(200, "dict/data/index.html", nil)
		})
	}
}
