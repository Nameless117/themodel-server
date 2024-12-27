package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysSiteMenuItemsGetPageReq struct {
	dto.Pagination     `search:"-"`
    SysSiteMenuItemsOrder
}

type SysSiteMenuItemsOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:sys_site_menu_items"`
    Title string `form:"titleOrder"  search:"type:order;column:title;table:sys_site_menu_items"`
    Path string `form:"pathOrder"  search:"type:order;column:path;table:sys_site_menu_items"`
    SortOrder string `form:"sortOrderOrder"  search:"type:order;column:sort_order;table:sys_site_menu_items"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_site_menu_items"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_site_menu_items"`
    
}

func (m *SysSiteMenuItemsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysSiteMenuItemsInsertReq struct {
    Id int `json:"-" comment:""` // 
    Title string `json:"title" comment:"菜单标题"`
    Path string `json:"path" comment:"菜单路径"`
    SortOrder string `json:"sortOrder" comment:"排序顺序"`
    common.ControlBy
}

func (s *SysSiteMenuItemsInsertReq) Generate(model *models.SysSiteMenuItems)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Title = s.Title
    model.Path = s.Path
    model.SortOrder = s.SortOrder
}

func (s *SysSiteMenuItemsInsertReq) GetId() interface{} {
	return s.Id
}

type SysSiteMenuItemsUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Title string `json:"title" comment:"菜单标题"`
    Path string `json:"path" comment:"菜单路径"`
    SortOrder string `json:"sortOrder" comment:"排序顺序"`
    common.ControlBy
}

func (s *SysSiteMenuItemsUpdateReq) Generate(model *models.SysSiteMenuItems)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Title = s.Title
    model.Path = s.Path
    model.SortOrder = s.SortOrder
}

func (s *SysSiteMenuItemsUpdateReq) GetId() interface{} {
	return s.Id
}

// SysSiteMenuItemsGetReq 功能获取请求参数
type SysSiteMenuItemsGetReq struct {
     Id int `uri:"id"`
}
func (s *SysSiteMenuItemsGetReq) GetId() interface{} {
	return s.Id
}

// SysSiteMenuItemsDeleteReq 功能删除请求参数
type SysSiteMenuItemsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysSiteMenuItemsDeleteReq) GetId() interface{} {
	return s.Ids
}
