package service

import (
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/dao"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

var Dict = new(DictService)

type DictService struct{}

// GetDictById 根据id获取字典
func (service *DictService) GetDictById(id int) (*model.Dict, error) {
	return dao.Dict.GetDictById(global.GormDao, id)
}

// GetDictByType 根据字典类型获取字典
func (service *DictService) GetDictByType(dictType string) (*model.Dict, error) {
	return dao.Dict.GetDictByType(global.GormDao, dictType)
}

// Page 获取字典分页列表
func (service *DictService) Page(req view.DictReqPageVO) (*common.PageInfo, error) {
	return dao.Dict.Page(req)
}

func (service *DictService) Add(m *model.Dict) error {
	return dao.Dict.Add(m)
}

func (service *DictService) Delete(id int) error {
	return dao.Dict.Delete(id)
}

func (service *DictService) Update(m *model.Dict) error {
	return dao.Dict.Update(m)
}

func (service *DictService) UpdateStatus(id int, status string) error {
	return dao.Dict.UpdateStatus(id, status)
}
