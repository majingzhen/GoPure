package api

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/app/service"
	"matuto.com/GoPure/src/common/response"
	"matuto.com/GoPure/src/utils"
)

var Dict = new(DictAPI)

type DictAPI struct{}

// Page 获取字典分页列表
func (api *DictAPI) Page(c *gin.Context) {
	var req view.DictReqPageVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	page, err := service.Dict.Page(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(page, c)
}

// Add 添加字典
func (api *DictAPI) Add(c *gin.Context) {
	var req model.Dict
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.Dict.Add(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Delete 删除字典
func (api *DictAPI) Delete(c *gin.Context) {
	id := utils.GetIntParam(c, "id")
	if id == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	err := service.Dict.Delete(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Update 更新字典
func (api *DictAPI) Update(c *gin.Context) {
	var req view.DictReqEditVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dict := model.Dict{
		Id:       req.Id,
		DictType: req.DictType,
		Status:   req.Status,
		Remark:   req.Remark,
		DictName: req.DictName,
	}
	err := service.Dict.Update(&dict)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func (api *DictAPI) Get(c *gin.Context) {
	id := utils.GetIntParam(c, "id")
	if id == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	dict, err := service.Dict.GetDictById(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(dict, c)
}

// EditStatus 修改字典状态
func (api *DictAPI) EditStatus(c *gin.Context) {
	var req view.DictStatusReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := service.Dict.UpdateStatus(req.Id, req.Status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
