package models

import (

	"go-admin/common/models"

)

type SysSiteSubscriptionConfig struct {
    models.Model
    
    Title string `json:"title" gorm:"type:varchar(100);comment:订阅标题"` 
    ExampleDesc string `json:"exampleDesc" gorm:"type:varchar(100);comment:样章描述"` 
    ExampleUrl string `json:"exampleUrl" gorm:"type:varchar(255);comment:样章下载链接"` 
    models.ModelTime
    models.ControlBy
}

func (SysSiteSubscriptionConfig) TableName() string {
    return "sys_site_subscription_config"
}

func (e *SysSiteSubscriptionConfig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysSiteSubscriptionConfig) GetId() interface{} {
	return e.Id
}