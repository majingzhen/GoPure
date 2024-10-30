package model

import (
	"time"
)

// DictData 字典数据模型
type DictData struct {
	Id              int        `json:"id" gorm:"primary_key"`
	DictLabel       string     `json:"dictLabel" gorm:"column:dict_label;size:512;not null;comment:展示值"`
	DictValue       string     `json:"dictValue" gorm:"column:dict_value;size:512;not null;comment:字典值"`
	DictExtendValue string     `json:"dictExtendValue" gorm:"column:dict_extend_value;size:512;comment:扩展值"`
	Status          string     `json:"status" gorm:"column:status;type:char(1);default:'0';comment:状态0:启用,1禁用"`
	Seq             int        `json:"seq" gorm:"column:seq;not null;default:0;comment:排序"`
	DictType        string     `json:"dictType" gorm:"column:dict_type;size:255;not null;comment:字典类型"`
	ParentDictType  string     `json:"parentDictType" gorm:"column:parent_dict_type;size:255;not null;comment:父级字典类型"`
	CreateTime      time.Time  `json:"createTime" gorm:"column:create_time"`
	UpdateTime      *time.Time `json:"updateTime" gorm:"column:update_time"`
	CreateUserID    *int       `json:"createUserId,omitempty" gorm:"column:create_user_id"`
	UpdateUserID    *int       `json:"updateUserId,omitempty" gorm:"column:update_user_id"`
}

// TableName 指定表名
func (DictData) TableName() string {
	return "p_dict_data"
}

// Status 字典数据状态常量
const (
	DictDataStatusEnabled  = "0" // 启用
	DictDataStatusDisabled = "1" // 禁用
)
