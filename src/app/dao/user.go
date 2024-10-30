package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

var User = new(UserDAO)

type UserDAO struct{}

// GetUserById 根据id获取用户信息
func (dao *UserDAO) GetUserById(tx *gorm.DB, id int) (*model.User, error) {
	user := &model.User{}
	err := tx.Where("id = ?", id).First(user).Error
	return user, err
}

// GetByAccount 根据账号获取用户信息
func (dao *UserDAO) GetByAccount(tx *gorm.DB, account string) (*model.User, error) {
	user := &model.User{}
	err := tx.Where("account = ?", account).First(user).Error
	return user, err
}

// Page 获取用户分页列表
//
// req: UserReqPageVO
func (dao *UserDAO) Page(req view.UserReqPageVO) (page *common.PageInfo, err error) {
	db := global.GormDao.Model(&model.User{})
	if req.Account != "" {
		db.Where("account like ?", "%"+req.Account+"%")
	}
	if req.Status != "" {
		db.Where("status = ?", req.Status)
	}
	if req.UserName != "" {
		db.Where("user_name like ?", "%"+req.UserName+"%")
	}
	page = common.CreatePageInfo(req.PageNum, req.PageSize)
	if err = db.Count(&page.Total).Error; err != nil {
		return
	}
	// 计算分页信息
	page.Calculate()
	var dataList []*model.User
	err = db.Offset(page.Offset).Limit(page.Limit).Find(&dataList).Error
	page.Rows = dataList
	return
}

// CheckAccountExists 检查账号是否已存在
func (dao *UserDAO) CheckAccountExists(tx *gorm.DB, account string) (bool, error) {
	var count int64
	err := tx.Model(&model.User{}).
		Where("account = ?", account).
		Count(&count).Error
	return count > 0, err
}
