package model

import (
	"time"
)

// Role 角色模型
type Role struct {
	Id           int        `json:"id" gorm:"primary_key"`
	Name         string     `json:"name" gorm:"column:name;size:32;not null;comment:角色名"`
	Description  string     `json:"description" gorm:"column:description;size:256;comment:角色描述"`
	Code         string     `json:"code" gorm:"column:code;size:32;not null;comment:角色码"`
	CreateTime   time.Time  `json:"createTime" gorm:"column:create_time"`
	UpdateTime   *time.Time `json:"updateTime" gorm:"column:update_time"`
	CreateUserID *int       `json:"createUserId,omitempty" gorm:"column:create_user_id"`
	UpdateUserID *int       `json:"updateUserId,omitempty" gorm:"column:update_user_id"`
}

// TableName 指定表名
func (Role) TableName() string {
	return "p_role"
}

// 角色编码常量
const (
	RoleAdmin = "admin" // 管理员角色编码
)
