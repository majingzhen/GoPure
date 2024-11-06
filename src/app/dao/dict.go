package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/api/view"
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
func (dao *DictDAO) Page(req view.DictReqPageVO) (*common.PageInfo, error) {
	db := global.GormDao.Model(&model.Dict{})

	// 添加查询条件
	if req.DictType != "" {
		db = db.Where("dict_type = ?", req.DictType)
	}
	if req.DictName != "" {
		db = db.Where("dict_name like ?", "%"+req.DictName+"%")
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	page := common.CreatePageInfo(req.PageNum, req.PageSize)
	if err := db.Count(&page.Total).Error; err != nil {
		return nil, err
	}

	page.Calculate()
	var dataList []*model.Dict
	err := db.Order("seq").Offset(page.Offset).Limit(page.Limit).Find(&dataList).Error
	page.Rows = dataList
	return page, err
}

func (dao *DictDAO) Add(m *model.Dict) error {
	return global.GormDao.Create(m).Error
}

func (dao *DictDAO) Delete(id int) error {
	return global.GormDao.Where("id = ?", id).Delete(&model.Dict{}).Error
}

func (dao *DictDAO) Update(m *model.Dict) error {
	return global.GormDao.Updates(m).Error
}

func (dao *DictDAO) UpdateStatus(id int, status string) error {
	return global.GormDao.Model(&model.Dict{}).Where("id = ?", id).Update("status", status).Error
}
