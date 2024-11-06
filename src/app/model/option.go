package model

import (
	"time"
)

// Option 选项模型
type Option struct {
	Id             int        `json:"id" gorm:"primary_key"`
	Key            string     `json:"key" gorm:"column:key;size:256;not null;uniqueIndex;comment:key"`
	Value          string     `json:"value" gorm:"column:value;type:text;comment:value"`
	Title          string     `json:"title" gorm:"column:title;size:512;comment:标题"`
	Identification string     `json:"identification" gorm:"column:identification;size:256;comment:标识"`
	CreateTime     time.Time  `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime     *time.Time `json:"updateTime" gorm:"autoCreateTime"`
	CreateUserID   *int       `json:"createUserId,omitempty" gorm:"column:create_user_id"`
	UpdateUserID   *int       `json:"updateUserId,omitempty" gorm:"column:update_user_id"`
}

// TableName 指定表名
func (Option) TableName() string {
	return "p_option"
}
