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
