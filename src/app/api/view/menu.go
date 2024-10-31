package view

import (
	"matuto.com/GoPure/src/app/model"
)

// MenuListReqVO 菜单列表请求
type MenuListReqVO struct {
	Name         string `form:"name"`         // 菜单名称
	Status       string `form:"status"`       // 状态
	MenuType     string `form:"menuType"`     // 菜单类型
	MenuPosition string `form:"menuPosition"` // 菜单位置
}

// MenuVO 菜单树节点
type MenuVO struct {
	*model.Menu
	Children []*MenuVO `json:"children"` // 子菜单
}

// MenuAddReqVO 添加菜单请求
type MenuAddReqVO struct {
	Pid          string `json:"pid"`                     // 父级id
	Name         string `json:"name" binding:"required"` // 菜单名称
	Url          string `json:"url"`                     // 菜单链接
	Icon         string `json:"icon"`                    // 菜单图标
	Seq          int    `json:"seq,string"`              // 排序序号
	Target       string `json:"target"`                  // 菜单打开方式
	Status       string `json:"status"`                  // 菜单状态
	MenuType     string `json:"menuType"`                // 菜单类型
	MenuPosition string `json:"menuPosition"`            // 菜单位置
	IsFrame      string `json:"isFrame"`                 // 是否为外链
}

// MenuUpdateReqVO 更新菜单请求
type MenuUpdateReqVO struct {
	Id           string `json:"id" binding:"required"`   // 主键
	Pid          string `json:"pid"`                     // 父级id
	Name         string `json:"name" binding:"required"` // 菜单名称
	Url          string `json:"url"`                     // 菜单链接
	Icon         string `json:"icon"`                    // 菜单图标
	Seq          int    `json:"seq,string"`              // 排序序号
	Target       string `json:"target"`                  // 菜单打开方式
	Status       string `json:"status"`                  // 菜单状态
	MenuType     string `json:"menuType"`                // 菜单类型
	MenuPosition string `json:"menuPosition"`            // 菜单位置
	IsFrame      string `json:"isFrame"`                 // 是否为外链
}

// MenuDeleteReqVO 删除菜单请求
type MenuDeleteReqVO struct {
	Id string `json:"id" binding:"required"`
}

// MenuStatusReqVO 更新菜单状态请求
type MenuStatusReqVO struct {
	Id     string `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}
