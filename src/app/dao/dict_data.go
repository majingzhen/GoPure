package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

var DictData = new(DictDataDAO)

type DictDataDAO struct{}

// GetDictDataById 根据id获取字典数据
func (dao *DictDataDAO) GetDictDataById(tx *gorm.DB, id int) (*model.DictData, error) {
	dictData := &model.DictData{}
	err := tx.Where("id = ?", id).First(dictData).Error
	return dictData, err
}

// GetDictDataByType 根据字典类型获取字典数据列表
func (dao *DictDataDAO) GetDictDataByType(tx *gorm.DB, dictType string) ([]*model.DictData, error) {
	var dictDataList []*model.DictData
	err := tx.Where("dict_type = ?", dictType).Order("seq").Find(&dictDataList).Error
	return dictDataList, err
}

// Page 获取字典数据分页列表
func (dao *DictDataDAO) Page(pageNum, pageSize int, query map[string]interface{}) (*common.PageInfo, error) {
	db := global.GormDao.Model(&model.DictData{})

	// 添加查询条件
	if dictType, ok := query["dictType"].(string); ok && dictType != "" {
		db = db.Where("dictType = ?", dictType)
	}
	if status, ok := query["status"].(int8); ok {
		db = db.Where("status = ?", status)
	}

	page := common.CreatePageInfo(pageNum, pageSize)
	if err := db.Count(&page.Total).Error; err != nil {
		return nil, err
	}

	page.Calculate()
	var dataList []*model.DictData
	err := db.Order("seq").Offset(page.Offset).Limit(page.Limit).Find(&dataList).Error
	page.Rows = dataList
	return page, err
}
