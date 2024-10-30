package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

var Menu = new(MenuDAO)

type MenuDAO struct{}

// GetMenuById 根据id获取菜单信息
func (dao *MenuDAO) GetMenuById(tx *gorm.DB, id string) (*model.Menu, error) {
	menu := &model.Menu{}
	err := tx.Where("id = ?", id).First(menu).Error
	return menu, err
}

// GetMenusByPid 根据父ID获取子菜单列表
func (dao *MenuDAO) GetMenusByPid(tx *gorm.DB, pid string) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := tx.Where("pid = ?", pid).Order("seq").Find(&menus).Error
	return menus, err
}

// Page 获取菜单分页列表
func (dao *MenuDAO) Page(pageNum, pageSize int, query map[string]interface{}) (*common.PageInfo, error) {
	db := global.GormDao.Model(&model.Menu{})

	// 添加查询条件
	if name, ok := query["name"].(string); ok && name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if status, ok := query["status"].(int); ok {
		db = db.Where("status = ?", status)
	}
	if menuType, ok := query["menuType"].(int); ok {
		db = db.Where("menu_type = ?", menuType)
	}

	page := common.CreatePageInfo(pageNum, pageSize)
	if err := db.Count(&page.Total).Error; err != nil {
		return nil, err
	}

	page.Calculate()
	var dataList []*model.Menu
	err := db.Order("seq").Offset(page.Offset).Limit(page.Limit).Find(&dataList).Error
	page.Rows = dataList
	return page, err
}

// GetMenusByPidAndType 根据父ID和菜单位置获取子菜单列表
func (dao *MenuDAO) GetMenusByPidAndType(tx *gorm.DB, pid string, menuPosition int) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := tx.Where("pid = ? AND menu_position = ? AND status = ?", pid, menuPosition, model.StatusEnabled).
		Order("seq").
		Find(&menus).Error
	return menus, err
}

// GetAllEnabledMenusByType 获取所有启用的指定位置的菜单
func (dao *MenuDAO) GetAllEnabledMenusByType(tx *gorm.DB, menuPosition string) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := tx.Where("status = ? AND menu_position = ?", model.StatusEnabled, menuPosition).
		Order("seq").
		Find(&menus).Error
	return menus, err
}

// GetMenusByRoleIdAndType 根据角色ID和菜单位置获取菜单列表
func (dao *MenuDAO) GetMenusByRoleIdAndType(tx *gorm.DB, roleId int, menuPosition string) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := tx.Table("p_menu").
		Joins("JOIN p_role_menu ON p_menu.id = p_role_menu.menu_id").
		Where("p_role_menu.role_id = ? AND p_menu.menu_position = ? AND p_menu.status = ?",
			roleId, menuPosition, model.StatusEnabled).
		Order("p_menu.seq").
		Find(&menus).Error
	return menus, err
}
