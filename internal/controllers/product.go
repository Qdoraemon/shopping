package controllers

import (
	"shopping/internal/models"
	"shopping/internal/services"
	"shopping/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *services.ProductService
}

func NewProductController(ProductService *services.ProductService) *ProductController {
	return &ProductController{ProductService: ProductService}
}

func (l *ProductController) GetAllProducts(c *gin.Context) {

	result, err := l.ProductService.GetAllProducts()
	if err != nil {
		c.JSON(200, utils.Error(400, "获取失败"))
		return
	}
	// 3. 返回结果
	c.JSON(200, utils.Success(result, "成功獲取所有商品"))

}

func (l *ProductController) GetProductById(c *gin.Context) {
	id := c.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, utils.Error(400, "id传入错误"))
		return
	}
	result, err := l.ProductService.GetProductById(idNum)
	if err != nil {
		c.JSON(200, utils.Error(400, "获取失败"))
		return
	}
	// 3. 返回结果
	c.JSON(200, utils.Success(result, "成功獲取商品"))

}

func (l *ProductController) AddProduct(c *gin.Context) {
	var product models.Product
	c.ShouldBindJSON(&product)
	// data := c.PostForm("data")
	// fmt.Println("data: ", product)
	err := l.ProductService.AddProduct(&product)
	if err != nil {
		// fmt.Println(err)
		c.JSON(200, utils.Error(400, "获取失败"))
		return
	}
	c.JSON(200, utils.Success("", "登录成功"))
}

func (l *ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	idNum, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, utils.Error(400, "id传入错误"))
		return
	}
	err = l.ProductService.DeleteProduct(idNum)
	if err != nil {
		c.JSON(200, utils.Error(400, "刪除失败"))
		return
	}
	// 3. 返回结果
	c.JSON(200, utils.Success("", "刪除成功"))
}

// DeleteProducts handles the batch deletion of products
func (l *ProductController) DeleteProducts(c *gin.Context) {

	var req struct {
		IDs []int `json:"ids"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(200, utils.Error(400, "id传入错误"))
		return
	}

	err = l.ProductService.DeleteProducts(req.IDs)
	if err != nil {
		c.JSON(200, utils.Error(400, "批量刪除失败"))
		return
	}

	// 3. 返回结果
	c.JSON(200, utils.Success("", "批量刪除成功"))

}

func (l *ProductController) CopyProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, utils.Error(400, "id传入错误"))
		return
	}

	newProduct, err := l.ProductService.CopyProduct(id)
	if err != nil {
		c.JSON(200, utils.Error(400, "复制失败"))
		return
	}

	c.JSON(200, utils.Success(newProduct, "复制成功"))
}

func (l *ProductController) UpdateProduct(c *gin.Context) {
	var product models.Product
	err := c.ShouldBindJSON(&product)

	// fmt.Println("product: ", product)
	if err != nil {
		c.JSON(200, utils.Error(400, "product格式傳入失敗"))
		return
	}

	err = l.ProductService.UpdateProduct(&product)
	if err != nil {
		c.JSON(200, utils.Error(400, "product 更新失敗"))
		return
	}

	// 3. 返回结果
	c.JSON(200, utils.Success("", "更新成功"))
}

// 根據分類ID獲取數據
func (l *ProductController) GetProductsByCategoryID(c *gin.Context) {
	categoryID := c.Param("categoryID")
	categoryIDNum, err := strconv.Atoi(categoryID)
	if err != nil {
		c.JSON(200, utils.Error(400, "categoryID传入错误"))
		return
	}
	result, err := l.ProductService.GetProductsByCategoryID(categoryIDNum)
	if err != nil {
		c.JSON(200, utils.Error(400, "通過categoryID獲取商品失敗"))
	}
	c.JSON(200, utils.Success(result, "通過categoryID成功獲取商品"))
}

// 根據品牌ID獲取數據
func (l *ProductController) GetProductsByBrandID(c *gin.Context) {
	brandID := c.Param("brandID")
	brandIDNum, err := strconv.Atoi(brandID)
	if err != nil {
		c.JSON(200, utils.Error(400, "brandID传入错误"))
		return
	}

	result, err := l.ProductService.GetProductsByBrandID(brandIDNum)
	if err != nil {
		c.JSON(200, utils.Error(400, "通過BrandID獲取商品失败"))
	}
	c.JSON(200, utils.Success(result, "通過BrandID成功獲取商品"))
}

// 根據名稱獲取數據
func (l *ProductController) GetProductsByName(c *gin.Context) {
	name := c.Param("name")

	result, err := l.ProductService.GetProductsByName(name)
	if err != nil {
		c.JSON(200, utils.Error(400, "通過name獲取商品失败"))
	}

	c.JSON(200, utils.Success(result, "通過name成功獲取商品"))

}

// 根據頁數獲取數據
func (l *ProductController) GetProductsByPage(c *gin.Context) {
	var request models.SearchProductsRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(200, utils.Error(400, "request格式傳入失敗"))
		return
	}

	result, err := l.ProductService.GetProductsByPage(request)
	if err != nil {
		c.JSON(200, utils.Error(400, "獲取商品失敗"))
	}

	c.JSON(200, utils.Success(result, "獲取商品成功"))
}
