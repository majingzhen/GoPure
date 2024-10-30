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
		// 字典数据接口
		dictGroup.GET("/data/list", api.DictData.List)
		dictGroup.GET("/data/page", api.DictData.Page)
	}
}
