package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetIntParam 获取整型参数
// 支持从URL路径参数、查询字符串、表单参数中获取
func GetIntParam(c *gin.Context, key string) int {
	// 1. 尝试从路径参数获取
	if param := c.Param(key); param != "" {
		if val, err := strconv.Atoi(param); err == nil {
			return val
		}
	}

	// 2. 尝试从查询字符串获取
	if param := c.Query(key); param != "" {
		if val, err := strconv.Atoi(param); err == nil {
			return val
		}
	}

	// 3. 尝试从表单参数获取
	if param := c.PostForm(key); param != "" {
		if val, err := strconv.Atoi(param); err == nil {
			return val
		}
	}

	// 4. 尝试从JSON body获取
	var jsonMap map[string]interface{}
	if err := c.ShouldBindJSON(&jsonMap); err == nil {
		if val, exists := jsonMap[key]; exists {
			switch v := val.(type) {
			case float64:
				return int(v)
			case int:
				return v
			case string:
				if intVal, err := strconv.Atoi(v); err == nil {
					return intVal
				}
			}
		}
	}

	return 0
}

// GetStringParam 获取字符串参数
// 支持从URL路径参数、查询字符串、表单参数中获取
func GetStringParam(c *gin.Context, key string) string {
	// 1. 尝试从路径参数获取
	if param := c.Param(key); param != "" {
		return param
	}

	// 2. 尝试从查询字符串获取
	if param := c.Query(key); param != "" {
		return param
	}

	// 3. 尝试从表单参数获取
	if param := c.PostForm(key); param != "" {
		return param
	}

	// 4. 尝试从JSON body获取
	var jsonMap map[string]interface{}
	if err := c.ShouldBindJSON(&jsonMap); err == nil {
		if val, exists := jsonMap[key]; exists {
			if strVal, ok := val.(string); ok {
				return strVal
			}
		}
	}

	return ""
}

// GetIntArrayParam 获取整型数组参数
// 支持从查询字符串、表单参数中获取
func GetIntArrayParam(c *gin.Context, key string) []int {
	var result []int

	// 1. 尝试从查询字符串获取
	if values := c.QueryArray(key); len(values) > 0 {
		for _, v := range values {
			if val, err := strconv.Atoi(v); err == nil {
				result = append(result, val)
			}
		}
		return result
	}

	// 2. 尝试从表单参数获取
	if values := c.PostFormArray(key); len(values) > 0 {
		for _, v := range values {
			if val, err := strconv.Atoi(v); err == nil {
				result = append(result, val)
			}
		}
		return result
	}

	// 3. 尝试从JSON body获取
	var jsonMap map[string]interface{}
	if err := c.ShouldBindJSON(&jsonMap); err == nil {
		if val, exists := jsonMap[key]; exists {
			switch v := val.(type) {
			case []interface{}:
				for _, item := range v {
					switch num := item.(type) {
					case float64:
						result = append(result, int(num))
					case int:
						result = append(result, num)
					case string:
						if intVal, err := strconv.Atoi(num); err == nil {
							result = append(result, intVal)
						}
					}
				}
			}
		}
	}

	return result
}

// GetStringArrayParam 获取字符串数组参数
// 支持从查询字符串、表单参数中获取
func GetStringArrayParam(c *gin.Context, key string) []string {
	// 1. 尝试从查询字符串获取
	if values := c.QueryArray(key); len(values) > 0 {
		return values
	}

	// 2. 尝试从表单参数获取
	if values := c.PostFormArray(key); len(values) > 0 {
		return values
	}

	// 3. 尝试从JSON body获取
	var jsonMap map[string]interface{}
	if err := c.ShouldBindJSON(&jsonMap); err == nil {
		if val, exists := jsonMap[key]; exists {
			switch v := val.(type) {
			case []interface{}:
				result := make([]string, 0, len(v))
				for _, item := range v {
					if str, ok := item.(string); ok {
						result = append(result, str)
					}
				}
				return result
			}
		}
	}

	return nil
}
