package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/site/service"
	"go-admin/app/site/service/dto"
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

func (w WebSite) Upload(c *gin.Context) {

	// 获取上传文件
	file, err := c.FormFile("file")
	if err != nil {
		w.Error(400, err, "获取上传文件失败")
		return
	}
	// 创建上传器实例，设置不同文件类型的限制
	uploader := NewLocalUploader(UploadConfig{
		SavePath: "./uploads",
		FileTypeLimits: map[string]FileTypeLimit{
			"image": {
				Extensions: []string{".jpg", ".jpeg", ".png", ".gif"},
				MaxSize:    10 * 1024 * 1024, // 10MB
			},
			"pdf": {
				Extensions: []string{".pdf"},
				MaxSize:    50 * 1024 * 1024, // 50MB
			},
		},
	})

	// 上传文件
	response, err := uploader.Upload(file)
	if err != nil {
		w.Error(500, err, fmt.Sprintf("上传文件，\r\n失败信息 %s", err.Error()))
		return
	}
	// 返回结果
	w.OK(response, "上传成功")
}
