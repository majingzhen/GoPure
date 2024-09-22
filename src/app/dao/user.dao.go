package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/model"
)

type UserDAO struct{}

// GetUserById 根据id获取用户信息
func (dao *UserDAO) GetUserById(tx *gorm.DB, id int) (*model.User, error) {
	user := &model.User{}
	err := tx.Where("id = ?", id).First(user).Error
	return user, err
}

// GetByUserName 根据账号获取用户信息
func (dao *UserDAO) GetByAccount(tx *gorm.DB, account string) (*model.User, error) {
	user := &model.User{}
	err := tx.Where("account = ?", account).First(user).Error
	return user, err
}
