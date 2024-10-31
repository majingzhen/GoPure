package view

import (
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
)

// RoleReqPageVO 角色分页请求
type RoleReqPageVO struct {
	common.PageView
	Name string `form:"name"`
	Code string `form:"code"`
}

// RoleVO 角色详情响应
type RoleVO struct {
	*model.Role
	MenuIds []int `json:"menuIds"` // 菜单ID列表
}

// RoleAddReqVO 添加角色请求
type RoleAddReqVO struct {
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Description string `json:"description"`
}

// RoleUpdateReqVO 更新角色请求
type RoleUpdateReqVO struct {
	Id          int    `json:"id,string" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Description string `json:"description"`
}

// RoleDeleteReqVO 删除角色请求
type RoleDeleteReqVO struct {
	Ids []int `json:"ids" binding:"required"`
}

// RoleAuthReqVO 角色授权请求
type RoleAuthReqVO struct {
	Id      int   `json:"id" binding:"required"`
	MenuIds []int `json:"menuIds"`
}
