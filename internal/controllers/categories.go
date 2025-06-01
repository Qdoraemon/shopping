package controllers

import (
	"shopping/internal/models"
	"shopping/internal/services"
	"shopping/internal/utils"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	Service *services.CategoryService
}

func NewCategoryController(service *services.CategoryService) *CategoryController {
	return &CategoryController{Service: service}
}

func (controller *CategoryController) GetAllCategories(c *gin.Context) {
	categories, err := controller.Service.GetAllCategories()
	if err != nil {
		c.JSON(200, utils.Error(400, err.Error()))
		return
	}
	c.JSON(200, utils.Success(categories, "獲取商品分類成功"))
}

func (l *CategoryController) GetCategoriesByPage(c *gin.Context) {

	var request models.SearchCategoriesRequest

	// 解析请求参数
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(200, utils.Error(400, err.Error()))
		return
	}

	// 调用服务层方法获取品牌列表
	result, err := l.Service.GetCategoriesByPage(request)
	if err != nil {
		c.JSON(200, utils.Error(400, "获取失败"))
		return
	}
	// 3. 返回结果
	c.JSON(200, utils.Success(result, "登录成功"))

}

// AddBrand 处理添加品牌请求
func (l *CategoryController) AddCategory(c *gin.Context) {
	var category models.Categories

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(200, utils.Error(400, "请求数据解析失败"))
		return
	}

	if err := l.Service.AddCategory(&category); err != nil {
		c.JSON(200, utils.Error(400, "添加分類失败"))
		return
	}

	c.JSON(200, utils.Success(nil, "分類添加成功"))
}

func (l *CategoryController) UpdateCategory(c *gin.Context) {
	var category models.Categories

	// 绑定请求数据到 certificate 结构体
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(200, utils.Error(400, "请求数据解析失败"+err.Error()))
		return
	}
	// fmt.Println(brand)
	// 调用服务层方法更新品牌
	if err := l.Service.UpdateCategory(&category); err != nil {
		c.JSON(200, utils.Error(400, "更新分類失败"))
		return
	}

	c.JSON(200, utils.Success(nil, "分類更新成功"))
}

// DeleteBrand 处理删除品牌请求
func (l *CategoryController) DeleteCategory(c *gin.Context) {
	// 从请求路径中获取品牌 ID
	id := c.Param("id")
	if id == "" {
		c.JSON(200, utils.Error(400, "分類 ID 不能为空"))
		return
	}

	// 调用服务层方法标记品牌为已删除
	if err := l.Service.DeleteCategory(id); err != nil {
		c.JSON(200, utils.Error(400, "分類删除失败"))
		return
	}

	c.JSON(200, utils.Success(nil, "分類删除成功"))
}
