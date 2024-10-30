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
