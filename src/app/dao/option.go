package dao

import (
	"gorm.io/gorm"
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
func (dao *OptionDAO) Page(pageNum, pageSize int, query map[string]interface{}) (*common.PageInfo, error) {
	db := global.GormDao.Model(&model.Option{})

	// 添加查询条件
	if key, ok := query["key"].(string); ok && key != "" {
		db = db.Where("`key` LIKE ?", "%"+key+"%")
	}
	if title, ok := query["title"].(string); ok && title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if identification, ok := query["identification"].(string); ok && identification != "" {
		db = db.Where("identification = ?", identification)
	}

	page := common.CreatePageInfo(pageNum, pageSize)
	if err := db.Count(&page.Total).Error; err != nil {
		return nil, err
	}
	page.Calculate()
	var dataList []*model.Option
	err := db.Offset(page.Offset).Limit(page.Limit).Find(&dataList).Error
	page.Rows = dataList
	return page, err
}
