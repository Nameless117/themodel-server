package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/site/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerWebSiteConfigRouter)
	routerCheckRole = append(routerCheckRole, registerFileRouter)
}

// registerWebSiteConfigRouter
func registerWebSiteConfigRouter(v1 *gin.RouterGroup) {
	api := apis.WebSite{}
	r := v1.Group("/site")
	{
		r.GET("/info", api.GetInfo)
		r.GET("/content", api.GetContent)
	}
}

// 需认证的路由代码
func registerFileRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.WebSite{}
	r := v1.Group("/site").Use(authMiddleware.MiddlewareFunc())
	{
		r.POST("/upload", api.Upload)
	}
}
