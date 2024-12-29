package models

import (
	"go-admin/common/models"
)

type SysSiteConfig struct {
	models.Model

	LogoLarge     string `json:"logoLarge" gorm:"type:varchar(255);comment:左上角大logo URL"`
	LogoSmall     string `json:"logoSmall" gorm:"type:varchar(255);comment:右上角小logo URL"`
	LogoSmallAddr string `json:"logoSmallAddr" gorm:"type:varchar(255);comment:右上角小logo 跳转地址URL"`
	BannerBgImg   string `json:"bannerBgImg" gorm:"type:varchar(255);comment:背景图片 URL"`
	BannerBook    string `json:"bannerBook" gorm:"type:varchar(255);comment:书封面图片 URL"`
	models.ModelTime
	models.ControlBy
}

func (SysSiteConfig) TableName() string {
	return "sys_site_config"
}

func (e *SysSiteConfig) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysSiteConfig) GetId() interface{} {
	return e.Id
}
