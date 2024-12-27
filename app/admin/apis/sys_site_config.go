package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type SysSiteConfig struct {
	api.Api
}

// GetPage 获取站点配置管理列表
// @Summary 获取站点配置管理列表
// @Description 获取站点配置管理列表
// @Tags 站点配置管理
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysSiteConfig}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-site-config [get]
// @Security Bearer
func (e SysSiteConfig) GetPage(c *gin.Context) {
    req := dto.SysSiteConfigGetPageReq{}
    s := service.SysSiteConfig{}
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
	list := make([]models.SysSiteConfig, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取站点配置管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取站点配置管理
// @Summary 获取站点配置管理
// @Description 获取站点配置管理
// @Tags 站点配置管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysSiteConfig} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-site-config/{id} [get]
// @Security Bearer
func (e SysSiteConfig) Get(c *gin.Context) {
	req := dto.SysSiteConfigGetReq{}
	s := service.SysSiteConfig{}
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
	var object models.SysSiteConfig

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取站点配置管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建站点配置管理
// @Summary 创建站点配置管理
// @Description 创建站点配置管理
// @Tags 站点配置管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysSiteConfigInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-site-config [post]
// @Security Bearer
func (e SysSiteConfig) Insert(c *gin.Context) {
    req := dto.SysSiteConfigInsertReq{}
    s := service.SysSiteConfig{}
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
		e.Error(500, err, fmt.Sprintf("创建站点配置管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改站点配置管理
// @Summary 修改站点配置管理
// @Description 修改站点配置管理
// @Tags 站点配置管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysSiteConfigUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-site-config/{id} [put]
// @Security Bearer
func (e SysSiteConfig) Update(c *gin.Context) {
    req := dto.SysSiteConfigUpdateReq{}
    s := service.SysSiteConfig{}
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
		e.Error(500, err, fmt.Sprintf("修改站点配置管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除站点配置管理
// @Summary 删除站点配置管理
// @Description 删除站点配置管理
// @Tags 站点配置管理
// @Param data body dto.SysSiteConfigDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-site-config [delete]
// @Security Bearer
func (e SysSiteConfig) Delete(c *gin.Context) {
    s := service.SysSiteConfig{}
    req := dto.SysSiteConfigDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除站点配置管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
