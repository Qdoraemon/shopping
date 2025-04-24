package controllers

import (
	"shopping/internal/services"
	"shopping/internal/utils"

	"github.com/gin-gonic/gin"
)

type CarouselsController struct {
	CarouselsService *services.CarouselsService
}

func NewCarouselsController(CarouselsService *services.CarouselsService) *CarouselsController {
	return &CarouselsController{CarouselsService: CarouselsService}
}

func (l *CarouselsController) GetAllCarousels(c *gin.Context) {

	result, err := l.CarouselsService.GetAllCarousels()
	if err != nil {
		c.JSON(200, utils.Error(400, "获取失败"))
		return
	}
	// 3. 返回结果
	c.JSON(200, utils.Success(result, "登录成功"))

}
