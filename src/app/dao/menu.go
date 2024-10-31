package dao

import (
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/global"
)

var Menu = new(MenuDAO)

type MenuDAO struct{}

// List 获取菜单列表
func (d *MenuDAO) List(req view.MenuListReqVO) ([]model.Menu, error) {
	var menus []model.Menu
	db := global.GormDao.Model(&model.Menu{})

	// 根据条件查询
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	if req.MenuType != "" {
		db = db.Where("menu_type = ?", req.MenuType)
	}
	if req.MenuPosition != "" {
		db = db.Where("menu_position = ?", req.MenuPosition)
	}

	err := db.Order("seq").Find(&menus).Error
	return menus, err
}

// GetById 根据ID获取菜单
func (d *MenuDAO) GetById(id string) (*model.Menu, error) {
	var menu model.Menu
	err := global.GormDao.First(&menu, "id = ?", id).Error
	return &menu, err
}

// Add 添加菜单
func (d *MenuDAO) Add(menu *model.Menu) error {
	return global.GormDao.Create(menu).Error
}

// Update 更新菜单
func (d *MenuDAO) Update(menu *model.Menu) error {
	return global.GormDao.Save(menu).Error
}

// Delete 删除菜单
func (d *MenuDAO) Delete(id string) error {
	return global.GormDao.Delete(&model.Menu{}, "id = ?", id).Error
}

// HasChildren 检查是否有子菜单
func (d *MenuDAO) HasChildren(id string) (bool, error) {
	var count int64
	err := global.GormDao.Model(&model.Menu{}).Where("pid = ?", id).Count(&count).Error
	return count > 0, err
}

// GetByRoleId 获取角色菜单
func (d *MenuDAO) GetByRoleId(roleId int) ([]model.Menu, error) {
	var menus []model.Menu
	err := global.GormDao.Model(&model.Menu{}).
		Where("id in (select menu_id from p_role_menu where role_id = ?)", roleId).
		Order("seq").
		Find(&menus).Error
	return menus, err
}

// GetByRoleIds 获取多个角色的菜单
func (d *MenuDAO) GetByRoleIds(roleIds []int) ([]model.Menu, error) {
	var menus []model.Menu
	err := global.GormDao.Model(&model.Menu{}).
		Where("id in (select menu_id from p_role_menu where role_id in ?)", roleIds).
		Order("seq").
		Find(&menus).Error
	return menus, err
}
