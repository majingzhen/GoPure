package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/model"
)

var UserRole = new(UserRoleDAO)

type UserRoleDAO struct{}

// GetRolesByUserId 根据用户ID获取角色列表
func (dao *UserRoleDAO) GetRolesByUserId(tx *gorm.DB, userId int) ([]*model.Role, error) {
	var roles []*model.Role
	err := tx.Table("p_role").
		Joins("JOIN p_user_role ON p_role.id = p_user_role.role_id").
		Where("p_user_role.user_id = ?", userId).
		Find(&roles).Error
	return roles, err
}

func (dao *UserRoleDAO) DeleteByUserIds(tx *gorm.DB, ids []int) error {
	return tx.Where("user_id in ?", ids).Delete(&model.UserRole{}).Error
}

func (dao *UserRoleDAO) DeleteByRoleIds(tx *gorm.DB, ids []int) error {
	return tx.Where("role_id in ?", ids).Delete(&model.UserRole{}).Error
}

func (dao *UserRoleDAO) Create(tx *gorm.DB, roles []*model.UserRole) error {
	return tx.Create(&roles).Error
}

func (dao *UserRoleDAO) DeleteByUserId(tx *gorm.DB, id int) error {
	return tx.Where("user_id = ?", id).Delete(&model.UserRole{}).Error
}
