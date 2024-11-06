package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

var Option = new(OptionDAO)

type OptionDAO struct{}

// GetOptionById 根据id获取选项信息
func (dao *OptionDAO) GetOptionById(tx *gorm.DB, id int) (*model.Option, error) {
	option := &model.Option{}
	err := tx.Where("id = ?", id).First(option).Error
	return option, err
}

// GetOptionByKey 根据key获取选项信息
func (dao *OptionDAO) GetOptionByKey(tx *gorm.DB, key string) (*model.Option, error) {
	option := &model.Option{}
	err := tx.Where("`key` = ?", key).First(option).Error
	return option, err
}

// Page 获取选项分页列表
func (dao *OptionDAO) Page(pageVo view.OptionReqPageVO) (*common.PageInfo, error) {
	db := global.GormDao.Model(&model.Option{})

	// 添加查询条件
	if pageVo.Key != "" {
		db = db.Where("`key` like ?", "%"+pageVo.Key+"%")
	}
	page := common.CreatePageInfo(pageVo.PageNum, pageVo.PageSize)
	if err := db.Count(&page.Total).Error; err != nil {
		return nil, err
	}
	page.Calculate()
	var dataList []*model.Option
	err := db.Offset(page.Offset).Limit(page.Limit).Find(&dataList).Error
	page.Rows = dataList
	return page, err
}

// GetList 获取选项列表
func (dao *OptionDAO) GetList() ([]*model.Option, error) {
	var list []*model.Option
	err := global.GormDao.Find(&list).Error
	return list, err
}

// Add 添加选项
func (dao *OptionDAO) Add(m *model.Option) error {
	return global.GormDao.Create(m).Error
}

// Update 更新选项
func (dao *OptionDAO) Update(m *model.Option) error {
	return global.GormDao.Updates(m).Error
}

// Delete 删除选项
func (dao *OptionDAO) Delete(id int) error {
	return global.GormDao.Delete(&model.Option{}, id).Error
}

// GetById 根据id获取选项
func (dao *OptionDAO) GetById(id int) (*model.Option, error) {
	return dao.GetOptionById(global.GormDao, id)
}
