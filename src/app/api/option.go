package api

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/app/service"
	"matuto.com/GoPure/src/common/response"
	"matuto.com/GoPure/src/utils"
)

var Option = new(OptionController)

type OptionController struct{}

// Page 获取选项分页列表
func (c *OptionController) Page(ctx *gin.Context) {
	var req view.OptionReqPageVO
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	page, err := service.Option.Page(req)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(page, ctx)
}

// List 获取选项列表
func (c *OptionController) List(ctx *gin.Context) {
	list, err := service.Option.GetList()
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(list, ctx)
}

// Add 添加选项
func (c *OptionController) Add(ctx *gin.Context) {
	var option model.Option
	if err := ctx.ShouldBindJSON(&option); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if err := service.Option.Add(&option); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.Ok(ctx)
}

// Update 更新选项
func (c *OptionController) Update(ctx *gin.Context) {
	var option view.OptionReqEditVO
	if err := ctx.ShouldBindJSON(&option); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	if option.Id == 0 {
		response.FailWithMessage("参数错误", ctx)
		return
	}
	// vo -> model
	optionData := model.Option{
		Id:             option.Id,
		Key:            option.Key,
		Value:          option.Value,
		Title:          option.Title,
		Identification: option.Identification,
	}
	if err := service.Option.Update(&optionData); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.Ok(ctx)
}

// Delete 删除选项
func (c *OptionController) Delete(ctx *gin.Context) {
	id := utils.GetIntParam(ctx, "id")
	if id == 0 {
		response.FailWithMessage("参数错误", ctx)
		return
	}
	if err := service.Option.Delete(id); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

// Get 获取单个选项
func (c *OptionController) Get(ctx *gin.Context) {
	id := utils.GetIntParam(ctx, "id")
	if id == 0 {
		response.FailWithMessage("参数错误", ctx)
		return
	}
	//
	option, err := service.Option.GetById(id)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(option, ctx)
}
