package router

import (
	"shopping/app/base"

	"github.com/gin-gonic/gin"
)

func RegisterApiRouter(r *gin.Engine) {
	r.Group("/")
	r.POST("login/", base.Login)
	r.POST("register/", base.Register)
	r.GET("products/", base.GetProducts)
	// 新增公司详情展示路由
	r.GET("company/", base.GetCompanyInfo)
}
