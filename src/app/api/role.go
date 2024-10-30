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

// JumpRoleView 跳转角色管理页面
func (api *RoleAPI) JumpRoleView(c *gin.Context) {
	response.JumpView(c, "role/index.html")
}

// JumpRoleFormView 跳转角色表单页面
func (api *RoleAPI) JumpRoleFormView(c *gin.Context) {
	response.JumpView(c, "role/form.html")
}

// JumpRoleAuthView 跳转角色权限分配页面
func (api *RoleAPI) JumpRoleAuthView(c *gin.Context) {
	response.JumpView(c, "role/auth.html")
}

// Page 获取角色列表
func (api *RoleAPI) Page(c *gin.Context) {
	var req view.RoleReqPageVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	page, err := service.Role.Page(req.PageNum, req.PageSize, map[string]interface{}{
		"name": req.Name,
		"code": req.Code,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(page, c)
}

// Get 获取角色详情
func (api *RoleAPI) Get(c *gin.Context) {
	id := utils.GetIntParam(c, "id")
	if id == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	role, err := service.Role.GetRoleById(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取角色菜单
	menuIds, err := service.Menu.GetMenuIdsByRoleId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 构建返回数据
	roleVO := &view.RoleVO{
		Role:    role,
		MenuIds: menuIds,
	}
	response.OkWithData(roleVO, c)
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
	err := service.Role.Add(role, req.MenuIds)
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
	role, err := service.Role.GetRoleById(req.Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 更新基本信息
	role.Name = req.Name
	role.Code = req.Code
	role.Description = req.Description
	// 更新角色
	err = service.Role.Update(role, req.MenuIds)
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

// GetMenuTree 获取角色菜单树
func (api *RoleAPI) GetMenuTree(c *gin.Context) {
	menuType := utils.GetIntParam(c, "menuType")
	menus, err := service.Menu.GetMenuTree(menuType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(menus, c)
}

// SaveMenus 保存角色菜单权限
func (api *RoleAPI) SaveMenus(c *gin.Context) {
	var req struct {
		Id      int      `json:"id" binding:"required"`
		MenuIds []string `json:"menuIds"`
	}
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := service.Role.SaveMenus(req.Id, req.MenuIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// List 获取所有角色列表
func (api *RoleAPI) List(c *gin.Context) {
	roles, err := service.Role.List()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(roles, c)
}
