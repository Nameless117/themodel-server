package apis

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"os"
	"path/filepath"
)

type PDFService struct {
	StoragePath string // PDF 存储路径
	CachePath   string // 预览图缓存路径
}

func NewPDFService(storagePath, cachePath string) *PDFService {
	return &PDFService{
		StoragePath: storagePath,
		CachePath:   cachePath,
	}
}

// GeneratePreview 生成 PDF 预览图
func (s *PDFService) GeneratePreview(pdfPath string) (string, error) {
	// 1. 构建完整的 PDF 文件路径
	fullPDFPath := filepath.Join(s.StoragePath, pdfPath)

	// 2. 检查文件是否存在
	if _, err := os.Stat(fullPDFPath); os.IsNotExist(err) {
		return "", fmt.Errorf("PDF file not found: %s", pdfPath)
	}

	// 3. 构建预览图路径
	previewName := fmt.Sprintf("%s.jpg", filepath.Base(pdfPath))
	previewPath := filepath.Join(s.CachePath, previewName)

	// 4. 检查是否已有缓存
	if _, err := os.Stat(previewPath); err == nil {
		return previewName, nil // 返回已存在的预览图
	}

	// 5. 打开 PDF 文件
	//doc, err := fitz.New(fullPDFPath)
	//if err != nil {
	//	return "", fmt.Errorf("failed to open PDF: %v", err)
	//}
	//defer doc.Close()
	//
	//// 6. 只获取第一页
	//img, err := doc.Image(0)
	//if err != nil {
	//	return "", fmt.Errorf("failed to extract image: %v", err)
	//}
	//
	//// 7. 调整图片大小（可选）
	//resizedImg := resize(img, 800) // 假设宽度调整为 800px
	//
	//// 8. 保存为 JPEG
	//out, err := os.Create(previewPath)
	//if err != nil {
	//	return "", fmt.Errorf("failed to create preview file: %v", err)
	//}
	//defer out.Close()
	//
	//err = jpeg.Encode(out, resizedImg, &jpeg.Options{Quality: 85})
	//if err != nil {
	//	return "", fmt.Errorf("failed to encode preview image: %v", err)
	//}

	return previewName, nil
}

// resize 调整图片大小，保持宽高比
func resize(img image.Image, width int) image.Image {
	bounds := img.Bounds()
	ratio := float64(width) / float64(bounds.Dx())
	height := int(float64(bounds.Dy()) * ratio)

	// 使用imaging库进行缩放
	dst := imaging.Resize(img, width, height, imaging.Lanczos)
	return dst
}

// GetPreview 获取预览图
func (s *PDFService) GetPreview(pdfPath string) (string, error) {
	// 1. 尝试获取已有预览图
	previewName := fmt.Sprintf("%s.jpg", filepath.Base(pdfPath))
	previewPath := filepath.Join(s.CachePath, previewName)

	// 2. 如果预览图存在，直接返回
	if _, err := os.Stat(previewPath); err == nil {
		return previewName, nil
	}

	// 3. 如果不存在，生成预览图
	return s.GeneratePreview(pdfPath)
}
