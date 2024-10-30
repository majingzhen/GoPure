package service

import (
	"errors"
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/dao"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

var Role = new(RoleService)

type RoleService struct{}

// GetRoleById 根据id获取角色
func (service *RoleService) GetRoleById(id int) (*model.Role, error) {
	return dao.Role.GetRoleById(global.GormDao, id)
}

// GetRoleByCode 根据角色编码获取角色
func (service *RoleService) GetRoleByCode(code string) (*model.Role, error) {
	return dao.Role.GetRoleByCode(global.GormDao, code)
}

// Page 获取角色分页列表
func (service *RoleService) Page(pageNum, pageSize int, query map[string]interface{}) (*common.PageInfo, error) {
	return dao.Role.Page(pageNum, pageSize, query)
}

// GetByUserId 根据用户ID获取角色列表
func (service *RoleService) GetByUserId(id int) ([]*model.Role, error) {
	return dao.UserRole.GetRolesByUserId(global.GormDao, id)
}

// Add 添加角色
func (service *RoleService) Add(role *model.Role, menuIds []string) error {
	return global.GormDao.Transaction(func(tx *gorm.DB) error {
		// 1. 检查角色编码是否已存在
		exists, err := dao.Role.CheckCodeExists(tx, role.Code)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("角色编码已存在")
		}

		// 2. 保存角色信息
		if err := dao.Role.Add(tx, role); err != nil {
			return err
		}

		// 3. 保存角色菜单关联
		if len(menuIds) > 0 {
			roleMenus := make([]*model.RoleMenu, len(menuIds))
			for i, menuId := range menuIds {
				roleMenus[i] = &model.RoleMenu{
					RoleId: role.Id,
					MenuId: menuId,
				}
			}
			if err := tx.Create(&roleMenus).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// Update 更新角色
func (service *RoleService) Update(role *model.Role, menuIds []string) error {
	return global.GormDao.Transaction(func(tx *gorm.DB) error {
		// 1. 检查角色编码是否已存在（排除自身）
		exists, err := dao.Role.CheckCodeExists(tx, role.Code, role.Id)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("角色编码已存在")
		}

		// 2. 更新角色信息
		if err := dao.Role.Update(tx, role); err != nil {
			return err
		}

		// 3. 更新角色菜单关联
		// 先删除原有关联
		if err := tx.Where("role_id = ?", role.Id).Delete(&model.RoleMenu{}).Error; err != nil {
			return err
		}
		// 添加新关联
		if len(menuIds) > 0 {
			roleMenus := make([]*model.RoleMenu, len(menuIds))
			for i, menuId := range menuIds {
				roleMenus[i] = &model.RoleMenu{
					RoleId: role.Id,
					MenuId: menuId,
				}
			}
			if err := tx.Create(&roleMenus).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// Delete 删除角色
func (service *RoleService) Delete(ids []int) error {
	return global.GormDao.Transaction(func(tx *gorm.DB) error {
		// 1. 删除角色菜单关联
		if err := tx.Where("role_id IN ?", ids).Delete(&model.RoleMenu{}).Error; err != nil {
			return err
		}

		// 2. 删除用户角色关联
		if err := tx.Where("role_id IN ?", ids).Delete(&model.UserRole{}).Error; err != nil {
			return err
		}

		// 3. 删除角色
		if err := dao.Role.Delete(tx, ids); err != nil {
			return err
		}

		return nil
	})
}

// List 获取所有角色列表
func (service *RoleService) List() ([]*model.Role, error) {
	return dao.Role.List(global.GormDao)
}

// SaveMenus 保存角色菜单权限
func (service *RoleService) SaveMenus(roleId int, menuIds []string) error {
	return global.GormDao.Transaction(func(tx *gorm.DB) error {
		// 先删除原有关联
		if err := tx.Where("role_id = ?", roleId).Delete(&model.RoleMenu{}).Error; err != nil {
			return err
		}

		// 批量插入新关联
		if len(menuIds) > 0 {
			roleMenus := make([]*model.RoleMenu, len(menuIds))
			for i, menuId := range menuIds {
				roleMenus[i] = &model.RoleMenu{
					RoleId: roleId,
					MenuId: menuId,
				}
			}
			if err := tx.Create(&roleMenus).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
