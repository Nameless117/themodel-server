package dto

import (
	"go-admin/app/site/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysSiteSubscriptionLinksGetPageReq struct {
	dto.Pagination `search:"-"`
	SysSiteSubscriptionLinksOrder
}

type SysSiteSubscriptionLinksOrder struct {
	Id        string `form:"idOrder"  search:"type:order;column:id;table:sys_site_subscription_links"`
	Title     string `form:"titleOrder"  search:"type:order;column:title;table:sys_site_subscription_links"`
	Url       string `form:"urlOrder"  search:"type:order;column:url;table:sys_site_subscription_links"`
	SortOrder string `form:"sortOrderOrder"  search:"type:order;column:sort_order;table:sys_site_subscription_links"`
	CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_site_subscription_links"`
	UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_site_subscription_links"`
}

func (m *SysSiteSubscriptionLinksGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysSiteSubscriptionLinksInsertReq struct {
	Id        int    `json:"-" comment:""` //
	Title     string `json:"title" comment:"订阅平台名称"`
	Url       string `json:"url" comment:"订阅链接"`
	SortOrder string `json:"sortOrder" comment:"排序顺序"`
	common.ControlBy
}

func (s *SysSiteSubscriptionLinksInsertReq) Generate(model *models.SysSiteSubscriptionLinks) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Url = s.Url
	model.SortOrder = s.SortOrder
}

func (s *SysSiteSubscriptionLinksInsertReq) GetId() interface{} {
	return s.Id
}

type SysSiteSubscriptionLinksUpdateReq struct {
	Id        int    `uri:"id" comment:""` //
	Title     string `json:"title" comment:"订阅平台名称"`
	Url       string `json:"url" comment:"订阅链接"`
	SortOrder string `json:"sortOrder" comment:"排序顺序"`
	common.ControlBy
}

func (s *SysSiteSubscriptionLinksUpdateReq) Generate(model *models.SysSiteSubscriptionLinks) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Url = s.Url
	model.SortOrder = s.SortOrder
}

func (s *SysSiteSubscriptionLinksUpdateReq) GetId() interface{} {
	return s.Id
}

// SysSiteSubscriptionLinksGetReq 功能获取请求参数
type SysSiteSubscriptionLinksGetReq struct {
	Id int `uri:"id"`
}

func (s *SysSiteSubscriptionLinksGetReq) GetId() interface{} {
	return s.Id
}

// SysSiteSubscriptionLinksDeleteReq 功能删除请求参数
type SysSiteSubscriptionLinksDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysSiteSubscriptionLinksDeleteReq) GetId() interface{} {
	return s.Ids
}
