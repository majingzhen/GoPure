package utils

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
)

// GenerateSalt 生成指定长度的随机盐值（简单版）
func GenerateSalt(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		result[i] = charset[num.Int64()]
	}

	return string(result)
}

// GenerateSaltBase64 生成Base64编码的盐值
func GenerateSaltBase64(length int) (string, error) {
	bytes := make([]byte, length)

	// 使用crypto/rand生成随机字节
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("生成随机字节失败: %v", err)
	}

	// 返回base64编码的字符串
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// GenerateSaltHex 生成16进制编码的盐值
func GenerateSaltHex(length int) (string, error) {
	bytes := make([]byte, length)

	// 使用crypto/rand生成随机字节
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("生成随机字节失败: %v", err)
	}

	// 返回16进制编码的字符串
	return hex.EncodeToString(bytes), nil
}

// GenerateSaltWithPrefix 生成带前缀的盐值
func GenerateSaltWithPrefix(prefix string, length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("生成随机字节失败: %v", err)
	}

	// 组合前缀和随机盐值
	return fmt.Sprintf("%s%s", prefix, hex.EncodeToString(bytes)), nil
}

// SaltConfig 盐值配置
type SaltConfig struct {
	Length     int    // 盐值长度
	Prefix     string // 前缀
	UseBase64  bool   // 是否使用Base64编码
	UseHex     bool   // 是否使用16进制编码
	Complexity int    // 复杂度 1-简单 2-中等 3-复杂
}

// GenerateSaltWithConfig 根据配置生成盐值
func GenerateSaltWithConfig(config SaltConfig) (string, error) {
	if config.Length <= 0 {
		return "", fmt.Errorf("盐值长度必须大于0")
	}

	var salt string
	var err error

	// 根据复杂度生成不同字符集的盐值
	switch config.Complexity {
	case 1:
		salt = GenerateSalt(config.Length)
	case 2:
		salt, err = GenerateSaltBase64(config.Length)
	case 3:
		salt, err = GenerateSaltHex(config.Length)
	default:
		salt = GenerateSalt(config.Length)
	}

	if err != nil {
		return "", err
	}

	// 如果有前缀，添加前缀
	if config.Prefix != "" {
		salt = config.Prefix + salt
	}

	return salt, nil
}
