package v1

import (
	"shopping/internal/controllers"
	"shopping/internal/middleware"
	"shopping/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterApiRouter(r *gin.Engine, engine *gorm.DB) {
	base := r.Group("/v1/admins")

	base.POST("/login", controllers.NewUserController(services.NewUserService(engine)).Login)
	base.POST("/register", controllers.NewUserController(services.NewUserService(engine)).Register)

	basicInformation := r.Group("/v1/basicInformation")
	basicInformation.Use(middleware.JWTAuthMiddleware())
	basicInformation.GET("/getAllBasicInformation", controllers.NewBaseInfoController(services.NewBaseInfoService(engine)).GetAllBasicInformation)

}
