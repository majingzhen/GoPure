package model

import (
	"time"
)

// Status 用户状态常量
const (
	StatusEnabled  = "0" // 启用
	StatusDisabled = "1" // 禁用
)

// Sex 性别常量
const (
	SexMale   = "1" // 男
	SexFemale = "2" // 女
)

// User 用户模型
type User struct {
	Id           int        `json:"id" gorm:"primary_key"`
	Account      string     `json:"account" gorm:"unique"`
	UserName     string     `json:"userName" gorm:"column:user_name"`
	Password     string     `json:"password"`
	Salt         string     `json:"salt"`
	Status       string     `json:"status" gorm:"type:char(1);default:'0'"`
	Avatar       string     `json:"avatar,omitempty"`
	Email        string     `json:"email,omitempty"`
	Website      string     `json:"website,omitempty"`
	Remark       string     `json:"remark,omitempty"`
	Mobile       string     `json:"mobile"`
	Sex          string     `json:"sex" gorm:"type:char(1)"`
	LoginIP      string     `json:"loginIp" gorm:"column:login_ip"`
	LoginDate    *time.Time `json:"loginDate,omitempty" gorm:"column:login_date"`
	CreateTime   time.Time  `json:"createTime" gorm:"column:create_time"`
	UpdateTime   *time.Time `json:"updateTime" gorm:"column:update_time"`
	CreateUserID *int       `json:"createUserId,omitempty" gorm:"column:create_user_id"`
	UpdateUserID *int       `json:"updateUserId,omitempty" gorm:"column:update_user_id"`
}

func (User) TableName() string {
	return "p_user"
}
