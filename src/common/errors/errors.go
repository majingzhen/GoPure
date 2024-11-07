package errors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// ErrorCode 错误码
type ErrorCode int

// 错误码定义
const (
	// 系统及错误码
	Success ErrorCode = 200 // 成功
	Failed  ErrorCode = 500 // 失败
	// 客户端错误码
	InvalidParams ErrorCode = 400
	Unauthorized  ErrorCode = 401
	Forbidden     ErrorCode = 403
	NotFound      ErrorCode = 404
	// 服务端错误码
	InternalError ErrorCode = 500
	DBError       ErrorCode = 501
	CacheError    ErrorCode = 502
	RpcError      ErrorCode = 503
)

// CustomError 自定义错误
type CustomError struct {
	Code    ErrorCode // 错误码
	Message string    // 错误消息
	Err     error     // 原始错误
	Stack   string    // 错误堆栈
	Data    any       // 附加数据
}

// Error 实现 error 接口
func (e *CustomError) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code, e.Message)
}

// WithStack 增加错误堆栈
func (e *CustomError) WithStack() *CustomError {
	var sb strings.Builder
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		sb.WriteString(fmt.Sprintf("\n\t%s:%d - %s", file, line, fn.Name()))
	}

	e.Stack = sb.String()
	return e
}

// WithData 增加附加数据
func (e *CustomError) WithData(data any) *CustomError {
	e.Data = data
	return e
}

// New 创建自定义错误
func New(code ErrorCode, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

// Wrap 包装错误
func Wrap(err error, code ErrorCode, message string) *CustomError {
	if err == nil {
		return nil
	}
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Is 判断错误类型
func Is(err error, code ErrorCode) bool {
	var customErr *CustomError
	if errors.As(err, &customErr) {
		return customErr.Code == code
	}
	return false
}
