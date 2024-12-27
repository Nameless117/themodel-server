package models

import (

	"go-admin/common/models"

)

type SysSiteFooterConfig struct {
    models.Model
    
    Title string `json:"title" gorm:"type:varchar(100);comment:联系标题"` 
    Email string `json:"email" gorm:"type:varchar(100);comment:联系邮箱"` 
    SubTitle string `json:"subTitle" gorm:"type:varchar(100);comment:副标题"` 
    Text string `json:"text" gorm:"type:text;comment:页脚文本内容"` 
    models.ModelTime
    models.ControlBy
}

func (SysSiteFooterConfig) TableName() string {
    return "sys_site_footer_config"
}

func (e *SysSiteFooterConfig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysSiteFooterConfig) GetId() interface{} {
	return e.Id
}