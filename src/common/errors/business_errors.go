package errors

// 业务错误码定义
const (
	// 用户模块
	UserNotFound     ErrorCode = 10001
	UserExist        ErrorCode = 10002
	PasswordInvalid  ErrorCode = 10003
	AdminIsNotEdit   ErrorCode = 10004
	SelfIsNotDelete  ErrorCode = 10005
	AdminIsNotDelete ErrorCode = 10006

	// 角色模块
	RoleNotFound ErrorCode = 20001
	RoleExist    ErrorCode = 20002

	// 菜单模块
	MenuNotFound ErrorCode = 30001
	MenuExist    ErrorCode = 30002

	// 字典模块
	DictNotFound ErrorCode = 40001
	DictExist    ErrorCode = 40002
)

// 预定义业务错误
var (
	// 用户模块
	ErrUserNotFound     = New(UserNotFound, "用户不存在")
	ErrUserExist        = New(UserExist, "用户已存在")
	ErrPasswordInvalid  = New(PasswordInvalid, "密码不正确")
	ErrAdminIsNotEdit   = New(AdminIsNotEdit, "管理员不可编辑")
	ErrSelfDelete       = New(SelfIsNotDelete, "不能删除自己")
	ErrAdminIsNotDelete = New(AdminIsNotDelete, "管理员不可删除")

	// 角色模块
	ErrRoleNotFound = New(RoleNotFound, "角色不存在")
	ErrRoleExist    = New(RoleExist, "角色已存在")

	// 菜单模块
	ErrMenuNotFound = New(MenuNotFound, "菜单不存在")
	ErrMenuExist    = New(MenuExist, "菜单已存在")

	// 字典模块
	ErrDictNotFound = New(DictNotFound, "字典不存在")
	ErrDictExist    = New(DictExist, "字典已存在")
)
