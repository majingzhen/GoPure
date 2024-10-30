package service

import (
	"errors"
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/dao"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

var User = new(UserService)

type UserService struct{}

// GetUserById 根据id获取用户
func (service *UserService) GetUserById(id int) (*model.User, error) {
	return dao.User.GetUserById(global.GormDao, id)
}

// GetByAccount 根据账号获取用户
func (service *UserService) GetByAccount(account string) (*model.User, error) {
	return dao.User.GetByAccount(global.GormDao, account)
}

// Page 获取用户分页列表
func (service *UserService) Page(req view.UserReqPageVO) (*common.PageInfo, error) {
	return dao.User.Page(req)
}

// Add 添加用户
func (service *UserService) Add(user *model.User, roleIds []int) error {
	return global.GormDao.Transaction(func(tx *gorm.DB) error {
		// 1. 检查账号是否已存在
		exists, err := dao.User.CheckAccountExists(tx, user.Account)
		if err != nil {
			return err
		}
		if exists {
			return errors.New("账号已存在")
		}

		// 2. 保存用户信息
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// 3. 保存用户角色关联
		if len(roleIds) > 0 {
			userRoles := make([]*model.UserRole, len(roleIds))
			for i, roleId := range roleIds {
				userRoles[i] = &model.UserRole{
					UserId: user.Id,
					RoleId: roleId,
				}
			}
			if err := tx.Create(&userRoles).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// Update 更新用户
func (service *UserService) Update(user *model.User, roleIds []int) error {
	return global.GormDao.Transaction(func(tx *gorm.DB) error {
		// 1. 更新用户信息
		if err := tx.Updates(user).Error; err != nil {
			return err
		}

		// 2. 更新用户角色关联
		// 先删除原有关联
		if err := tx.Where("user_id = ?", user.Id).Delete(&model.UserRole{}).Error; err != nil {
			return err
		}
		// 添加新关联
		if len(roleIds) > 0 {
			userRoles := make([]*model.UserRole, len(roleIds))
			for i, roleId := range roleIds {
				userRoles[i] = &model.UserRole{
					UserId: user.Id,
					RoleId: roleId,
				}
			}
			if err := tx.Create(&userRoles).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// Delete 删除用户
func (service *UserService) Delete(ids []int) error {
	return global.GormDao.Transaction(func(tx *gorm.DB) error {
		// 1. 删除用户角色关联
		if err := tx.Where("user_id IN ?", ids).Delete(&model.UserRole{}).Error; err != nil {
			return err
		}

		// 2. 删除用户
		if err := tx.Where("id IN ?", ids).Delete(&model.User{}).Error; err != nil {
			return err
		}

		return nil
	})
}

// UpdateStatus 更新用户状态
func (service *UserService) UpdateStatus(id int, status string) error {
	return global.GormDao.Model(&model.User{}).
		Where("id = ?", id).
		Update("status", status).Error
}

// UpdatePassword 更新用户密码
func (service *UserService) UpdatePassword(id int, password, salt string) error {
	return global.GormDao.Model(&model.User{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"password": password,
			"salt":     salt,
		}).Error
}
