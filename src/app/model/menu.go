package model

import "time"

// Menu 菜单模型
type Menu struct {
	Id           string    `json:"id" gorm:"primary_key"`
	Pid          string    `json:"pid" gorm:"default:-1"`         // 父级id
	Name         string    `json:"name" gorm:"not null"`          // 菜单名
	Url          string    `json:"url"`                           // 菜单链接
	Icon         string    `json:"icon"`                          // 菜单图标
	Seq          int       `json:"seq" gorm:"default:0"`          // 排序序号
	Target       string    `json:"target" gorm:"default:0"`       // 菜单打开方式(0本页,1新窗口)
	Status       string    `json:"status" gorm:"default:0"`       // 菜单状态(0启用,1禁用)
	MenuType     string    `json:"menuType"`                      // 菜单类型（0目录1菜单2按钮）
	CreateUserId int       `json:"createUserId"`                  // 添加人
	UpdateUserId int       `json:"updateUserId"`                  // 更新人
	MenuPosition string    `json:"menuPosition" gorm:"default:0"` // 菜单位置(0前台,1后台)
	IsFrame      string    `json:"isFrame" gorm:"default:1"`      // 是否为外链（0是 1否）
	CreateTime   time.Time `json:"createTime" gorm:"autoCreateTime"`
	UpdateTime   time.Time `json:"updateTime" gorm:"autoUpdateTime"`
}

// TableName 指定表名
func (Menu) TableName() string {
	return "p_menu"
}

// 菜单类型常量
const (
	MENU_TYPE_DIR  = "0" // 目录
	MENU_TYPE_MENU = "1" // 菜单
	MENU_TYPE_BTN  = "2" // 按钮
)

// 菜单状态常量
const (
	MENU_STATUS_NORMAL = "0" // 正常
	MENU_STATUS_STOP   = "1" // 停用
)

// 菜单位置常量
const (
	MENU_POSITION_FRONT   = "0" // 前台
	MENU_POSITION_BACKEND = "1" // 后台
)

// 菜单打开方式常量
const (
	MENU_TARGET_PAGE  = "0" // 本页
	MENU_TARGET_BLANK = "1" // 新窗口
)

// 是否外链常量
const (
	MENU_FRAME_YES = "0" // 是
	MENU_FRAME_NO  = "1" // 否
)
