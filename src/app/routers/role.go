package routers

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api"
	"matuto.com/GoPure/src/framework/middleware"
)

var Role = new(RoleRouter)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(e *gin.Engine) {
	roleGroup := e.Group("role").Use(middleware.AuthMiddleware())
	{
		roleGroup.GET("/", api.Role.JumpRoleView)
		roleGroup.GET("/form", api.Role.JumpRoleFormView)
		roleGroup.GET("/auth", api.Role.JumpRoleAuthView)
		roleGroup.GET("/list", api.Role.List)
		roleGroup.GET("/page", api.Role.Page)
		roleGroup.GET("/get", api.Role.Get)
		roleGroup.POST("/add", api.Role.Add)
		roleGroup.POST("/update", api.Role.Update)
		roleGroup.POST("/delete", api.Role.Delete)
		roleGroup.GET("/getMenuTree", api.Role.GetMenuTree)
		roleGroup.POST("/saveMenus", api.Role.SaveMenus)
	}
}
