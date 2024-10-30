package view

import "matuto.com/GoPure/src/app/model"

// RoleReqPageVO 角色分页请求
type RoleReqPageVO struct {
	Name     string `form:"name"`
	Code     string `form:"code"`
	PageNum  int    `form:"page" binding:"required"`
	PageSize int    `form:"limit" binding:"required"`
}

// RoleVO 角色详情响应
type RoleVO struct {
	*model.Role
	MenuIds []string `json:"menuIds"` // 菜单ID列表
}

// RoleAddReqVO 添加角色请求
type RoleAddReqVO struct {
	Name        string   `json:"name" binding:"required"`
	Code        string   `json:"code" binding:"required"`
	Description string   `json:"description"`
	MenuIds     []string `json:"menuIds"` // 菜单ID列表
}

// RoleUpdateReqVO 更新角色请求
type RoleUpdateReqVO struct {
	Id          int      `json:"id" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Code        string   `json:"code" binding:"required"`
	Description string   `json:"description"`
	MenuIds     []string `json:"menuIds"` // 菜单ID列表
}

// RoleDeleteReqVO 删除角色请求
type RoleDeleteReqVO struct {
	Ids []int `json:"ids" binding:"required"`
}
