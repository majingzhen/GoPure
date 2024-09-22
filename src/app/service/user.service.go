package service

import (
	"matuto.com/GoPure/src/app/dao"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/global"
)

type UserService struct {
	dao dao.UserDAO
}

func (service *UserService) GetUserById(id int) (string, error) {
	user, err := service.dao.GetUserById(global.GormDao, id)
	if err != nil {
		return "", err
	}
	return user.Account, nil
}

func (service *UserService) GetByAccount(name string) (*model.User, error) {
	return service.dao.GetByAccount(global.GormDao, name)
}
