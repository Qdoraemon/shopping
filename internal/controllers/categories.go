package controllers

import (
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
