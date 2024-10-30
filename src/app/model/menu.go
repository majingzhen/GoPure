package model

import (
	"time"
)

// Menu 菜单模型
type Menu struct {
	Id           string     `json:"id" gorm:"primary_key;comment:主键"`
	Pid          string     `json:"pid" gorm:"column:pid;size:64;default:-1;comment:父级id"`
	Name         string     `json:"name" gorm:"column:name;size:128;not null;comment:菜单名"`
	Url          string     `json:"url" gorm:"column:url;size:128;comment:菜单链接"`
	Icon         string     `json:"icon" gorm:"column:icon;size:64;comment:菜单图标"`
	Seq          int        `json:"seq" gorm:"column:seq;default:0;comment:排序序号"`
	Target       string     `json:"target" gorm:"column:target;type:char(1);default:'0';comment:菜单打开方式:0本页,1:新窗口"`
	Status       string     `json:"status" gorm:"column:status;type:char(1);default:'0';comment:菜单状态0:启用,1禁用;index"`
	MenuType     string     `json:"menuType" gorm:"column:menu_type;type:char(1);comment:菜单类型（0目录1菜单2按钮）"`
	MenuPosition string     `json:"menuPosition" gorm:"column:menu_position;type:char(1);default:'0';comment:菜单位置0:前台,1:后台"`
	IsFrame      string     `json:"isFrame" gorm:"column:is_frame;type:char(1);default:'1';comment:是否为外链（0是 1否）"`
	CreateTime   time.Time  `json:"createTime" gorm:"column:create_time;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime   *time.Time `json:"updateTime" gorm:"column:update_time;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;comment:更新时间"`
	CreateUserID *int       `json:"createUserId,omitempty" gorm:"column:create_user_id;comment:添加人"`
	UpdateUserID *int       `json:"updateUserId,omitempty" gorm:"column:update_user_id;comment:更新人"`
	Children     []*Menu    `json:"children" gorm:"-"` // 子菜单列表，不映射到数据库
}

// TableName 指定表名
func (Menu) TableName() string {
	return "p_menu"
}

// MenuType 菜单类型常量
const (
	MenuTypeDirectory = "0" // 目录
	MenuTypeMenu      = "1" // 菜单
	MenuTypeButton    = "2" // 按钮
)

// Target 菜单打开方式常量
const (
	TargetSelf  = "0" // 本页打开
	TargetBlank = "1" // 新窗口打开
)

// MenuPosition 菜单位置常量
const (
	MenuPositionFrontend = "0" // 前台菜单
	MenuPositionBackend  = "1" // 后台菜单
)

// IsFrame 是否外链常量
const (
	IsFrameYes = "0" // 是外链
	IsFrameNo  = "1" // 不是外链
)
