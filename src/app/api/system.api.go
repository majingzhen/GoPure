package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/service"
	"matuto.com/GoPure/src/common/response"
	"matuto.com/GoPure/src/global"
	"matuto.com/GoPure/src/utils"
)

type SystemAPI struct {
	userService service.UserService
}

func (api *SystemAPI) JumpLoginView(c *gin.Context) {
	response.JumpView(c, "login.html")
}

// store 验证码
var store = base64Captcha.DefaultMemStore

// Login 登录
func (api *SystemAPI) Login(c *gin.Context) {
	var loginUserView view.LoginUserView
	_ = c.ShouldBindBodyWithJSON(&loginUserView)

	// 校验验证码
	captcha := VerifyCaptcha(loginUserView.VerifyUuid, loginUserView.VerifyCode)
	if !captcha {
		response.FailWithMessage("验证码错误", c)
		return
	}
	if loginUserView.Account == "" || loginUserView.Password == "" {
		response.FailWithMessage("账号密码不能为空", c)
		return
	}
	byUserName, err := api.userService.GetByAccount(loginUserView.Account)
	if err != nil || byUserName == nil {
		response.FailWithMessage("用户不存在", c)
		return
	}
	// 取加密密码
	hashedPassword := utils.EncryptionPassword(loginUserView.Password, byUserName.Salt)
	if hashedPassword != byUserName.Password {
		global.Logger.Error("登录失败")
		response.FailWithMessage("登录失败", c)
		return
	}

	response.OkWithData(byUserName, c)
}

// CaptchaImage 验证码
func (api *SystemAPI) CaptchaImage(c *gin.Context) {
	//字符,公式,验证码配置
	//定义一个driver
	var driver base64Captcha.Driver
	//创建一个字符串类型的验证码驱动DriverString, DriverChinese :中文驱动
	driverString := base64Captcha.DriverString{
		Height:          40,    //高度
		Width:           100,   //宽度
		NoiseCount:      0,     //干扰数
		ShowLineOptions: 2 | 4, //展示个数
		Length:          4,     //长度
		// Source:          "1234567890qwertyuiplkjhgfdsazxcvbnm", //验证码随机字符串来源
		Source: "1234567890", //验证码随机字符串来源
		BgColor: &color.RGBA{ // 背景颜色
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"}, // 字体
	}
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		// 处理生成验证码时的错误
		response.FailWithMessage("登录失败", c)
	}
	response.OkWithData(&view.CaptchaVO{
		Key: id,
		Img: b64s,
	}, c)
}

// VerifyCaptcha 校验验证码
func VerifyCaptcha(id string, VerifyValue string) bool {
	// 参数说明: id 验证码id, verifyValue 验证码的值, true: 验证成功后是否删除原来的验证码
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
