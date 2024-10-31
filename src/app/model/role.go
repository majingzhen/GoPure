package model

import "time"

// Role 角色模型
type Role struct {
	Id          int       `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"not null"`
	Code        string    `json:"code" gorm:"not null;unique"`
	Description string    `json:"description"`
	CreateTime  time.Time `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime  time.Time `json:"updateTime" gorm:"autoUpdateTime"`
}

// TableName 指定表名
func (Role) TableName() string {
	return "p_role"
}

// 角色编码常量
const (
	RoleAdmin = "admin" // 管理员角色编码
)
