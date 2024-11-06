package routers

import (
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter(r *gin.Engine) {
	// 初始化系统路由
	InitSystemRouter(r)
	// 初始化用户路由
	User.InitUserRouter(r)
	// 初始化角色路由
	Role.InitRoleRouter(r)
	// 初始化菜单路由
	Menu.InitMenuRouter(r)
	// 初始化字典路由
	Dict.InitDictRouter(r)
	// 初始化选项路由
	Option.InitOptionRoutes(r)
}
