package dao

import (
	"gorm.io/gorm"
	"matuto.com/GoPure/src/app/api/view"
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
func (dao *DictDataDAO) Page(req view.DictDataReqPageVO) (*common.PageInfo, error) {
	db := global.GormDao.Model(&model.DictData{})
	// 添加查询条件
	if req.DictType != "" {
		db = db.Where("dict_type = ?", req.DictType)
	}
	if req.DictLabel != "" {
		db = db.Where("dict_label like ?", "%"+req.DictLabel+"%")
	}
	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	page := common.CreatePageInfo(req.PageNum, req.PageSize)
	if err := db.Count(&page.Total).Error; err != nil {
		return nil, err
	}
	page.Calculate()
	var dataList []*model.DictData
	err := db.Order("seq").Offset(page.Offset).Limit(page.Limit).Find(&dataList).Error
	page.Rows = dataList
	return page, err
}

func (dao *DictDataDAO) Add(i *model.DictData) error {
	return global.GormDao.Create(i).Error
}

func (dao *DictDataDAO) Update(data *model.DictData) error {
	return global.GormDao.Updates(data).Error
}

func (dao *DictDataDAO) Delete(id int) error {
	return global.GormDao.Where("id = ?", id).Delete(&model.DictData{}).Error
}

func (dao *DictDataDAO) EditStatus(req view.DictDataStatusReqVO) error {
	return global.GormDao.Model(&model.DictData{}).Where("id = ?", req.Id).Update("status", req.Status).Error
}
