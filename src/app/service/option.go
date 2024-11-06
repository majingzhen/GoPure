package service

import (
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/dao"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/common"
	"matuto.com/GoPure/src/global"
)

var Option = new(OptionService)

type OptionService struct{}

// GetOptionById 根据id获取选项
func (service *OptionService) GetOptionById(id int) (*model.Option, error) {
	return dao.Option.GetOptionById(global.GormDao, id)
}

// GetOptionByKey 根据key获取选项
func (service *OptionService) GetOptionByKey(key string) (*model.Option, error) {
	return dao.Option.GetOptionByKey(global.GormDao, key)
}

// Page 获取选项分页列表
func (service *OptionService) Page(pageVo view.OptionReqPageVO) (*common.PageInfo, error) {
	return dao.Option.Page(pageVo)
}

func (service *OptionService) GetList(req view.OptionReqListVO) ([]*model.Option, error) {
	return dao.Option.GetList(req)
}

func (service *OptionService) Add(m *model.Option) error {
	return dao.Option.Add(m)
}

func (service *OptionService) Update(m *model.Option) error {
	return dao.Option.Update(m)
}

func (service *OptionService) Delete(id int) error {
	return dao.Option.Delete(id)
}

func (service *OptionService) GetById(id int) (*model.Option, error) {
	return dao.Option.GetById(id)
}
