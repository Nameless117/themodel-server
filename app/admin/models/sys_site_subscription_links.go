package models

import (

	"go-admin/common/models"

)

type SysSiteSubscriptionLinks struct {
    models.Model
    
    Title string `json:"title" gorm:"type:varchar(50);comment:订阅平台名称"` 
    Url string `json:"url" gorm:"type:varchar(255);comment:订阅链接"` 
    SortOrder int64 `json:"sortOrder" gorm:"type:int;comment:排序顺序"` 
    models.ModelTime
    models.ControlBy
}

func (SysSiteSubscriptionLinks) TableName() string {
    return "sys_site_subscription_links"
}

func (e *SysSiteSubscriptionLinks) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysSiteSubscriptionLinks) GetId() interface{} {
	return e.Id
}