package v1

import (
	"shopping/internal/controllers"
	"shopping/internal/middleware"
	"shopping/internal/services"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterApiRouter(r *gin.Engine, engine *gorm.DB) {
	r.Use(middleware.RateLimitMiddleware(time.Second, 100, 100))
	base := r.Group("/v1/admins")

	base.POST("/login", controllers.NewUserController(services.NewUserService(engine)).Login)
	base.POST("/register", controllers.NewUserController(services.NewUserService(engine)).Register)

	r.Use(middleware.JWTAuthMiddleware()) // 验证token ,从这里往下都需要验证token

	basicInformation := r.Group("/v1/basicInformation")
	basicInformationController := controllers.NewBaseInfoController(services.NewBaseInfoService(engine))
	basicInformation.GET("/getAllBasicInformation", basicInformationController.GetAllBasicInformation)

	carousels := r.Group("/v1/carousels")
	carouselsController := controllers.NewCarouselsController(services.NewCarouselsService(engine))
	carousels.GET("/getAllCarousels", carouselsController.GetAllCarousels)

	products := r.Group("/v1/products")
	productController := controllers.NewProductController(services.NewProductsService(engine))
	products.GET("/getAllProducts", productController.GetAllProducts)
	products.GET("/getProductById/:id", productController.GetProductById)

	certificates := r.Group("/v1/certificates")
	certificatesController := controllers.NewCertificatesController(services.NewCertificatesService(engine))
	certificates.GET("/page", certificatesController.GetCertificatesByPage)
}
