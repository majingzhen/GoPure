package utils

import (
	"fmt"
	"time"
)

// MicrosecondsStr 将时间转换为毫秒字符串
func MicrosecondsStr(d time.Duration) string {
	return fmt.Sprintf("%.2fms", float64(d.Microseconds())/1000)
}
