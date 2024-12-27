package models

import (

	"go-admin/common/models"

)

type SysSiteMenuItems struct {
    models.Model
    
    Title string `json:"title" gorm:"type:varchar(50);comment:菜单标题"` 
    Path string `json:"path" gorm:"type:varchar(100);comment:菜单路径"` 
    SortOrder string `json:"sortOrder" gorm:"type:int;comment:排序顺序"` 
    models.ModelTime
    models.ControlBy
}

func (SysSiteMenuItems) TableName() string {
    return "sys_site_menu_items"
}

func (e *SysSiteMenuItems) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysSiteMenuItems) GetId() interface{} {
	return e.Id
}