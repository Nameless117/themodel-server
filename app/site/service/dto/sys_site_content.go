package dto

import (

	"go-admin/app/site/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysSiteContentGetPageReq struct {
	dto.Pagination     `search:"-"`
    SysSiteContentOrder
}

type SysSiteContentOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:sys_site_content"`
    Type string `form:"typeOrder"  search:"type:order;column:type;table:sys_site_content"`
    Title string `form:"titleOrder"  search:"type:order;column:title;table:sys_site_content"`
    Content string `form:"contentOrder"  search:"type:order;column:content;table:sys_site_content"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_site_content"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_site_content"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:sys_site_content"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:sys_site_content"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:sys_site_content"`
    
}

func (m *SysSiteContentGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysSiteContentInsertReq struct {
    Id int `json:"-" comment:""` // 
    Type string `json:"type" comment:"内容类型：首页/关于/有声读物"`
    Title string `json:"title" comment:"内容名称"`
    Content string `json:"content" comment:"富文本内容"`
    common.ControlBy
}

func (s *SysSiteContentInsertReq) Generate(model *models.SysSiteContent)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Type = s.Type
    model.Title = s.Title
    model.Content = s.Content
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SysSiteContentInsertReq) GetId() interface{} {
	return s.Id
}

type SysSiteContentUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Type string `json:"type" comment:"内容类型：首页/关于/有声读物"`
    Title string `json:"title" comment:"内容名称"`
    Content string `json:"content" comment:"富文本内容"`
    common.ControlBy
}

func (s *SysSiteContentUpdateReq) Generate(model *models.SysSiteContent)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Type = s.Type
    model.Title = s.Title
    model.Content = s.Content
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SysSiteContentUpdateReq) GetId() interface{} {
	return s.Id
}

// SysSiteContentGetReq 功能获取请求参数
type SysSiteContentGetReq struct {
     Id int `uri:"id"`
}
func (s *SysSiteContentGetReq) GetId() interface{} {
	return s.Id
}

// SysSiteContentDeleteReq 功能删除请求参数
type SysSiteContentDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysSiteContentDeleteReq) GetId() interface{} {
	return s.Ids
}
