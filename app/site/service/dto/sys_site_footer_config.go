package dto

import (

	"go-admin/app/site/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysSiteFooterConfigGetPageReq struct {
	dto.Pagination     `search:"-"`
    SysSiteFooterConfigOrder
}

type SysSiteFooterConfigOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:sys_site_footer_config"`
    Title string `form:"titleOrder"  search:"type:order;column:title;table:sys_site_footer_config"`
    Email string `form:"emailOrder"  search:"type:order;column:email;table:sys_site_footer_config"`
    SubTitle string `form:"subTitleOrder"  search:"type:order;column:sub_title;table:sys_site_footer_config"`
    Text string `form:"textOrder"  search:"type:order;column:text;table:sys_site_footer_config"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_site_footer_config"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_site_footer_config"`
    
}

func (m *SysSiteFooterConfigGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysSiteFooterConfigInsertReq struct {
    Id int `json:"-" comment:""` // 
    Title string `json:"title" comment:"联系标题"`
    Email string `json:"email" comment:"联系邮箱"`
    SubTitle string `json:"subTitle" comment:"副标题"`
    Text string `json:"text" comment:"页脚文本内容"`
    common.ControlBy
}

func (s *SysSiteFooterConfigInsertReq) Generate(model *models.SysSiteFooterConfig)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Title = s.Title
    model.Email = s.Email
    model.SubTitle = s.SubTitle
    model.Text = s.Text
}

func (s *SysSiteFooterConfigInsertReq) GetId() interface{} {
	return s.Id
}

type SysSiteFooterConfigUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Title string `json:"title" comment:"联系标题"`
    Email string `json:"email" comment:"联系邮箱"`
    SubTitle string `json:"subTitle" comment:"副标题"`
    Text string `json:"text" comment:"页脚文本内容"`
    common.ControlBy
}

func (s *SysSiteFooterConfigUpdateReq) Generate(model *models.SysSiteFooterConfig)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Title = s.Title
    model.Email = s.Email
    model.SubTitle = s.SubTitle
    model.Text = s.Text
}

func (s *SysSiteFooterConfigUpdateReq) GetId() interface{} {
	return s.Id
}

// SysSiteFooterConfigGetReq 功能获取请求参数
type SysSiteFooterConfigGetReq struct {
     Id int `uri:"id"`
}
func (s *SysSiteFooterConfigGetReq) GetId() interface{} {
	return s.Id
}

// SysSiteFooterConfigDeleteReq 功能删除请求参数
type SysSiteFooterConfigDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysSiteFooterConfigDeleteReq) GetId() interface{} {
	return s.Ids
}
