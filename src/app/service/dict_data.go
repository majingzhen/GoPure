package service

import (
	"matuto.com/GoPure/src/app/dao"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

var DictData = new(DictDataService)

type DictDataService struct{}

// GetDictDataById 根据id获取字典数据
func (service *DictDataService) GetDictDataById(id int) (*model.DictData, error) {
	return dao.DictData.GetDictDataById(global.GormDao, id)
}

// GetDictDataByType 根据字典类型获取字典数据列表
func (service *DictDataService) GetDictDataByType(dictType string) ([]*model.DictData, error) {
	return dao.DictData.GetDictDataByType(global.GormDao, dictType)
}

// Page 获取字典数据分页列表
func (service *DictDataService) Page(pageNum, pageSize int, query map[string]interface{}) (*common.PageInfo, error) {
	return dao.DictData.Page(pageNum, pageSize, query)
}
