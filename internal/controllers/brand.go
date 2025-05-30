package controllers

import (
	"fmt"
	"shopping/internal/models"
	"shopping/internal/services"
	"shopping/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BrandController struct {
	BrandService *services.BrandService
}

func NewBrandController(BrandService *services.BrandService) *BrandController {
	return &BrandController{BrandService: BrandService}
}

func (l *BrandController) GetBrandByPage(c *gin.Context) {

	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageNum, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(200, utils.Error(400, "page参数错误"))
		return
	}
	pageSizeNum, err := strconv.Atoi(pageSize)
	if err != nil {
		c.JSON(200, utils.Error(400, "pageSize参数错误"))
		return
	}
	result, err := l.BrandService.GetBrandByPage(pageNum, pageSizeNum)
	if err != nil {
		c.JSON(200, utils.Error(400, "获取失败"))
		return
	}
	// 3. 返回结果
	c.JSON(200, utils.Success(result, "登录成功"))

}

// AddBrand 处理添加品牌请求
func (l *BrandController) AddBrand(c *gin.Context) {
	var brand models.Brand

	// 绑定请求数据到 certificate 结构体
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(200, utils.Error(400, "请求数据解析失败"))
		return
	}
	fmt.Println(brand)
	// 调用服务层方法添加品牌
	if err := l.BrandService.AddBrand(&brand); err != nil {
		c.JSON(200, utils.Error(400, "添加品牌失败"))
		return
	}

	c.JSON(200, utils.Success(nil, "品牌添加成功"))
}

// UpdateBrand 处理更新品牌请求
func (l *BrandController) UpdateBrand(c *gin.Context) {
	var brand models.Brand

	// 绑定请求数据到 certificate 结构体
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(200, utils.Error(400, "请求数据解析失败"))
		return
	}

	// 调用服务层方法更新品牌
	if err := l.BrandService.UpdateBrand(&brand); err != nil {
		c.JSON(200, utils.Error(400, "更新品牌失败"))
		return
	}

	c.JSON(200, utils.Success(nil, "品牌更新成功"))
}

// DeleteBrand 处理删除品牌请求
func (l *BrandController) DeleteBrand(c *gin.Context) {
	// 从请求路径中获取品牌 ID
	id := c.Param("id")
	if id == "" {
		c.JSON(200, utils.Error(400, "品牌 ID 不能为空"))
		return
	}

	// 调用服务层方法标记品牌为已删除
	if err := l.BrandService.DeleteBrand(id); err != nil {
		c.JSON(200, utils.Error(400, "品牌删除失败"))
		return
	}

	c.JSON(200, utils.Success(nil, "品牌已标记为删除"))
}

func (l *BrandController) GetAllBrand(c *gin.Context) {
	result, err := l.BrandService.GetAllBrand()
	if err != nil {
		c.JSON(200, utils.Error(400, "品牌获取失败"))
		return
	}

	c.JSON(200, utils.Success(result, "品牌获取成功"))
}
