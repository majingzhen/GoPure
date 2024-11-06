package view

import "matuto.com/GoPure/src/common"

// DictDataReqPageVO 字典数据分页请求
type DictDataReqPageVO struct {
	common.PageView
	DictType  string `form:"dictType"`
	Status    string `form:"status"`
	DictLabel string `form:"dictLabel"`
}

// DictDataReqEditVO 字典数据编辑请求
type DictDataReqEditVO struct {
	Id              int    `json:"id,string" form:"id"`
	DictType        string `json:"dictType" form:"dictType"`
	DictLabel       string `json:"dictLabel" form:"dictLabel"`
	DictValue       string `json:"dictValue" form:"dictValue"`
	DictExtendValue string `json:"dictExtendValue" form:"dictExtendValue"`
	Status          string `json:"status" form:"status"`
	Seq             int    `json:"seq,string" form:"seq"`
}

// DictDataReqAddVO 字典数据添加请求
type DictDataReqAddVO struct {
	DictType        string `json:"dictType" form:"dictType"`
	DictLabel       string `json:"dictLabel" form:"dictLabel"`
	DictValue       string `json:"dictValue" form:"dictValue"`
	DictExtendValue string `json:"dictExtendValue" form:"dictExtendValue"`
	Status          string `json:"status" form:"status"`
	Seq             int    `json:"seq" form:"seq"`
}

// DictDataStatusReqVO 字典数据状态更新请求
type DictDataStatusReqVO struct {
	Id     int    `json:"id" form:"id"`
	Status string `json:"status" form:"status"`
}
