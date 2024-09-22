package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/model"
)

type UserDAO struct{}

func (dao *UserDAO) GetUserById(tx *gorm.DB, id int) (*model.User, error) {
	user := &model.User{}
	err := tx.Where("id = ?", id).First(user).Error
	return user, err
}
