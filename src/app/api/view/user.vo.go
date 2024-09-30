package view

import "matuto.com/GoPure/src/common"

type UserReqPageVO struct {
	common.PageView
	Account string `json:"account" form:"account"`
}

type LoginUserVO struct {
	Account  string `json:"account"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}
