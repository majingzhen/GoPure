package service

import (
	"matuto.com/GoPure/src/app/api/view"
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
func (service *DictDataService) Page(req view.DictDataReqPageVO) (*common.PageInfo, error) {
	return dao.DictData.Page(req)
}

func (service *DictDataService) Add(i *model.DictData) error {
	return dao.DictData.Add(i)
}

func (service *DictDataService) Update(data *model.DictData) error {
	return dao.DictData.Update(data)
}

func (service *DictDataService) Delete(id int) error {
	return dao.DictData.Delete(id)
}

func (service *DictDataService) EditStatus(req view.DictDataStatusReqVO) error {
	return dao.DictData.EditStatus(req)
}
