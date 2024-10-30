package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

var Role = new(RoleDAO)

type RoleDAO struct{}

// GetRoleById 根据id获取角色
func (dao *RoleDAO) GetRoleById(tx *gorm.DB, id int) (*model.Role, error) {
	role := &model.Role{}
	err := tx.Where("id = ?", id).First(role).Error
	return role, err
}

// GetRoleByCode 根据角色编码获取角色
func (dao *RoleDAO) GetRoleByCode(tx *gorm.DB, code string) (*model.Role, error) {
	role := &model.Role{}
	err := tx.Where("code = ?", code).First(role).Error
	return role, err
}

// CheckCodeExists 检查角色编码是否已存在
func (dao *RoleDAO) CheckCodeExists(tx *gorm.DB, code string, excludeId ...int) (bool, error) {
	var count int64
	db := tx.Model(&model.Role{}).Where("code = ?", code)
	if len(excludeId) > 0 {
		db = db.Where("id != ?", excludeId[0])
	}
	err := db.Count(&count).Error
	return count > 0, err
}

// Page 获取角色分页列表
func (dao *RoleDAO) Page(pageNum, pageSize int, query map[string]interface{}) (*common.PageInfo, error) {
	db := global.GormDao.Model(&model.Role{})

	// 添加查询条件
	if name, ok := query["name"].(string); ok && name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}
	if code, ok := query["code"].(string); ok && code != "" {
		db = db.Where("code LIKE ?", "%"+code+"%")
	}

	page := common.CreatePageInfo(pageNum, pageSize)
	if err := db.Count(&page.Total).Error; err != nil {
		return nil, err
	}

	page.Calculate()
	var dataList []*model.Role
	err := db.Order("id desc").Offset(page.Offset).Limit(page.Limit).Find(&dataList).Error
	page.Rows = dataList
	return page, err
}

// Add 添加角色
func (dao *RoleDAO) Add(tx *gorm.DB, role *model.Role) error {
	return tx.Create(role).Error
}

// Update 更新角色
func (dao *RoleDAO) Update(tx *gorm.DB, role *model.Role) error {
	return tx.Model(role).Updates(map[string]interface{}{
		"name":        role.Name,
		"code":        role.Code,
		"description": role.Description,
	}).Error
}

// Delete 删除角色
func (dao *RoleDAO) Delete(tx *gorm.DB, ids []int) error {
	return tx.Where("id IN ?", ids).Delete(&model.Role{}).Error
}

// List 获取所有角色列表
func (dao *RoleDAO) List(tx *gorm.DB) ([]*model.Role, error) {
	var roles []*model.Role
	err := tx.Find(&roles).Error
	return roles, err
}
