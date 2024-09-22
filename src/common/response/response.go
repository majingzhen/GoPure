package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// JumpView 跳转页面
func JumpView(c *gin.Context, view string) {
	c.HTML(http.StatusOK, view, nil)
}

func FailWithMessage(msg string, c *gin.Context) {
	response := Response{
		Code: 1,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, response)
}

func OkWithData(data interface{}, c *gin.Context) {
	response := Response{
		Code: 0,
		Msg:  "Success",
		Data: data,
	}
	c.JSON(http.StatusOK, response)
}

func OkWithMessage(msg string, c *gin.Context) {
	response := Response{
		Code: 0,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, response)
}

func OkWithDetailed(data interface{}, msg string, c *gin.Context) {
	response := Response{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
	c.JSON(http.StatusOK, response)
}

func Unauthorized(c *gin.Context) {
	response := Response{
		Code: 0,
		Msg:  "",
		Data: nil,
	}
	c.JSON(http.StatusUnauthorized, response)
}
