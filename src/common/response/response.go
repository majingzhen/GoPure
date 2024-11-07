package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 统一响应结构
type Response struct {
	Code int         `json:"code"`           // 响应码
	Msg  string      `json:"msg"`            // 响应消息
	Data interface{} `json:"data,omitempty"` // 响应数据
}

const (
	SUCCESS = 0   // 成功
	ERROR   = 500 // 错误
)

var (
	// 常用响应消息
	SuccessMsg = "操作成功"
	ErrorMsg   = "操作失败"
)

// Result 统一返回结果
func Result(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// Ok 成功响应，无数据
func Ok(c *gin.Context) {
	Result(SUCCESS, SuccessMsg, nil, c)
}

// OkWithMessage 成功响应，自定义消息
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, message, nil, c)
}

// OkWithData 成功响应，带数据
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, SuccessMsg, data, c)
}

// OkWithDetailed 成功响应，自定义消息和数据
func OkWithDetailed(message string, data interface{}, c *gin.Context) {
	Result(SUCCESS, message, data, c)
}

// Fail 失败响应，默认消息
func Fail(c *gin.Context) {
	Result(ERROR, ErrorMsg, nil, c)
}

// FailWithMessage 失败响应，自定义消息
func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, message, nil, c)
}

// FailWithData 失败响应，带数据
func FailWithData(data interface{}, c *gin.Context) {
	Result(ERROR, ErrorMsg, data, c)
}

// FailWithDetailed 失败响应，自定义消息和数据
func FailWithDetailed(message string, data interface{}, c *gin.Context) {
	Result(ERROR, message, data, c)
}
func FailWithDetailedCode(code int, message string, data interface{}, c *gin.Context) {
	Result(code, message, data, c)
}

// FailWithCode 失败响应，自定义错误码和消息
func FailWithCode(code int, message string, c *gin.Context) {
	Result(code, message, nil, c)
}

// JumpView 跳转视图
func JumpView(c *gin.Context, view string) {
	c.HTML(http.StatusOK, view, nil)
}

// JumpViewWithData 跳转视图，带数据
func JumpViewWithData(c *gin.Context, view string, data interface{}) {
	c.HTML(http.StatusOK, view, data)
}

// Unauthorized 未登录
func Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code":     401,
		"message":  "未登录",
		"redirect": "/login",
	})
	c.Abort()
	return
}
