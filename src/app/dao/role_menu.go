package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/model"
)

var RoleMenu = new(RoleMenuDAO)

type RoleMenuDAO struct{}

// GetMenusByRoleId 根据角色ID获取菜单列表
func (dao *RoleMenuDAO) GetMenusByRoleId(tx *gorm.DB, roleId int) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := tx.Table("p_menu").
		Joins("JOIN p_role_menu ON p_menu.id = p_role_menu.menu_id").
		Where("p_role_menu.role_id = ?", roleId).
		Order("p_menu.seq").
		Find(&menus).Error
	return menus, err
}

// GetMenuIdsByRoleId 根据角色ID获取菜单ID列表
func (dao *RoleMenuDAO) GetMenuIdsByRoleId(tx *gorm.DB, roleId int) ([]string, error) {
	var menuIds []string
	err := tx.Model(&model.RoleMenu{}).
		Where("role_id = ?", roleId).
		Pluck("menu_id", &menuIds).Error
	return menuIds, err
}

// GetMenusByRoleIdAndType 根据角色ID和菜单位置获取菜单列表
func (dao *RoleMenuDAO) GetMenusByRoleIdAndType(tx *gorm.DB, roleId int, menuPosition int) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := tx.Table("p_menu").
		Joins("JOIN p_role_menu ON p_menu.id = p_role_menu.menu_id").
		Where("p_role_menu.role_id = ? AND p_menu.menu_position = ? AND p_menu.status = ?",
			roleId, menuPosition, model.StatusEnabled).
		Order("p_menu.seq").
		Find(&menus).Error
	return menus, err
}

// BatchSave 批量保存角色菜单关联
func (dao *RoleMenuDAO) BatchSave(tx *gorm.DB, roleId int, menuIds []string) error {
	// 先删除原有关联
	if err := tx.Where("role_id = ?", roleId).Delete(&model.RoleMenu{}).Error; err != nil {
		return err
	}

	// 批量插入新关联
	roleMenus := make([]*model.RoleMenu, len(menuIds))
	for i, menuId := range menuIds {
		roleMenus[i] = &model.RoleMenu{
			RoleId: roleId,
			MenuId: menuId,
		}
	}
	return tx.Create(&roleMenus).Error
}
