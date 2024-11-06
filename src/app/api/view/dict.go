package view

import "matuto.com/GoPure/src/common"

// DictReqPageVO 字典分页请求
type DictReqPageVO struct {
	common.PageView
	DictName string `json:"dictName" form:"dictName"`
	DictType string `json:"dictType" form:"dictType"`
	Status   string `json:"status" form:"status"`
}

// DictReqEditVO 字典编辑请求
type DictReqEditVO struct {
	Id       int    `json:"id,string" form:"id"`
	DictName string `json:"dictName" form:"dictName"`
	DictType string `json:"dictType" form:"dictType"`
	Remark   string `json:"remark" form:"remark"`
	Status   string `json:"status" form:"status"`
	Seq      int    `json:"seq" form:"seq"`
}

// DictStatusReqVO 更新字典状态请求
type DictStatusReqVO struct {
	Id     int    `json:"id" form:"id"`
	Status string `json:"status" form:"status"`
}
