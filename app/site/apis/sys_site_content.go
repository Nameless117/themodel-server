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

type SysSiteContent struct {
	api.Api
}

// GetPage 获取站点内容管理列表
// @Summary 获取站点内容管理列表
// @Description 获取站点内容管理列表
// @Tags 站点内容管理
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysSiteContent}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-site-content [get]
// @Security Bearer
func (e SysSiteContent) GetPage(c *gin.Context) {
	req := dto.SysSiteContentGetPageReq{}
	s := service.SysSiteContent{}
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
	list := make([]models.SysSiteContent, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取站点内容管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取站点内容管理
// @Summary 获取站点内容管理
// @Description 获取站点内容管理
// @Tags 站点内容管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysSiteContent} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-site-content/{id} [get]
// @Security Bearer
func (e SysSiteContent) Get(c *gin.Context) {
	req := dto.SysSiteContentGetReq{}
	s := service.SysSiteContent{}
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
	var object models.SysSiteContent

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取站点内容管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建站点内容管理
// @Summary 创建站点内容管理
// @Description 创建站点内容管理
// @Tags 站点内容管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysSiteContentInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-site-content [post]
// @Security Bearer
func (e SysSiteContent) Insert(c *gin.Context) {
	req := dto.SysSiteContentInsertReq{}
	s := service.SysSiteContent{}
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
		e.Error(500, err, fmt.Sprintf("创建站点内容管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改站点内容管理
// @Summary 修改站点内容管理
// @Description 修改站点内容管理
// @Tags 站点内容管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysSiteContentUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-site-content/{id} [put]
// @Security Bearer
func (e SysSiteContent) Update(c *gin.Context) {
	req := dto.SysSiteContentUpdateReq{}
	s := service.SysSiteContent{}
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
		e.Error(500, err, fmt.Sprintf("修改站点内容管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除站点内容管理
// @Summary 删除站点内容管理
// @Description 删除站点内容管理
// @Tags 站点内容管理
// @Param data body dto.SysSiteContentDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-site-content [delete]
// @Security Bearer
func (e SysSiteContent) Delete(c *gin.Context) {
	s := service.SysSiteContent{}
	req := dto.SysSiteContentDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除站点内容管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
