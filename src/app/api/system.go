package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/app/service"
	"matuto.com/GoPure/src/common/response"
	"matuto.com/GoPure/src/global"
	"matuto.com/GoPure/src/utils"
)

var System = new(SystemAPI)

// SystemAPI 系统api
type SystemAPI struct{}

// JumpHomeView 跳转首页
func (api *SystemAPI) JumpHomeView(c *gin.Context) {
	response.JumpView(c, "index.html")
}

// JumpLoginView 跳转登录页面
func (api *SystemAPI) JumpLoginView(c *gin.Context) {
	response.JumpView(c, "login.html")
}

// store 验证码
var store = base64Captcha.DefaultMemStore

// Login 登录
func (api *SystemAPI) Login(c *gin.Context) {
	var loginUserView view.LoginUserView
	// 接收form表单参数
	_ = c.ShouldBind(&loginUserView)

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
	byUserName, err := service.User.GetByAccount(loginUserView.Account)
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
	// 查询角色
	roles, err := service.Role.GetByUserId(byUserName.Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 查询菜单
	menus, err := service.Menu.GetByUserId(byUserName.Id, model.MENU_POSITION_BACKEND)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 转换成LoginUserVO对象
	loginUserVO := &view.LoginUserVO{
		Account:  byUserName.Account,
		UserName: byUserName.UserName,
		Avatar:   byUserName.Avatar,
		Email:    byUserName.Email,
		Mobile:   byUserName.Mobile,
		Roles:    roles,
		Menus:    menus,
	}
	// 记录登录状态
	session := sessions.Default(c)
	session.Set("user", loginUserVO)
	if err = session.Save(); err != nil {
		global.Logger.Error("登录失败")
		response.FailWithMessage("登录失败", c)
		return
	}
	user := session.Get("user")
	response.OkWithData(user, c)
}

// Logout 退出登录
func (api *SystemAPI) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	if err := session.Save(); err != nil {
		response.FailWithMessage("退出登录失败", c)
		return
	}
	response.OkWithMessage("退出登录成功", c)
}

// GetLoginUser 获取登录用户
func (api *SystemAPI) GetLoginUser(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	response.OkWithData(user, c)
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

// JumpWelcomeView 跳转欢迎页面
func (api *SystemAPI) JumpWelcomeView(c *gin.Context) {
	response.JumpView(c, "welcome.html")
}
