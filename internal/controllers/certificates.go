package controllers

import (
	"shopping/internal/models"
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

// AddCertificate 处理添加证书请求
func (l *CertificatesController) AddCertificate(c *gin.Context) {
	var certificate models.Certificates
	// 绑定请求数据到 certificate 结构体
	if err := c.ShouldBindJSON(&certificate); err != nil {
		c.JSON(200, utils.Error(400, "请求数据解析失败"))
		return
	}

	// 调用服务层方法添加证书
	if err := l.CertificatesService.AddCertificate(&certificate); err != nil {
		c.JSON(200, utils.Error(400, "添加证书失败"))
		return
	}

	c.JSON(200, utils.Success(nil, "证书添加成功"))
}

// UpdateCertificate 处理更新证书请求
func (l *CertificatesController) UpdateCertificate(c *gin.Context) {
	var certificate models.Certificates

	// 绑定请求数据到 certificate 结构体
	if err := c.ShouldBindJSON(&certificate); err != nil {
		c.JSON(200, utils.Error(400, "请求数据解析失败"))
		return
	}

	// 调用服务层方法更新证书
	if err := l.CertificatesService.UpdateCertificate(&certificate); err != nil {
		c.JSON(200, utils.Error(400, "更新证书失败"))
		return
	}

	c.JSON(200, utils.Success(nil, "证书更新成功"))
}

// DeleteCertificate 处理删除证书请求
func (l *CertificatesController) DeleteCertificate(c *gin.Context) {
	// 从请求路径中获取证书 ID
	id := c.Param("id")
	if id == "" {
		c.JSON(200, utils.Error(400, "证书 ID 不能为空"))
		return
	}

	// 调用服务层方法标记证书为已删除
	if err := l.CertificatesService.DeleteCertificate(id); err != nil {
		c.JSON(200, utils.Error(400, "标记证书删除失败"))
		return
	}

	c.JSON(200, utils.Success(nil, "证书已标记为删除"))
}
