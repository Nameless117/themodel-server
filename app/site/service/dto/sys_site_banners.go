package dto

import (

	"go-admin/app/site/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysSiteBannersGetPageReq struct {
	dto.Pagination     `search:"-"`
    SysSiteBannersOrder
}

type SysSiteBannersOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:sys_site_banners"`
    Type string `form:"typeOrder"  search:"type:order;column:type;table:sys_site_banners"`
    Description string `form:"descriptionOrder"  search:"type:order;column:description;table:sys_site_banners"`
    SubDescription string `form:"subDescriptionOrder"  search:"type:order;column:sub_description;table:sys_site_banners"`
    PdfTitle string `form:"pdfTitleOrder"  search:"type:order;column:pdf_title;table:sys_site_banners"`
    PdfUrl string `form:"pdfUrlOrder"  search:"type:order;column:pdf_url;table:sys_site_banners"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_site_banners"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_site_banners"`
    
}

func (m *SysSiteBannersGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysSiteBannersInsertReq struct {
    Id int `json:"-" comment:""` // 
    Type string `json:"type" comment:"banner类型：首页/有声读物"`
    Description string `json:"description" comment:"banner描述"`
    SubDescription string `json:"subDescription" comment:"二级描述（用于有声读物）"`
    PdfTitle string `json:"pdfTitle" comment:"PDF标题"`
    PdfUrl string `json:"pdfUrl" comment:"PDF下载链接"`
    common.ControlBy
}

func (s *SysSiteBannersInsertReq) Generate(model *models.SysSiteBanners)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Type = s.Type
    model.Description = s.Description
    model.SubDescription = s.SubDescription
    model.PdfTitle = s.PdfTitle
    model.PdfUrl = s.PdfUrl
}

func (s *SysSiteBannersInsertReq) GetId() interface{} {
	return s.Id
}

type SysSiteBannersUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Type string `json:"type" comment:"banner类型：首页/有声读物"`
    Description string `json:"description" comment:"banner描述"`
    SubDescription string `json:"subDescription" comment:"二级描述（用于有声读物）"`
    PdfTitle string `json:"pdfTitle" comment:"PDF标题"`
    PdfUrl string `json:"pdfUrl" comment:"PDF下载链接"`
    common.ControlBy
}

func (s *SysSiteBannersUpdateReq) Generate(model *models.SysSiteBanners)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Type = s.Type
    model.Description = s.Description
    model.SubDescription = s.SubDescription
    model.PdfTitle = s.PdfTitle
    model.PdfUrl = s.PdfUrl
}

func (s *SysSiteBannersUpdateReq) GetId() interface{} {
	return s.Id
}

// SysSiteBannersGetReq 功能获取请求参数
type SysSiteBannersGetReq struct {
     Id int `uri:"id"`
}
func (s *SysSiteBannersGetReq) GetId() interface{} {
	return s.Id
}

// SysSiteBannersDeleteReq 功能删除请求参数
type SysSiteBannersDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysSiteBannersDeleteReq) GetId() interface{} {
	return s.Ids
}
