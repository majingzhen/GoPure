package view

import (
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
)

type LoginUserVO struct {
	Account  string        `json:"account"`
	UserName string        `json:"userName"`
	Avatar   string        `json:"avatar"`
	Email    string        `json:"email"`
	Mobile   string        `json:"mobile"`
	Roles    []*model.Role `json:"roles"`
	Menus    []*MenuVO     `json:"menus"`
}

// UserReqPageVO 用户分页请求
type UserReqPageVO struct {
	*common.PageView
	Account  string `form:"account"`
	UserName string `form:"userName"`
	Status   string `form:"status"`
}

// UserVO 用户详情响应
type UserVO struct {
	*model.User
	RoleIds []int `json:"roleIds"` // 角色ID列表
}

// UserAddReqVO 添加用户请求
type UserAddReqVO struct {
	Account  string `json:"account" binding:"required"`
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password" binding:"required"`
	Status   string `json:"status"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Sex      string `json:"sex"`
	Remark   string `json:"remark"`
	RoleIds  []int  `json:"roleIds"` // 角色ID列表
}

// UserUpdateReqVO 更新用户请求
type UserUpdateReqVO struct {
	Id       int    `json:"id,string" binding:"required"`
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password"` // 可选，不填则不修改密码
	Status   string `json:"status"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Sex      string `json:"sex"`
	Remark   string `json:"remark"`
	RoleIds  []int  `json:"roleIds"` // 角色ID列表
	Avatar   string `json:"avatar"`
}

// UserDeleteReqVO 删除用户请求
type UserDeleteReqVO struct {
	Ids []int `json:"ids,string" binding:"required"`
}

// UserStatusReqVO 更新用户状态请求
type UserStatusReqVO struct {
	Id     int    `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

// UserResetPwdReqVO 重置密码请求
type UserResetPwdReqVO struct {
	Id       int    `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
