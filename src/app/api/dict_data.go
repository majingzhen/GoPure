package api

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/service"
	"matuto.com/GoPure/src/common/response"
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

	page, err := service.DictData.Page(req.PageNum, req.PageSize, map[string]interface{}{
		"dictType": req.DictType,
		"status":   req.Status,
	})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(page, c)
}
