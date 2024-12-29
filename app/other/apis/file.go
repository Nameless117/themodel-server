package apis

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/utils"
	"github.com/google/uuid"

	"go-admin/common/file_store"
)

type FileResponse struct {
	Size     int64  `json:"size"`
	Path     string `json:"path"`
	FullPath string `json:"full_path"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

const path = "static/uploadfile/"

type File struct {
	api.Api
}

// UploadFile 上传图片
// @Summary 上传图片
// @Description 获取JSON
// @Tags 公共接口
// @Accept multipart/form-data
// @Param type query string true "type" (1：单图，2：多图, 3：base64图片)
// @Param file formData file true "file"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/public/uploadFile [post]
// @Security Bearer
func (e File) UploadFile(c *gin.Context) {
	e.MakeContext(c)
	tag, _ := c.GetPostForm("type")
	urlPrefix := fmt.Sprintf("%s://%s/", "http", c.Request.Host)
	var fileResponse FileResponse

	switch tag {
	case "1": // 单图
		var done bool
		fileResponse, done = e.singleFile(c, fileResponse, urlPrefix)
		if done {
			return
		}
		e.OK(fileResponse, "上传成功")
		return
	case "2": // 多图
		multipartFile := e.multipleFile(c, urlPrefix)
		e.OK(multipartFile, "上传成功")
		return
	case "3": // base64
		fileResponse = e.baseImg(c, fileResponse, urlPrefix)
		e.OK(fileResponse, "上传成功")
	case "4": // pdf
		var done bool
		fileResponse, done = e.uploadPDF(c, fileResponse, urlPrefix)
		if done {
			return
		}
		e.OK(fileResponse, "上传成功")
	default:
		var done bool
		fileResponse, done = e.singleFile(c, fileResponse, urlPrefix)
		if done {
			return
		}
		e.OK(fileResponse, "上传成功")
		return
	}

}

func (e File) baseImg(c *gin.Context, fileResponse FileResponse, urlPerfix string) FileResponse {
	files, _ := c.GetPostForm("file")
	file2list := strings.Split(files, ",")
	ddd, _ := base64.StdEncoding.DecodeString(file2list[1])
	guid := uuid.New().String()
	fileName := guid + ".jpg"
	err := utils.IsNotExistMkDir(path)
	if err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败")
	}
	base64File := path + fileName
	_ = ioutil.WriteFile(base64File, ddd, 0666)
	typeStr := strings.Replace(strings.Replace(file2list[0], "data:", "", -1), ";base64", "", -1)
	fileResponse = FileResponse{
		Size:     pkg.GetFileSize(base64File),
		Path:     base64File,
		FullPath: urlPerfix + base64File,
		Name:     "",
		Type:     typeStr,
	}
	source, _ := c.GetPostForm("source")
	err = thirdUpload(source, fileName, base64File)
	if err != nil {
		e.Error(200, errors.New(""), "上传第三方失败")
		return fileResponse
	}
	if source != "1" {
		fileResponse.Path = "/static/uploadfile/" + fileName
		fileResponse.FullPath = "/static/uploadfile/" + fileName
	}
	return fileResponse
}

func (e File) multipleFile(c *gin.Context, urlPerfix string) []FileResponse {
	files := c.Request.MultipartForm.File["file"]
	source, _ := c.GetPostForm("source")
	var multipartFile []FileResponse
	for _, f := range files {
		guid := uuid.New().String()
		fileName := guid + utils.GetExt(f.Filename)

		err := utils.IsNotExistMkDir(path)
		if err != nil {
			e.Error(500, errors.New(""), "初始化文件路径失败")
		}
		multipartFileName := path + fileName
		err1 := c.SaveUploadedFile(f, multipartFileName)
		fileType, _ := utils.GetType(multipartFileName)
		if err1 == nil {
			err := thirdUpload(source, fileName, multipartFileName)
			if err != nil {
				e.Error(500, errors.New(""), "上传第三方失败")
			} else {
				fileResponse := FileResponse{
					Size:     pkg.GetFileSize(multipartFileName),
					Path:     multipartFileName,
					FullPath: urlPerfix + multipartFileName,
					Name:     f.Filename,
					Type:     fileType,
				}
				if source != "1" {
					fileResponse.Path = "/static/uploadfile/" + fileName
					fileResponse.FullPath = "/static/uploadfile/" + fileName
				}
				multipartFile = append(multipartFile, fileResponse)
			}
		}
	}
	return multipartFile
}

func (e File) singleFile(c *gin.Context, fileResponse FileResponse, urlPerfix string) (FileResponse, bool) {
	files, err := c.FormFile("file")

	if err != nil {
		e.Error(200, errors.New(""), "图片不能为空")
		return FileResponse{}, true
	}
	// 上传文件至指定目录
	guid := uuid.New().String()

	fileName := guid + utils.GetExt(files.Filename)

	err = utils.IsNotExistMkDir(path)
	if err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败")
	}
	singleFile := path + fileName
	_ = c.SaveUploadedFile(files, singleFile)
	fileType, _ := utils.GetType(singleFile)
	fileResponse = FileResponse{
		Size:     pkg.GetFileSize(singleFile),
		Path:     singleFile,
		FullPath: urlPerfix + singleFile,
		Name:     files.Filename,
		Type:     fileType,
	}
	//source, _ := c.GetPostForm("source")
	//err = thirdUpload(source, fileName, singleFile)
	//if err != nil {
	//	e.Error(200, errors.New(""), "上传第三方失败")
	//	return FileResponse{}, true
	//}
	//fileResponse.Path = "/static/uploadfile/" + fileName
	//fileResponse.FullPath = "/static/uploadfile/" + fileName
	return fileResponse, false
}

func (e File) uploadPDF(c *gin.Context, fileResponse FileResponse, urlPrefix string) (FileResponse, bool) {
	// 获取上传的文件
	files, err := c.FormFile("file")
	if err != nil {
		e.Error(200, errors.New(""), "文件不能为空")
		return FileResponse{}, true
	}

	// 验证文件类型
	if !strings.HasSuffix(strings.ToLower(files.Filename), ".pdf") {
		e.Error(200, errors.New(""), "只支持PDF文件上传")
		return FileResponse{}, true
	}

	// 生成唯一文件名
	guid := uuid.New().String()
	fileName := guid + ".pdf"

	// 确保目录存在
	err = utils.IsNotExistMkDir(path)
	if err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败")
		return FileResponse{}, true
	}

	// 保存文件
	pdfFile := path + fileName
	err = c.SaveUploadedFile(files, pdfFile)
	if err != nil {
		e.Error(500, errors.New(""), "保存PDF文件失败")
		return FileResponse{}, true
	}

	// 验证文件是否为有效的PDF
	file, err := os.Open(pdfFile)
	if err != nil {
		e.Error(500, errors.New(""), "读取PDF文件失败")
		return FileResponse{}, true
	}
	defer file.Close()

	// 读取文件头部来验证PDF格式
	buffer := make([]byte, 4)
	_, err = file.Read(buffer)
	if err != nil || !bytes.HasPrefix(buffer, []byte("%PDF")) {
		// 如果不是有效的PDF文件，删除已上传的文件
		os.Remove(pdfFile)
		e.Error(200, errors.New(""), "上传的文件不是有效的PDF格式")
		return FileResponse{}, true
	}
	err = utils.IsNotExistMkDir(path)
	if err != nil {
		e.Error(500, errors.New(""), "初始化文件路径失败")
	}
	singleFile := path + fileName

	// 构建响应
	fileResponse = FileResponse{
		Size:     pkg.GetFileSize(pdfFile),
		Path:     singleFile,
		FullPath: urlPrefix + singleFile,
		Name:     files.Filename,
		Type:     "application/pdf",
	}

	return fileResponse, false
}

func thirdUpload(source string, name string, path string) error {
	switch source {
	case "2":
		return ossUpload("img/"+name, path)
	case "3":
		return qiniuUpload("img/"+name, path)
	}
	return nil
}

func ossUpload(name string, path string) error {
	oss := file_store.ALiYunOSS{}
	return oss.UpLoad(name, path)
}

func qiniuUpload(name string, path string) error {
	oss := file_store.ALiYunOSS{}
	return oss.UpLoad(name, path)
}

// GetPDFPreview 获取 PDF 预览图
func (e File) GetPDFPreview(c *gin.Context) {
	// 1. 获取 PDF 相对路径
	pdfPath := c.Param("path")
	if pdfPath == "" {
		e.Error(400, nil, "PDF path is required")
		return
	}
	// 2. 生成或获取预览图
	h := NewPDFService("static/pdffile/", "static/pdffile/preview/")
	previewName, err := h.GetPreview(pdfPath)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("Failed to generate preview: %s", err.Error()))
		return
	}

	// 3. 返回预览图
	previewPath := filepath.Join(h.CachePath, previewName)
	c.File(previewPath)
}

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
