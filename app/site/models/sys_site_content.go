package models

import (

	"go-admin/common/models"

)

type SysSiteContent struct {
    models.Model
    
    Type string `json:"type" gorm:"type:enum('home','about','audio');comment:内容类型：首页/关于/有声读物"` 
    Title string `json:"title" gorm:"type:varchar(100);comment:内容名称"` 
    Content string `json:"content" gorm:"type:text;comment:富文本内容"` 
    models.ModelTime
    models.ControlBy
}

func (SysSiteContent) TableName() string {
    return "sys_site_content"
}

func (e *SysSiteContent) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysSiteContent) GetId() interface{} {
	return e.Id
}