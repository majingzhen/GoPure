package view

type LoginUserView struct {
	Account    string `json:"account" form:"account"`
	Password   string `json:"password" form:"password"`
	VerifyCode string `json:"code" form:"code"`
	VerifyUuid string `json:"uuid" form:"uuid"`
}

// CaptchaVO 验证码VO
type CaptchaVO struct {
	Img interface{} `json:"img"` //数据内容
	Key string      `json:"key"` //验证码ID
}
