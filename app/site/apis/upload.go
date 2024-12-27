package apis

import (
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

// FileTypeLimit 文件类型限制
type FileTypeLimit struct {
	Extensions []string // 允许的扩展名
	MaxSize    int64    // 最大大小（字节）
}

// UploadConfig 上传配置
type UploadConfig struct {
	SavePath       string                   // 文件保存路径
	FileTypeLimits map[string]FileTypeLimit // 不同类型文件的限制
}

// UploadResponse 上传响应
type UploadResponse struct {
	FileName string `json:"fileName"` // 文件名
	FileUrl  string `json:"fileUrl"`  // 文件访问URL
}

// LocalUploader 本地文件上传器
type LocalUploader struct {
	Config UploadConfig
}

// NewLocalUploader 创建上传器实例
func NewLocalUploader(config UploadConfig) *LocalUploader {
	return &LocalUploader{
		Config: config,
	}
}

// Upload 上传文件到本地
func (u *LocalUploader) Upload(file *multipart.FileHeader) (*UploadResponse, error) {
	// 1. 获取文件类型限制
	ext := strings.ToLower(path.Ext(file.Filename))
	typeLimit, err := u.getTypeLimit(ext)
	if err != nil {
		return nil, err
	}

	// 2. 检查文件大小
	if file.Size > typeLimit.MaxSize {
		return nil, fmt.Errorf("文件大小超过限制: %d MB", typeLimit.MaxSize/1024/1024)
	}

	// 3. 生成文件保存路径
	savePath := u.generateSavePath()
	if err := os.MkdirAll(savePath, 0755); err != nil {
		return nil, fmt.Errorf("创建目录失败: %v", err)
	}

	// 4. 生成新的文件名
	newFileName := u.generateFileName(file.Filename)
	fullPath := path.Join(savePath, newFileName)

	// 5. 保存文件
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("打开上传文件失败: %v", err)
	}
	defer src.Close()

	dst, err := os.Create(fullPath)
	if err != nil {
		return nil, fmt.Errorf("创建目标文件失败: %v", err)
	}
	defer dst.Close()

	// 6. 复制文件内容
	buffer := make([]byte, 1024*1024) // 1MB buffer
	for {
		n, err := src.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("读取文件失败: %v", err)
		}
		if _, err := dst.Write(buffer[:n]); err != nil {
			return nil, fmt.Errorf("写入文件失败: %v", err)
		}
	}

	// 7. 构建文件访问URL
	fileUrl := path.Join("/uploads", time.Now().Format("2006/01/02"), newFileName)

	return &UploadResponse{
		FileName: newFileName,
		FileUrl:  fileUrl,
	}, nil
}

// getTypeLimit 获取文件类型限制
func (u *LocalUploader) getTypeLimit(fileExt string) (*FileTypeLimit, error) {
	for _, typeLimit := range u.Config.FileTypeLimits {
		for _, ext := range typeLimit.Extensions {
			if strings.ToLower(ext) == fileExt {
				return &typeLimit, nil
			}
		}
	}
	return nil, fmt.Errorf("不支持的文件类型: %s", fileExt)
}

// generateSavePath 生成文件保存路径
func (u *LocalUploader) generateSavePath() string {
	datePath := time.Now().Format("2006/01/02")
	return path.Join(u.Config.SavePath, datePath)
}

// generateFileName 生成新的文件名
func (u *LocalUploader) generateFileName(originalName string) string {
	ext := path.Ext(originalName)
	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	return fileName
}
