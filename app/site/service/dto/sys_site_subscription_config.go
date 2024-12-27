package dto

import (

	"go-admin/app/site/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysSiteSubscriptionConfigGetPageReq struct {
	dto.Pagination     `search:"-"`
    SysSiteSubscriptionConfigOrder
}

type SysSiteSubscriptionConfigOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:sys_site_subscription_config"`
    Title string `form:"titleOrder"  search:"type:order;column:title;table:sys_site_subscription_config"`
    ExampleDesc string `form:"exampleDescOrder"  search:"type:order;column:example_desc;table:sys_site_subscription_config"`
    ExampleUrl string `form:"exampleUrlOrder"  search:"type:order;column:example_url;table:sys_site_subscription_config"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_site_subscription_config"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_site_subscription_config"`
    
}

func (m *SysSiteSubscriptionConfigGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysSiteSubscriptionConfigInsertReq struct {
    Id int `json:"-" comment:""` // 
    Title string `json:"title" comment:"订阅标题"`
    ExampleDesc string `json:"exampleDesc" comment:"样章描述"`
    ExampleUrl string `json:"exampleUrl" comment:"样章下载链接"`
    common.ControlBy
}

func (s *SysSiteSubscriptionConfigInsertReq) Generate(model *models.SysSiteSubscriptionConfig)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Title = s.Title
    model.ExampleDesc = s.ExampleDesc
    model.ExampleUrl = s.ExampleUrl
}

func (s *SysSiteSubscriptionConfigInsertReq) GetId() interface{} {
	return s.Id
}

type SysSiteSubscriptionConfigUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Title string `json:"title" comment:"订阅标题"`
    ExampleDesc string `json:"exampleDesc" comment:"样章描述"`
    ExampleUrl string `json:"exampleUrl" comment:"样章下载链接"`
    common.ControlBy
}

func (s *SysSiteSubscriptionConfigUpdateReq) Generate(model *models.SysSiteSubscriptionConfig)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Title = s.Title
    model.ExampleDesc = s.ExampleDesc
    model.ExampleUrl = s.ExampleUrl
}

func (s *SysSiteSubscriptionConfigUpdateReq) GetId() interface{} {
	return s.Id
}

// SysSiteSubscriptionConfigGetReq 功能获取请求参数
type SysSiteSubscriptionConfigGetReq struct {
     Id int `uri:"id"`
}
func (s *SysSiteSubscriptionConfigGetReq) GetId() interface{} {
	return s.Id
}

// SysSiteSubscriptionConfigDeleteReq 功能删除请求参数
type SysSiteSubscriptionConfigDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysSiteSubscriptionConfigDeleteReq) GetId() interface{} {
	return s.Ids
}
