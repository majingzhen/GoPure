package model

import (
	"time"
)

// Dict 字典模型
type Dict struct {
	Id           int        `json:"id" gorm:"primary_key"`
	DictType     string     `json:"dictType" gorm:"column:dict_type;size:255;not null;comment:字典类型"`
	Status       int8       `json:"status" gorm:"column:status;not null;default:0;comment:状态"`
	Remark       string     `json:"remark" gorm:"column:remark;size:500;comment:备注"`
	DictName     string     `json:"dictName" gorm:"column:dict_name;size:255;not null;comment:字典名"`
	Seq          int        `json:"seq" gorm:"column:seq;not null;default:0;comment:排序"`
	CreateTime   time.Time  `json:"createTime" gorm:"column:create_time"`
	UpdateTime   *time.Time `json:"updateTime" gorm:"column:update_time"`
	CreateUserID *int       `json:"createUserId,omitempty" gorm:"column:create_user_id"`
	UpdateUserID *int       `json:"updateUserId,omitempty" gorm:"column:update_user_id"`
}

// TableName 指定表名
func (Dict) TableName() string {
	return "p_dict"
}

// Status 字典状态常量
const (
	DictStatusEnabled  int8 = 0 // 启用
	DictStatusDisabled int8 = 1 // 禁用
)
