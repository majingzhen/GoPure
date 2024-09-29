package api

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/service"
	"matuto.com/GoPure/src/common/response"
	"strconv"
)

type UserAPI struct {
	userService service.UserService
}

// Page 获取用户分页列表
func (api *UserAPI) Page(c *gin.Context) {
	// 获取参数
	var userReq view.UserReqPageVO
	if err := c.ShouldBind(&userReq); err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	// 获取用户列表
	userPage, err := api.userService.Page(userReq)
	if err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 正常返回
	response.OkWithData(userPage, c)
}

// GetUserById 根据id获取用户信息
func (api *UserAPI) GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	byId, err := api.userService.GetUserById(id)
	if err != nil {
		// 错误处理
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 正常返回
	response.OkWithData(byId, c)
}
