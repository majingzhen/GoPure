package api

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/app/service"
	"matuto.com/GoPure/src/common/response"
	"matuto.com/GoPure/src/utils"
)

var Role = new(RoleAPI)

type RoleAPI struct{}

// Page 获取角色分页
func (api *RoleAPI) Page(c *gin.Context) {
	var req view.RoleReqPageVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	page, err := service.Role.Page(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(page, c)
}

// List 获取角色列表
func (api *RoleAPI) List(c *gin.Context) {
	roles, err := service.Role.List()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(roles, c)
}

// Get 获取角色详情
func (api *RoleAPI) Get(c *gin.Context) {
	id := utils.GetIntParam(c, "id")
	if id == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	role, err := service.Role.GetById(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(role, c)
}

// Add 添加角色
func (api *RoleAPI) Add(c *gin.Context) {
	var req view.RoleAddReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 构建角色对象
	role := &model.Role{
		Name:        req.Name,
		Code:        req.Code,
		Description: req.Description,
	}
	// 保存角色
	err := service.Role.Add(role)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Update 更新角色
func (api *RoleAPI) Update(c *gin.Context) {
	var req view.RoleUpdateReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取原角色信息
	role, err := service.Role.GetById(req.Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 更新基本信息
	role.Name = req.Name
	role.Code = req.Code
	role.Description = req.Description

	// 更新角色
	err = service.Role.Update(role)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Delete 删除角色
func (api *RoleAPI) Delete(c *gin.Context) {
	var req view.RoleDeleteReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.Role.Delete(req.Ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// JumpRoleView 跳转角色管理页面
func (api *RoleAPI) JumpRoleView(c *gin.Context) {
	response.JumpView(c, "role/index.html")
}

// JumpRoleAddView 跳转角色添加页面
func (api *RoleAPI) JumpRoleAddView(c *gin.Context) {
	response.JumpView(c, "role/add.html")
}

// JumpRoleEditView 跳转角色更新页面
func (api *RoleAPI) JumpRoleEditView(c *gin.Context) {
	response.JumpView(c, "role/edit.html")
}

// JumpRoleAuthView 跳转角色授权页面
func (api *RoleAPI) JumpRoleAuthView(c *gin.Context) {
	response.JumpView(c, "role/auth.html")
}

// GetRoleMenus 获取角色菜单
func (api *RoleAPI) GetRoleMenus(c *gin.Context) {
	id := utils.GetIntParam(c, "id")
	if id == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	menuIds, err := service.Role.GetRoleMenus(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(menuIds, c)
}

// AuthRole 角色授权
func (api *RoleAPI) AuthRole(c *gin.Context) {
	var req view.RoleAuthReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.Role.AuthRole(req.Id, req.MenuIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
