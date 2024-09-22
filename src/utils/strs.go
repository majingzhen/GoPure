package utils

import (
	"github.com/gogf/gf/crypto/gmd5"
	uuid "github.com/satori/go.uuid"

	"strings"
)

// ToTitle 将下划线命名转换为首字母大写的驼峰命名
func ToTitle(underscoreName string) string {
	words := strings.Split(strings.ToLower(underscoreName), "_")
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToTitle(words[i])
	}
	return strings.Join(words, "")
}

// ToCamelCase 将下划线命名转换为驼峰命名
func ToCamelCase(underscoreName string) string {
	words := strings.Split(strings.ToLower(underscoreName), "_")
	for i := 1; i < len(words); i++ {
		words[i] = strings.ToTitle(words[i])
	}
	return strings.Join(words, "")
}

// EncryptionPassword 加密密码
func EncryptionPassword(password, salt string) string {
	// 将盐与密码结合后进行哈希计算
	hashedPassword := gmd5.MustEncryptString(password + salt)
	return hashedPassword
}

// ReplaceEach 将字符串str中的所有与搜索列表中的子串匹配的部分替换为对应的替换列表中的子串
func ReplaceEach(str string, searchList []string, replacementList []string) string {
	if str == "" || searchList == nil || replacementList == nil {
		return ""
	}

	if len(searchList) != len(replacementList) {
		panic("Search and Replace array lengths don't match")
	}

	for i := 0; i < len(searchList); i++ {
		str = strings.Replace(str, searchList[i], replacementList[i], -1)
	}

	return str
}

// GenUID 生成UUID
func GenUID() string {
	uid := uuid.NewV4()
	uidStr := strings.ReplaceAll(uid.String(), "-", "")
	return uidStr
}

// EndsWithIgnoreCase 判断字符串s是否以字符串t结尾，忽略大小写
func EndsWithIgnoreCase(s, t string) bool {
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(t))
}

// BeginsWithIgnoreCase 判断字符串s是否以字符串t开头，忽略大小写
func BeginsWithIgnoreCase(s, t string) bool {
	return strings.HasPrefix(strings.ToLower(s), strings.ToLower(t))
}

// StrContains 判断字符串s是否包含字符串t
func StrContains(s, t string) bool {
	return strings.Contains(s, t)
}
