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

type SysSiteMenuItems struct {
	api.Api
}

// GetPage 获取站点菜单管理列表
// @Summary 获取站点菜单管理列表
// @Description 获取站点菜单管理列表
// @Tags 站点菜单管理
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysSiteMenuItems}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-site-menu-items [get]
// @Security Bearer
func (e SysSiteMenuItems) GetPage(c *gin.Context) {
    req := dto.SysSiteMenuItemsGetPageReq{}
    s := service.SysSiteMenuItems{}
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
	list := make([]models.SysSiteMenuItems, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取站点菜单管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取站点菜单管理
// @Summary 获取站点菜单管理
// @Description 获取站点菜单管理
// @Tags 站点菜单管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysSiteMenuItems} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-site-menu-items/{id} [get]
// @Security Bearer
func (e SysSiteMenuItems) Get(c *gin.Context) {
	req := dto.SysSiteMenuItemsGetReq{}
	s := service.SysSiteMenuItems{}
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
	var object models.SysSiteMenuItems

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取站点菜单管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建站点菜单管理
// @Summary 创建站点菜单管理
// @Description 创建站点菜单管理
// @Tags 站点菜单管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysSiteMenuItemsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-site-menu-items [post]
// @Security Bearer
func (e SysSiteMenuItems) Insert(c *gin.Context) {
    req := dto.SysSiteMenuItemsInsertReq{}
    s := service.SysSiteMenuItems{}
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
		e.Error(500, err, fmt.Sprintf("创建站点菜单管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改站点菜单管理
// @Summary 修改站点菜单管理
// @Description 修改站点菜单管理
// @Tags 站点菜单管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysSiteMenuItemsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-site-menu-items/{id} [put]
// @Security Bearer
func (e SysSiteMenuItems) Update(c *gin.Context) {
    req := dto.SysSiteMenuItemsUpdateReq{}
    s := service.SysSiteMenuItems{}
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
		e.Error(500, err, fmt.Sprintf("修改站点菜单管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除站点菜单管理
// @Summary 删除站点菜单管理
// @Description 删除站点菜单管理
// @Tags 站点菜单管理
// @Param data body dto.SysSiteMenuItemsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-site-menu-items [delete]
// @Security Bearer
func (e SysSiteMenuItems) Delete(c *gin.Context) {
    s := service.SysSiteMenuItems{}
    req := dto.SysSiteMenuItemsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除站点菜单管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
