package dto

import (

	"go-admin/app/site/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysSiteConfigGetPageReq struct {
	dto.Pagination     `search:"-"`
    SysSiteConfigOrder
}

type SysSiteConfigOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:sys_site_config"`
    LogoLarge string `form:"logoLargeOrder"  search:"type:order;column:logo_large;table:sys_site_config"`
    LogoSmall string `form:"logoSmallOrder"  search:"type:order;column:logo_small;table:sys_site_config"`
    BannerBgImg string `form:"bannerBgImgOrder"  search:"type:order;column:banner_bg_img;table:sys_site_config"`
    BannerBook string `form:"bannerBookOrder"  search:"type:order;column:banner_book;table:sys_site_config"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:sys_site_config"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:sys_site_config"`
    
}

func (m *SysSiteConfigGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysSiteConfigInsertReq struct {
    Id int `json:"-" comment:""` // 
    LogoLarge string `json:"logoLarge" comment:"左上角大logo URL"`
    LogoSmall string `json:"logoSmall" comment:"右上角小logo URL"`
    BannerBgImg string `json:"bannerBgImg" comment:"背景图片 URL"`
    BannerBook string `json:"bannerBook" comment:"书封面图片 URL"`
    common.ControlBy
}

func (s *SysSiteConfigInsertReq) Generate(model *models.SysSiteConfig)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.LogoLarge = s.LogoLarge
    model.LogoSmall = s.LogoSmall
    model.BannerBgImg = s.BannerBgImg
    model.BannerBook = s.BannerBook
}

func (s *SysSiteConfigInsertReq) GetId() interface{} {
	return s.Id
}

type SysSiteConfigUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    LogoLarge string `json:"logoLarge" comment:"左上角大logo URL"`
    LogoSmall string `json:"logoSmall" comment:"右上角小logo URL"`
    BannerBgImg string `json:"bannerBgImg" comment:"背景图片 URL"`
    BannerBook string `json:"bannerBook" comment:"书封面图片 URL"`
    common.ControlBy
}

func (s *SysSiteConfigUpdateReq) Generate(model *models.SysSiteConfig)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.LogoLarge = s.LogoLarge
    model.LogoSmall = s.LogoSmall
    model.BannerBgImg = s.BannerBgImg
    model.BannerBook = s.BannerBook
}

func (s *SysSiteConfigUpdateReq) GetId() interface{} {
	return s.Id
}

// SysSiteConfigGetReq 功能获取请求参数
type SysSiteConfigGetReq struct {
     Id int `uri:"id"`
}
func (s *SysSiteConfigGetReq) GetId() interface{} {
	return s.Id
}

// SysSiteConfigDeleteReq 功能删除请求参数
type SysSiteConfigDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SysSiteConfigDeleteReq) GetId() interface{} {
	return s.Ids
}
