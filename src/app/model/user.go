package model

import "time"

type User struct {
	ID           int        `json:"id" gorm:"primary_key"`
	Account      string     `json:"account" gorm:"unique"`
	UserName     string     `json:"userName" gorm:"column:user_name"`
	Password     string     `json:"password"`
	Salt         string     `json:"salt"`
	Status       int        `json:"status"`
	Avatar       string     `json:"avatar,omitempty"`
	Email        string     `json:"email,omitempty"`
	Website      string     `json:"website,omitempty"`
	CreateTime   time.Time  `json:"createTime"`
	UpdateTime   *time.Time `json:"updateTime,omitempty"`
	CreateUserID *int       `json:"createUserId,omitempty"`
	UpdateUserID *int       `json:"updateUserId,omitempty"`
	Remark       string     `json:"remark,omitempty"`
	Mobile       string     `json:"mobile"`
	Sex          *int       `json:"sex,omitempty"`
	LoginIP      string     `json:"loginIp"`
	LoginDate    *time.Time `json:"loginDate,omitempty"`
}

func (User) TableName() string {
	return "p_user"
}
