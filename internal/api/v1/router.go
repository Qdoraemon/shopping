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
	userController := controllers.NewUserController(services.NewUserService(engine))
	basicInformationController := controllers.NewBaseInfoController(services.NewBaseInfoService(engine))
	carouselsController := controllers.NewCarouselsController(services.NewCarouselsService(engine))
	productController := controllers.NewProductController(services.NewProductsService(engine))
	// certificatesController := controllers.NewCertificatesController(services.NewCertificatesService(engine))
	BrandController := controllers.NewBrandController(services.NewBrandService(engine))
	categoryController := controllers.NewCategoryController(services.NewCategoryService(engine))

	// TODO 文件上传 需要有一个专门的接口来处理文件上传
	fileUpload := r.Group("/v1/fileUpload")
	fileUpload.POST("/upload", baseController.UploadImage)

	base.GET("/getImage", baseController.GetImage)
	base.DELETE("/deleteImage", baseController.DeleteImage)

	base.POST("/admins/login", userController.Login)
	base.POST("/admins/register", userController.Register)
	notAuthcarousels := r.Group("/v1/carousels")
	notAuthcarousels.GET("/getAllCarousels", carouselsController.GetAllCarousels)

	basicInformation := r.Group("/v1/basicInformation")
	basicInformation.GET("/getAllBasicInformation", basicInformationController.GetAllBasicInformation)

	products := r.Group("/v1/products")

	// 獲取所有商品
	products.GET("/getAllProducts", productController.GetAllProducts)
	// 根據ID獲取商品
	products.GET("/getProductById/:id", productController.GetProductById)
	// 根據分類ID獲取商品
	products.GET("/getProductsByCategoryID/:categoryID", productController.GetProductsByCategoryID)
	// 根據品牌ID獲取商品
	products.GET("/getProductsByBrandID/:brandID", productController.GetProductsByBrandID)
	// 根據名稱獲取商品
	products.GET("/getProductsByName/:name", productController.GetProductsByName)

	// 定義categories分組
	categories := r.Group("/v1/categories")
	// 獲取所有商品分類
	categories.GET("/getAllCategories", categoryController.GetAllCategories)

	// 上傳logo
	logos := r.Group("/v1/logos")
	logos.POST("/uploadLogo", baseController.UploadLogo)

	// 定義Brand分組
	brands := r.Group("/v1/brands")
	// 獲取所有brand
	brands.GET("/getAllBrands", BrandController.GetAllBrand)

	r.Use(middleware.JWTAuthMiddleware()) // 验证token ,从这里往下都需要验证token

	// brands操作
	brands.POST("/addBrand", BrandController.AddBrand)
	brands.PUT("/updateBrand", BrandController.UpdateBrand)
	brands.DELETE("/deleteBrand/:id", BrandController.DeleteBrand)
	brands.POST("/getBrandsByPage", BrandController.GetBrandByPage)

	// category操作
	categories.POST("/addCategory", categoryController.AddCategory)
	categories.PUT("/updateCategory", categoryController.UpdateCategory)
	categories.DELETE("/deleteCategory/:id", categoryController.DeleteCategory)
	categories.POST("/getCategoriesByPage", categoryController.GetCategoriesByPage)

	// 獲取用戶信息
	base.GET("/client/getRoles", userController.GetInfo)

	// 添加商品
	products.POST("/addProduct", productController.AddProduct)

	// 刪除商品
	products.DELETE("/deleteProduct/:id", productController.DeleteProduct)

	// 批量刪除商品
	products.DELETE("/deleteProducts", productController.DeleteProducts)

	// 複製商品
	products.POST("/copyProduct/:id", productController.CopyProduct)

	// 更新商品
	products.PUT("/updateProduct", productController.UpdateProduct)
}
