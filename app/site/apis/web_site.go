package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/site/service"
	"go-admin/app/site/service/dto"
	"path/filepath"
)

type WebSite struct {
	api.Api
}

func (w WebSite) GetContent(c *gin.Context) {
	req := dto.SysApiGetReq{}
	s := service.WebSite{}
	err := w.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		w.Logger.Error(err)
		w.Error(500, err, err.Error())
		return
	}
	//
	object, err := s.GetContent(req.Type)
	if err != nil {
		w.Error(500, err, fmt.Sprintf("获取数据，\r\n失败信息 %s", err.Error()))
		return
	}

	w.OK(object, "查询成功")
}
func (w WebSite) GetInfo(c *gin.Context) {
	s := service.WebSite{}
	err := w.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		w.Logger.Error(err)
		w.Error(500, err, err.Error())
		return
	}
	//
	object, err := s.Get()
	if err != nil {
		w.Error(500, err, fmt.Sprintf("获取数据，\r\n失败信息 %s", err.Error()))
		return
	}

	w.OK(object, "查询成功")
}

// GetPDFPreview 获取 PDF 预览图
func (w WebSite) GetPDFPreview(c *gin.Context) {
	// 1. 获取 PDF 相对路径
	pdfPath := c.Param("path")
	if pdfPath == "" {
		w.Error(400, nil, "PDF path is required")
		return
	}
	// 2. 生成或获取预览图
	h := NewPDFService("static/pdffile/", "static/pdffile/preview/")
	previewName, err := h.GetPreview(pdfPath)
	if err != nil {
		w.Error(500, err, fmt.Sprintf("Failed to generate preview: %s", err.Error()))
		return
	}

	// 3. 返回预览图
	previewPath := filepath.Join(h.CachePath, previewName)
	c.File(previewPath)
}
