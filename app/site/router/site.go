package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/site/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerWebSiteConfigRouter)
	//routerCheckRole = append(routerCheckRole, registerFileRouter)
}

// registerWebSiteConfigRouter
func registerWebSiteConfigRouter(v1 *gin.RouterGroup) {
	api := apis.WebSite{}
	r := v1.Group("/site")
	{
		r.GET("/info", api.GetInfo)
		r.GET("/content/:type", api.GetContent)
		//r.GET("/pdf/preview/*path", api.GetPDFPreview)
	}
}

// 需认证的路由代码
//func registerFileRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
//	api := apis.WebSite{}
//	r := v1.Group("/site").Use(authMiddleware.MiddlewareFunc())
//	{
//		r.POST("/pdf/preview/*path", api.GetPDFPreview)
//	}
//}
