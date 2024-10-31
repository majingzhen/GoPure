package api

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/app/service"
	"matuto.com/GoPure/src/common/response"
)

var Menu = new(MenuAPI)

type MenuAPI struct{}

// List 获取菜单列表
func (api *MenuAPI) List(c *gin.Context) {
	var req view.MenuListReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	menus := service.Menu.List(req)
	response.OkWithData(menus, c)
}

// Get 获取菜单详情
func (api *MenuAPI) Get(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	menu, err := service.Menu.GetById(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(menu, c)
}

// Add 添加菜单
func (api *MenuAPI) Add(c *gin.Context) {
	var req view.MenuAddReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 构建菜单对象
	menu := &model.Menu{
		Pid:          req.Pid,
		Name:         req.Name,
		Url:          req.Url,
		Icon:         req.Icon,
		Seq:          req.Seq,
		Target:       req.Target,
		Status:       req.Status,
		MenuType:     req.MenuType,
		MenuPosition: req.MenuPosition,
		IsFrame:      req.IsFrame,
	}
	// 保存菜单
	err := service.Menu.Add(menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Update 更新菜单
func (api *MenuAPI) Update(c *gin.Context) {
	var req view.MenuUpdateReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取原菜单信息
	menu, err := service.Menu.GetById(req.Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 更新基本信息
	menu.Pid = req.Pid
	menu.Name = req.Name
	menu.Url = req.Url
	menu.Icon = req.Icon
	menu.Seq = req.Seq
	menu.Target = req.Target
	menu.Status = req.Status
	menu.MenuType = req.MenuType
	menu.MenuPosition = req.MenuPosition
	menu.IsFrame = req.IsFrame

	// 更新菜单
	err = service.Menu.Update(menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Delete 删除菜单
func (api *MenuAPI) Delete(c *gin.Context) {
	var req view.MenuDeleteReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.Menu.Delete(req.Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// UpdateStatus 更新菜单状态
func (api *MenuAPI) UpdateStatus(c *gin.Context) {
	var req view.MenuStatusReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.Menu.UpdateStatus(req.Id, req.Status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// JumpMenuView 跳转菜单管理页面
func (api *MenuAPI) JumpMenuView(c *gin.Context) {
	response.JumpView(c, "menu/index.html")
}

// JumpMenuAddView 跳转菜单添加页面
func (api *MenuAPI) JumpMenuAddView(c *gin.Context) {
	response.JumpView(c, "menu/add.html")
}

// JumpMenuEditView 跳转菜单更新页面
func (api *MenuAPI) JumpMenuEditView(c *gin.Context) {
	response.JumpView(c, "menu/edit.html")
}
