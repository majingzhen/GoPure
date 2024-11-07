package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"matuto.com/GoPure/src/common/errors"
	"matuto.com/GoPure/src/common/response"
	"matuto.com/GoPure/src/global"
)

// ErrorHandler 通义错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			// 捕获 panic
			if err := recover(); err != nil {
				// 转换自定义错误
				var customErr *errors.CustomError
				switch e := err.(type) {
				case *errors.CustomError:
					customErr = e
				case error:
					customErr = errors.Wrap(e, errors.InternalError, "系统内部错误")
				default:
					customErr = errors.New(errors.InternalError, "系统内部错误")
				}
				// 添加堆栈信息
				customErr.WithStack()
				// 记录错误日志
				global.Logger.Error("panic recovered", zap.Any("err", err), zap.Stack("stack"))
				// 返回错误信息
				response.FailWithDetailedCode(int(customErr.Code), customErr.Message, customErr.Data, c)
				c.Abort()
			}
		}()
		c.Next()
	}
}
