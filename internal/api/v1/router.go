package v1

import (
	"shopping/internal/controllers"
	"shopping/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterApiRouter(r *gin.Engine, engine *gorm.DB) {
	user := r.Group("/")
	// r.POST("login/", base.Login)
	// r.POST("register/", base.Register)
	// r.GET("products/", base.GetProducts)
	// // 新增公司详情展示路由
	// r.GET("company/", base.GetCompanyInfo)

	user.POST("login/", controllers.NewUserController(services.NewUserService(engine)).Login)
	user.POST("register/", controllers.NewUserController(services.NewUserService(engine)).Register)
	

}
