package controllers

import (
	"shopping/internal/models"
	"shopping/internal/services"
	"shopping/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BrandsController struct {
	BrandsService *services.BrandsService
}

func NewBrandsController(BrandsService *services.BrandsService) *BrandsController {
	return &BrandsController{BrandsService: BrandsService}
}

func (l *BrandsController) GetBrandsByPage(c *gin.Context) {

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
	result, err := l.BrandsService.GetBrandsByPage(pageNum, pageSizeNum)
	if err != nil {
		c.JSON(200, utils.Error(400, "获取失败"))
		return
	}
	// 3. 返回结果
	c.JSON(200, utils.Success(result, "登录成功"))

}

// AddBrands 处理添加品牌请求
func (l *BrandsController) AddBrands(c *gin.Context) {
	var certificate models.Brands
	// 绑定请求数据到 certificate 结构体
	if err := c.ShouldBindJSON(&certificate); err != nil {
		c.JSON(200, utils.Error(400, "请求数据解析失败"))
		return
	}

	// 调用服务层方法添加品牌
	if err := l.BrandsService.AddBrands(&certificate); err != nil {
		c.JSON(200, utils.Error(400, "添加品牌失败"))
		return
	}

	c.JSON(200, utils.Success(nil, "品牌添加成功"))
}

// UpdateBrands 处理更新品牌请求
func (l *BrandsController) UpdateBrands(c *gin.Context) {
	var certificate models.Brands

	// 绑定请求数据到 certificate 结构体
	if err := c.ShouldBindJSON(&certificate); err != nil {
		c.JSON(200, utils.Error(400, "请求数据解析失败"))
		return
	}

	// 调用服务层方法更新品牌
	if err := l.BrandsService.UpdateBrands(&certificate); err != nil {
		c.JSON(200, utils.Error(400, "更新品牌失败"))
		return
	}

	c.JSON(200, utils.Success(nil, "品牌更新成功"))
}

// DeleteBrands 处理删除品牌请求
func (l *BrandsController) DeleteBrands(c *gin.Context) {
	// 从请求路径中获取品牌 ID
	id := c.Param("id")
	if id == "" {
		c.JSON(200, utils.Error(400, "品牌 ID 不能为空"))
		return
	}

	// 调用服务层方法标记品牌为已删除
	if err := l.BrandsService.DeleteBrands(id); err != nil {
		c.JSON(200, utils.Error(400, "品牌删除失败"))
		return
	}

	c.JSON(200, utils.Success(nil, "品牌已标记为删除"))
}

func (l *BrandsController) GetAllBrands(c *gin.Context) {
	result, err := l.BrandsService.GetAllBrands()
	if err != nil {
		c.JSON(200, utils.Error(400, "品牌获取失败"))
		return
	}

	c.JSON(200, utils.Success(result, "品牌获取成功"))
}
