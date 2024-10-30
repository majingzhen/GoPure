package service

import (
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
func (service *OptionService) Page(pageNum, pageSize int, query map[string]interface{}) (*common.PageInfo, error) {
	return dao.Option.Page(pageNum, pageSize, query)
}
