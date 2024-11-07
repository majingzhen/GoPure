package errors

// HandleError 处理错误并返回自定义错误
func HandleError(err error, code ErrorCode, message string) error {
	if err == nil {
		return nil
	}
	return Wrap(err, code, message)
}

// MustNil 断言错误必须为 nil，否则 panic
func MustNil(err error) {
	if err != nil {
		panic(err)
	}
}

// Assert 断言条件必须为真，否则 panic
func Assert(condition bool, code ErrorCode, message string) {
	if !condition {
		panic(New(code, message))
	}
}
