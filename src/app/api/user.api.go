package api

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/service"
	"matuto.com/GoPure/src/common/response"
	"strconv"
)

type UserAPI struct {
	userService service.UserService
}

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
