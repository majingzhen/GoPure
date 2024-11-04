package api

import (
	"fmt"
	"matuto.com/GoPure/src/common/response"
	"matuto.com/GoPure/src/global"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var Upload = new(UploadApi)

type UploadApi struct{}

// Upload 处理文件上传
func (u *UploadApi) Upload(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("获取上传文件失败", c)
		return
	}

	// 检查文件类型
	if !isAllowedFileType(file) {
		response.FailWithMessage("不支持的文件类型", c)
		return
	}

	// 检查文件大小
	if !isAllowedFileSize(file) {
		response.FailWithMessage("文件大小超过限制", c)
		return
	}

	// 生成文件保存路径
	uploadPath := global.Config.Upload.ImagePath
	// 如果不存在上传目录，创建目录
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		if err := os.MkdirAll(uploadPath, 0755); err != nil {
			response.FailWithMessage("创建上传目录失败", c)
			return
		}
	}
	// 生成新的文件名
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadPath, newFileName)

	// 保存文件
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		response.FailWithMessage("保存文件失败", c)
		return
	}
	response.OkWithData("/uploads/images/"+newFileName, c)
	return
}

// 检查文件类型是否允许
func isAllowedFileType(file *multipart.FileHeader) bool {
	allowedTypes := global.Config.Upload.AllowedTypes
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(file.Filename), "."))
	for _, allowedType := range allowedTypes {
		if ext == allowedType {
			return true
		}
	}
	return false
}

// 检查文件大小是否允许
func isAllowedFileSize(file *multipart.FileHeader) bool {
	maxSize := global.Config.Upload.MaxSize
	return file.Size <= maxSize
}
