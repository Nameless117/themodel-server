package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/site/models"
	"go-admin/app/site/service"
	"go-admin/app/site/service/dto"
	"go-admin/common/actions"
)

type SysSiteBanners struct {
	api.Api
}

// GetPage 获取站点Banner管理列表
// @Summary 获取站点Banner管理列表
// @Description 获取站点Banner管理列表
// @Tags 站点Banner管理
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysSiteBanners}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-site-banners [get]
// @Security Bearer
func (e SysSiteBanners) GetPage(c *gin.Context) {
    req := dto.SysSiteBannersGetPageReq{}
    s := service.SysSiteBanners{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.SysSiteBanners, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取站点Banner管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取站点Banner管理
// @Summary 获取站点Banner管理
// @Description 获取站点Banner管理
// @Tags 站点Banner管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysSiteBanners} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-site-banners/{id} [get]
// @Security Bearer
func (e SysSiteBanners) Get(c *gin.Context) {
	req := dto.SysSiteBannersGetReq{}
	s := service.SysSiteBanners{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysSiteBanners

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取站点Banner管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建站点Banner管理
// @Summary 创建站点Banner管理
// @Description 创建站点Banner管理
// @Tags 站点Banner管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysSiteBannersInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-site-banners [post]
// @Security Bearer
func (e SysSiteBanners) Insert(c *gin.Context) {
    req := dto.SysSiteBannersInsertReq{}
    s := service.SysSiteBanners{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建站点Banner管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改站点Banner管理
// @Summary 修改站点Banner管理
// @Description 修改站点Banner管理
// @Tags 站点Banner管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysSiteBannersUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-site-banners/{id} [put]
// @Security Bearer
func (e SysSiteBanners) Update(c *gin.Context) {
    req := dto.SysSiteBannersUpdateReq{}
    s := service.SysSiteBanners{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改站点Banner管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除站点Banner管理
// @Summary 删除站点Banner管理
// @Description 删除站点Banner管理
// @Tags 站点Banner管理
// @Param data body dto.SysSiteBannersDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-site-banners [delete]
// @Security Bearer
func (e SysSiteBanners) Delete(c *gin.Context) {
    s := service.SysSiteBanners{}
    req := dto.SysSiteBannersDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除站点Banner管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
