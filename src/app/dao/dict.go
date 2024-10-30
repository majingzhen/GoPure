package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

var Dict = new(DictDAO)

type DictDAO struct{}

// GetDictById 根据id获取字典信息
func (dao *DictDAO) GetDictById(tx *gorm.DB, id int) (*model.Dict, error) {
	dict := &model.Dict{}
	err := tx.Where("id = ?", id).First(dict).Error
	return dict, err
}

// GetDictByType 根据字典类型获取字典信息
func (dao *DictDAO) GetDictByType(tx *gorm.DB, dictType string) (*model.Dict, error) {
	dict := &model.Dict{}
	err := tx.Where("dictType = ?", dictType).First(dict).Error
	return dict, err
}

// Page 获取字典分页列表
func (dao *DictDAO) Page(pageNum, pageSize int, query map[string]interface{}) (*common.PageInfo, error) {
	db := global.GormDao.Model(&model.Dict{})

	// 添加查询条件
	if dictType, ok := query["dictType"].(string); ok && dictType != "" {
		db = db.Where("dictType LIKE ?", "%"+dictType+"%")
	}
	if dictName, ok := query["dictName"].(string); ok && dictName != "" {
		db = db.Where("dictName LIKE ?", "%"+dictName+"%")
	}
	if status, ok := query["status"].(int8); ok {
		db = db.Where("status = ?", status)
	}

	page := common.CreatePageInfo(pageNum, pageSize)
	if err := db.Count(&page.Total).Error; err != nil {
		return nil, err
	}

	page.Calculate()
	var dataList []*model.Dict
	err := db.Order("seq").Offset(page.Offset).Limit(page.Limit).Find(&dataList).Error
	page.Rows = dataList
	return page, err
}
