package service

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/dao"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/common/errors"
	"matuto.com/GoPure/src/global"
)

var User = new(UserService)

type UserService struct{}

// GetUserById 根据id获取用户
func (service *UserService) GetUserById(id int) (*model.User, error) {
	user, err := dao.User.GetUserById(global.GormDao, id)
	if err != nil {
		return nil, errors.HandleError(err, errors.DBError, "获取用户失败")
	}
	if user == nil {
		return nil, errors.ErrUserNotFound.WithData(gin.H{"id": id})
	}
	return user, nil
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
	// 1. 检查账号是否已存在
	exists, err := dao.User.CheckAccountExists(user.Account)
	if err != nil {
		return errors.HandleError(err, errors.DBError, "检查用户失败")
	}
	if exists {
		return errors.ErrUserExist.WithData(gin.H{"account": user.Account})
	}

	tx := global.GormDao.Begin()
	if err := dao.User.Create(tx, user); err != nil {
		tx.Rollback()
		return err
	}

	if len(roleIds) > 0 {
		userRoles := make([]*model.UserRole, len(roleIds))
		for i, roleId := range roleIds {
			userRoles[i] = &model.UserRole{
				UserId: user.Id,
				RoleId: roleId,
			}
		}
		if err := dao.UserRole.Create(tx, userRoles); err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

// Update 更新用户
func (service *UserService) Update(user *model.User, roleIds []int) error {
	if user.Id == model.AdminId {
		return errors.ErrAdminIsNotEdit
	}
	tx := global.GormDao.Begin()

	if err := dao.User.Update(tx, user); err != nil {
		tx.Rollback()
		return err
	}
	// 删除原有角色关联
	if err := dao.UserRole.DeleteByUserId(tx, user.Id); err != nil {
		tx.Rollback()
		return err
	}
	// 添加新角色关联
	if len(roleIds) > 0 {
		userRoles := make([]*model.UserRole, len(roleIds))
		for i, roleId := range roleIds {
			userRoles[i] = &model.UserRole{
				UserId: user.Id,
				RoleId: roleId,
			}
		}
		if err := dao.UserRole.Create(tx, userRoles); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()
	return nil
}

// Delete 删除用户
func (service *UserService) Delete(loginUserId int, ids []int) error {
	// 不允许删除自己
	for _, id := range ids {
		if id == loginUserId {
			return errors.ErrSelfDelete
		}
		if id == model.AdminId {
			return errors.ErrAdminIsNotDelete
		}
	}
	// 删除用户
	tx := global.GormDao.Begin()
	if err := dao.User.Delete(tx, ids); err != nil {
		tx.Rollback()
		return err
	}

	// 删除用户角色关联
	if err := dao.UserRole.DeleteByUserIds(tx, ids); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// UpdateStatus 更新用户状态
func (service *UserService) UpdateStatus(id int, status string) error {
	// 不允许禁用管理员
	if id == model.AdminId {
		return errors.ErrAdminIsNotEdit
	}
	return dao.User.UpdateStatus(id, status)
}

// UpdatePassword 更新用户密码
func (service *UserService) UpdatePassword(id int, password, salt string) error {
	return dao.User.UpdatePassword(id, password, salt)
}

// CheckPermission 检查用户是否有权限
func (service *UserService) CheckPermission(loginUserId int, path string) bool {
	// 如果是超级管理员角色,直接返回true
	if loginUserId == model.AdminId {
		return true
	}
	// 获取用户角色
	roles, _ := Role.GetByUserId(loginUserId)
	if len(roles) == 0 {
		return false
	}
	// 获取角色菜单
	roleIds := make([]int, len(roles))
	for i, role := range roles {
		roleIds[i] = role.Id
	}
	menus, _ := Menu.GetByRoleIds(roleIds)
	// 判断是否有权限
	for _, menu := range menus {
		if menu.Url == path {
			return true
		}
	}
	return false
}
