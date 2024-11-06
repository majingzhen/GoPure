package api

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/app/service"
	"matuto.com/GoPure/src/common/response"
	"matuto.com/GoPure/src/utils"
)

var DictData = new(DictDataAPI)

type DictDataAPI struct{}

// List 获取字典数据列表
func (api *DictDataAPI) List(c *gin.Context) {
	dictType := c.Query("dictType")
	if dictType == "" {
		response.FailWithMessage("字典类型不能为空", c)
		return
	}
	// 获取字典数据列表
	dictDataList, err := service.DictData.GetDictDataByType(dictType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(dictDataList, c)
}

// Page 获取字典数据分页列表
func (api *DictDataAPI) Page(c *gin.Context) {
	var req view.DictDataReqPageVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	page, err := service.DictData.Page(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(page, c)
}

// Add 添加字典数据
func (api *DictDataAPI) Add(c *gin.Context) {
	var req view.DictDataReqAddVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data := &model.DictData{
		DictType:        req.DictType,
		DictLabel:       req.DictLabel,
		DictValue:       req.DictValue,
		Status:          req.Status,
		DictExtendValue: req.DictExtendValue,
		Seq:             req.Seq,
	}
	err := service.DictData.Add(data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Edit 修改字典数据
func (api *DictDataAPI) Edit(c *gin.Context) {
	var req view.DictDataReqEditVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data := &model.DictData{
		Id:              req.Id,
		DictType:        req.DictType,
		DictLabel:       req.DictLabel,
		DictValue:       req.DictValue,
		Status:          req.Status,
		DictExtendValue: req.DictExtendValue,
		Seq:             req.Seq,
	}
	err := service.DictData.Update(data)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Delete 删除字典数据
func (api *DictDataAPI) Delete(c *gin.Context) {
	id := utils.GetIntParam(c, "id")
	if id == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	err := service.DictData.Delete(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Get 获取字典数据
func (api *DictDataAPI) Get(c *gin.Context) {
	id := utils.GetIntParam(c, "id")
	if id == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	data, err := service.DictData.GetDictDataById(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(data, c)
}

// EditStatus 修改字典数据状态
func (api *DictDataAPI) EditStatus(c *gin.Context) {
	var req view.DictDataStatusReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.DictData.EditStatus(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
