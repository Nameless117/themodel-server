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

type SysSiteSubscriptionConfig struct {
	api.Api
}

// GetPage 获取站点订阅配置列表
// @Summary 获取站点订阅配置列表
// @Description 获取站点订阅配置列表
// @Tags 站点订阅配置
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysSiteSubscriptionConfig}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-site-subscription-config [get]
// @Security Bearer
func (e SysSiteSubscriptionConfig) GetPage(c *gin.Context) {
    req := dto.SysSiteSubscriptionConfigGetPageReq{}
    s := service.SysSiteSubscriptionConfig{}
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
	list := make([]models.SysSiteSubscriptionConfig, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取站点订阅配置失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取站点订阅配置
// @Summary 获取站点订阅配置
// @Description 获取站点订阅配置
// @Tags 站点订阅配置
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SysSiteSubscriptionConfig} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-site-subscription-config/{id} [get]
// @Security Bearer
func (e SysSiteSubscriptionConfig) Get(c *gin.Context) {
	req := dto.SysSiteSubscriptionConfigGetReq{}
	s := service.SysSiteSubscriptionConfig{}
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
	var object models.SysSiteSubscriptionConfig

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取站点订阅配置失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建站点订阅配置
// @Summary 创建站点订阅配置
// @Description 创建站点订阅配置
// @Tags 站点订阅配置
// @Accept application/json
// @Product application/json
// @Param data body dto.SysSiteSubscriptionConfigInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sys-site-subscription-config [post]
// @Security Bearer
func (e SysSiteSubscriptionConfig) Insert(c *gin.Context) {
    req := dto.SysSiteSubscriptionConfigInsertReq{}
    s := service.SysSiteSubscriptionConfig{}
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
		e.Error(500, err, fmt.Sprintf("创建站点订阅配置失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改站点订阅配置
// @Summary 修改站点订阅配置
// @Description 修改站点订阅配置
// @Tags 站点订阅配置
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysSiteSubscriptionConfigUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sys-site-subscription-config/{id} [put]
// @Security Bearer
func (e SysSiteSubscriptionConfig) Update(c *gin.Context) {
    req := dto.SysSiteSubscriptionConfigUpdateReq{}
    s := service.SysSiteSubscriptionConfig{}
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
		e.Error(500, err, fmt.Sprintf("修改站点订阅配置失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除站点订阅配置
// @Summary 删除站点订阅配置
// @Description 删除站点订阅配置
// @Tags 站点订阅配置
// @Param data body dto.SysSiteSubscriptionConfigDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sys-site-subscription-config [delete]
// @Security Bearer
func (e SysSiteSubscriptionConfig) Delete(c *gin.Context) {
    s := service.SysSiteSubscriptionConfig{}
    req := dto.SysSiteSubscriptionConfigDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除站点订阅配置失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
