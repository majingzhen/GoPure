package api

import (
	"github.com/gin-gonic/gin"
	"matuto.com/GoPure/src/app/api/view"
	"matuto.com/GoPure/src/app/model"
	"matuto.com/GoPure/src/app/service"
	"matuto.com/GoPure/src/common/response"
	"matuto.com/GoPure/src/utils"
)

var User = new(UserAPI)

type UserAPI struct{}

// JumpUserView 跳转用户管理页面
func (api *UserAPI) JumpUserView(c *gin.Context) {
	response.JumpView(c, "user/index.html")
}

// JumpUserAddView 跳转添加用户页面
func (api *UserAPI) JumpUserAddView(c *gin.Context) {
	response.JumpView(c, "user/add.html")
}

// JumpUserEditView 跳转编辑用户页面
func (api *UserAPI) JumpUserEditView(c *gin.Context) {
	response.JumpView(c, "user/edit.html")
}

// Page 获取用户分页
func (api *UserAPI) Page(c *gin.Context) {
	var req view.UserReqPageVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	page, err := service.User.Page(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(page, c)
}

// Get 获取用户详情
func (api *UserAPI) Get(c *gin.Context) {
	id := utils.GetIntParam(c, "id")
	if id == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}
	user, err := service.User.GetUserById(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取用户角色
	roles, err := service.Role.GetByUserId(id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 构建返回数据
	roleIds := make([]int, len(roles))
	for i, role := range roles {
		roleIds[i] = role.Id
	}
	userVO := &view.UserVO{
		User:    user,
		RoleIds: roleIds,
	}
	response.OkWithData(userVO, c)
}

// Add 添加用户
func (api *UserAPI) Add(c *gin.Context) {
	var req view.UserAddReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 生成盐值和密码
	salt := utils.GenerateSalt(16)
	hashedPassword := utils.EncryptionPassword(req.Password, salt)
	// 构建用户对象
	user := &model.User{
		Account:  req.Account,
		UserName: req.UserName,
		Password: hashedPassword,
		Salt:     salt,
		Status:   req.Status,
		Mobile:   req.Mobile,
		Email:    req.Email,
		Sex:      req.Sex,
		Remark:   req.Remark,
	}
	// 保存用户
	err := service.User.Add(user, req.RoleIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Update 更新用户
func (api *UserAPI) Update(c *gin.Context) {
	var req view.UserUpdateReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取原用户信息
	user, err := service.User.GetUserById(req.Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 判断用户是否为管理员
	roles, err := service.Role.GetByUserId(req.Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(roles) > 0 {
		// 如果是管理员用户，则不允许修改
		for _, role := range roles {
			if role.Code == model.RoleAdmin {
				response.FailWithMessage("管理员用户不允许修改", c)
				return
			}
		}
	}
	// 更新基本信息
	user.UserName = req.UserName
	user.Status = req.Status
	user.Mobile = req.Mobile
	user.Email = req.Email
	user.Sex = req.Sex
	user.Remark = req.Remark
	// 更新用户
	err = service.User.Update(user, req.RoleIds)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// Delete 删除用户
func (api *UserAPI) Delete(c *gin.Context) {
	var req view.UserDeleteReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	loginUserId, _ := c.Get("userId")
	err := service.User.Delete(loginUserId.(int), req.Ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// UpdateStatus 更新用户状态
func (api *UserAPI) UpdateStatus(c *gin.Context) {
	var req view.UserStatusReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	loginUserId, _ := c.Get("userId")
	if loginUserId.(int) == req.Id {
		response.FailWithMessage("不能修改自己的状态", c)
		return
	}
	err := service.User.UpdateStatus(req.Id, req.Status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

// ResetPassword 重置密码
func (api *UserAPI) ResetPassword(c *gin.Context) {
	var req view.UserResetPwdReqVO
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取用户信息
	user, err := service.User.GetUserById(req.Id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if user == nil {
		response.FailWithMessage("用户不存在", c)
		return
	}

	// 生成新的盐值和密码
	salt := utils.GenerateSalt(16)
	hashedPassword := utils.EncryptionPassword(req.Password, salt)

	// 更新密码
	err = service.User.UpdatePassword(req.Id, hashedPassword, salt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
