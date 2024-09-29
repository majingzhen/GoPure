package view

import "matuto.com/GoPure/src/common"

type UserReqPageVO struct {
	common.PageView
	Account string `json:"account" form:"account"`
}
