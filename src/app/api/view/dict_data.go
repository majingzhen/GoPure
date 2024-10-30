package view

// DictDataReqPageVO 字典数据分页请求
type DictDataReqPageVO struct {
	DictType string `form:"dictType"`
	Status   string `form:"status"`
	PageNum  int    `form:"page" binding:"required"`
	PageSize int    `form:"limit" binding:"required"`
}
