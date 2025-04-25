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
	base := r.Group("/v1")
	baseController := controllers.NewBaseController()
	// TODO 文件上传 需要有一个专门的接口来处理文件上传
	fileUpload := r.Group("/v1/fileUpload")
	fileUpload.POST("/upload", baseController.UploadImage)

	base.GET("/getLatestImage", baseController.GetImage)

	base.POST("/admins/login", controllers.NewUserController(services.NewUserService(engine)).Login)
	base.POST("/admins/register", controllers.NewUserController(services.NewUserService(engine)).Register)

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
	// 注册添加证书路由
	certificates.POST("/addCertificate", certificatesController.AddCertificate)
	certificates.PUT("/updateCertificate", certificatesController.UpdateCertificate)
	// 注册删除证书路由
	certificates.DELETE("/deleteCertificate/:id", certificatesController.DeleteCertificate)

	brands := r.Group("/v1/brands")
	brandsController := controllers.NewBrandsController(services.NewBrandsService(engine))
	brands.GET("/", brandsController.GetBrandsByPage)
	// 注册添加证书路由
	brands.POST("/addBrand", brandsController.AddBrands)
	brands.PUT("/updateBrand", brandsController.UpdateBrands)
	// 注册删除证书路由
	brands.DELETE("/deleteBrand/:id", brandsController.DeleteBrands)

}
