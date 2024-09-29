package service

import (
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/dao"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

type UserService struct {
	dao dao.UserDAO
}

// GetUserById 根据id获取用户
func (service *UserService) GetUserById(id int) (string, error) {
	user, err := service.dao.GetUserById(global.GormDao, id)
	if err != nil {
		return "", err
	}
	return user.Account, nil
}

// GetByAccount 根据账号获取用户
func (service *UserService) GetByAccount(name string) (*model.User, error) {
	return service.dao.GetByAccount(global.GormDao, name)
}

// Page 获取用户分页列表
func (service *UserService) Page(req view.UserReqPageVO) (res *common.PageInfo, err error) {
	return service.dao.Page(req)
}
