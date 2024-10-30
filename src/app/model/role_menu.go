package model

// RoleMenu 角色菜单关联模型
type RoleMenu struct {
	RoleId int    `json:"roleId" gorm:"column:role_id;not null;comment:角色id"`
	MenuId string `json:"menuId" gorm:"column:menu_id;size:64;not null;comment:菜单id"`
}

// TableName 指定表名
func (RoleMenu) TableName() string {
	return "p_role_menu"
}
