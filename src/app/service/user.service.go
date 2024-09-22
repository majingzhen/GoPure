package service

import (
	"matuto.com/GoPure/src/app/dao"
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
