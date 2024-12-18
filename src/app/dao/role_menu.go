package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/global"
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
func (dao *RoleMenuDAO) GetMenuIdsByRoleId(roleId int) ([]string, error) {
	var menuIds []string
	err := global.GormDao.Model(&model.RoleMenu{}).
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

// DeleteByMenuId 根据菜单ID删除角色菜单关联
func (dao *RoleMenuDAO) DeleteByMenuId(tx *gorm.DB, id string) error {
	return tx.Where("menu_id = ?", id).Delete(&model.RoleMenu{}).Error
}

func (dao *RoleMenuDAO) DeleteByRoleIds(tx *gorm.DB, ids []int) error {
	return tx.Where("role_id in ?", ids).Delete(&model.RoleMenu{}).Error
}

// AuthRole 角色授权
func (dao *RoleMenuDAO) AuthRole(roleId int, menuIds []string) error {
	return global.GormDao.Transaction(func(tx *gorm.DB) error {
		// 删除原有权限
		if err := tx.Model(model.RoleMenu{}).Where("role_id = ?", roleId).Delete(nil).Error; err != nil {
			return err
		}

		// 添加新权限
		if len(menuIds) > 0 {
			var roleMenus []map[string]interface{}
			for _, menuId := range menuIds {
				roleMenus = append(roleMenus, map[string]interface{}{
					"role_id": roleId,
					"menu_id": menuId,
				})
			}
			return tx.Model(model.RoleMenu{}).Create(roleMenus).Error
		}
		return nil
	})
}
