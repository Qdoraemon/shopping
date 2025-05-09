package controllers

import (
	"shopping/internal/services"
	"shopping/internal/utils"

	"github.com/gin-gonic/gin"
)

type BaseInfoController struct {
	BaseInfoService *services.BaseInfoService
}

func NewBaseInfoController(BaseInfoService *services.BaseInfoService) *BaseInfoController {
	return &BaseInfoController{BaseInfoService: BaseInfoService}
}

func (l *BaseInfoController) GetAllBasicInformation(c *gin.Context) {
	/*
		userId := c.GetInt("userId")
		username := c.GetString("username")
	*/

	result, err := l.BaseInfoService.GetAllBasicInformation()
	if err != nil {
		c.JSON(200, utils.Error(400, "获取失败"))
		return
	}
	// 3. 返回结果
	c.JSON(200, utils.Success(result, "登录成功"))

}
