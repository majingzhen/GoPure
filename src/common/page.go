package common

type PageInfo struct {
	Limit         int         `json:"limit" form:"limit"`                 // 每页大小
	Offset        int         `json:"offset" form:"offset"`               // 当前页码
	PageNum       int         `json:"pageNum" form:"pageNum"`             // 页码
	Total         int64       `json:"total" form:"total"`                 // 总页
	PageSize      int         `json:"pageSize" form:"pageSize"`           // 每页大小
	Rows          interface{} `json:"rows" form:"rows"`                   // 返回列表
	OrderByColumn string      `json:"orderByColumn" form:"orderByColumn"` //排序字段
	IsAsc         string      `json:"isAsc" form:"isAsc"`                 //排序方式
}

type PageView struct {
	OrderByColumn string `json:"orderByColumn" form:"orderByColumn"` //排序字段
	IsAsc         string `json:"isAsc" form:"isAsc"`                 //排序方式
	PageNum       int    `json:"pageNum" form:"pageNum"`             //当前页码
	PageSize      int    `json:"pageSize" form:"pageSize"`           //每页数
}

func CreatePageInfo(pageNum, pageSize int) *PageInfo {
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	return &PageInfo{
		PageNum:  pageNum,
		PageSize: pageSize,
		Limit:    pageSize,
		Offset:   (pageNum - 1) * pageSize,
	}

}

// Calculate 定义一个方法来计算 Limit 和 Offset
func (p *PageInfo) Calculate() {
	if p.PageNum < 1 {
		p.PageNum = 1
	}
	if p.PageSize < 1 {
		p.PageSize = 10
	}
	p.Limit = p.PageSize
	p.Offset = (p.PageNum - 1) * p.PageSize
}
