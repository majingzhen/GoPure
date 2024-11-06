package view

import "matuto.com/GoPure/src/common"

// OptionReqPageVO 选项分页请求
type OptionReqPageVO struct {
	common.PageView
	Key string `json:"key" form:"key"`
}

// OptionReqEditVO 选项编辑请求
type OptionReqEditVO struct {
	Id             int    `json:"id,string" form:"id"`
	Key            string `json:"key" form:"key"`
	Value          string `json:"value" form:"value"`
	Title          string `json:"title" form:"title"`
	Identification string `json:"identification" form:"identification"`
}

// OptionReqListVO 选项列表请求
type OptionReqListVO struct {
	Key            string `json:"key" form:"key"`
	Identification string `json:"identification" form:"identification"`
}
