package controllers

import (
	"shopping/internal/services"
	"shopping/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CertificatesController struct {
	CertificatesService *services.CertificatesService
}

func NewCertificatesController(CertificatesService *services.CertificatesService) *CertificatesController {
	return &CertificatesController{CertificatesService: CertificatesService}
}

func (l *CertificatesController) GetCertificatesByPage(c *gin.Context) {

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
	result, err := l.CertificatesService.GetCertificatesByPage(pageNum, pageSizeNum)
	if err != nil {
		c.JSON(200, utils.Error(400, "获取失败"))
		return
	}
	// 3. 返回结果
	c.JSON(200, utils.Success(result, "登录成功"))

}
