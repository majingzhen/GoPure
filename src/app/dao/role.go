package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

var Role = new(RoleDAO)

type RoleDAO struct{}

// Page 分页查询角色
func (d *RoleDAO) Page(req view.RoleReqPageVO) (*common.PageInfo, error) {
	var roles []model.Role
	var total int64

	// 构建查询条件
	db := global.GormDao.Model(&model.Role{})
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Code != "" {
		db = db.Where("code LIKE ?", "%"+req.Code+"%")
	}

	// 查询总数
	err := db.Count(&total).Error
	if err != nil {
		return nil, err
	}

	// 分页查询数据
	offset := (req.PageNum - 1) * req.PageSize
	err = db.Offset(offset).Limit(req.PageSize).Order("id desc").Find(&roles).Error
	if err != nil {
		return nil, err
	}

	// 构建分页结果
	return &common.PageInfo{
		Total: total,
		Rows:  roles,
	}, nil
}

// List 获取角色列表
func (d *RoleDAO) List() ([]model.Role, error) {
	var roles []model.Role
	err := global.GormDao.Order("id desc").Find(&roles).Error
	return roles, err
}

// GetById 根据ID获取角色
func (d *RoleDAO) GetById(id int) (*model.Role, error) {
	var role model.Role
	err := global.GormDao.First(&role, id).Error
	return &role, err
}

// Add 添加角色
func (d *RoleDAO) Add(role *model.Role) error {
	return global.GormDao.Create(role).Error
}

// Update 更新角色
func (d *RoleDAO) Update(role *model.Role) error {
	return global.GormDao.Save(role).Error
}

// Delete 删除角色
func (d *RoleDAO) Delete(tx *gorm.DB, ids []int) error {
	return tx.Delete(&model.Role{}, ids).Error
}

// UpdateStatus 更新角色状态
func (d *RoleDAO) UpdateStatus(id int, status string) error {
	return global.GormDao.Model(&model.Role{}).Where("id = ?", id).Update("status", status).Error
}

// GetByCode 根据角色编码获取角色
func (d *RoleDAO) GetByCode(code string) (*model.Role, error) {
	var role model.Role
	err := global.GormDao.Where("code = ?", code).First(&role).Error
	return &role, err
}

// CheckNameExist 检查角色名称是否存在
func (d *RoleDAO) CheckNameExist(name string, excludeId ...int) (bool, error) {
	var count int64
	db := global.GormDao.Model(&model.Role{}).Where("name = ?", name)
	if len(excludeId) > 0 {
		db = db.Where("id != ?", excludeId[0])
	}
	err := db.Count(&count).Error
	return count > 0, err
}

// CheckCodeExist 检查角色编码是否存在
func (d *RoleDAO) CheckCodeExist(code string, excludeId ...int) (bool, error) {
	var count int64
	db := global.GormDao.Model(&model.Role{}).Where("code = ?", code)
	if len(excludeId) > 0 {
		db = db.Where("id != ?", excludeId[0])
	}
	err := db.Count(&count).Error
	return count > 0, err
}

func (d *RoleDAO) GetByUserId(id int) ([]*model.Role, error) {
	var roles []*model.Role
	err := global.GormDao.Model(&model.Role{}).Where("id in (select role_id from p_user_role where user_id = ?)", id).Find(&roles).Error
	return roles, err
}

// GetRoleMenus 获取角色菜单ID列表
func (d *RoleDAO) GetRoleMenus(roleId int) ([]int, error) {
	var menuIds []int
	err := global.GormDao.Table("role_menu").
		Select("menu_id").
		Where("role_id = ?", roleId).
		Pluck("menu_id", &menuIds).Error
	return menuIds, err
}

// AuthRole 角色授权
func (d *RoleDAO) AuthRole(roleId int, menuIds []int) error {
	return global.GormDao.Transaction(func(tx *gorm.DB) error {
		// 删除原有权限
		if err := tx.Table("role_menu").Where("role_id = ?", roleId).Delete(nil).Error; err != nil {
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
			return tx.Table("role_menu").Create(roleMenus).Error
		}
		return nil
	})
}
