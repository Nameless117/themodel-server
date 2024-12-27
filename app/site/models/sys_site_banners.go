package models

import (

	"go-admin/common/models"

)

type SysSiteBanners struct {
    models.Model
    
    Type string `json:"type" gorm:"type:enum('home','audio');comment:banner类型：首页/有声读物"` 
    Description string `json:"description" gorm:"type:text;comment:banner描述"` 
    SubDescription string `json:"subDescription" gorm:"type:text;comment:二级描述（用于有声读物）"` 
    PdfTitle string `json:"pdfTitle" gorm:"type:varchar(50);comment:PDF标题"` 
    PdfUrl string `json:"pdfUrl" gorm:"type:varchar(255);comment:PDF下载链接"` 
    models.ModelTime
    models.ControlBy
}

func (SysSiteBanners) TableName() string {
    return "sys_site_banners"
}

func (e *SysSiteBanners) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysSiteBanners) GetId() interface{} {
	return e.Id
}