package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"shopping/internal/services"
	"shopping/internal/utils"
	"strconv"
	"time"

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

// UploadImage 处理图片上传请求
func (l *CertificatesController) UploadImage(c *gin.Context) {
	// 从表单中获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, utils.Error(400, "获取文件失败"))
		return
	}
	// 检查文件类型，这里简单判断文件后缀是否为图片
	if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
		c.JSON(200, utils.Error(400, "文件类型不支持"))
		return
	}
	// 检查文件大小，这里简单判断文件大小不超过 2MB
	if file.Size > 2*1024*1024 {
		c.JSON(200, utils.Error(400, "文件大小超过限制"))
		return
	}
	fmt.Println("00ol")
	// 修改图片名称，这里简单使用当前时间戳作为文件名
	// SaveFileName := fmt.Sprintf("%d%s", time.Now().UnixMicro(), filepath.Ext(file.Filename))
	SaveFileName := fmt.Sprintf("shopping_%d", time.Now().UnixMicro())

	// 定义保存文件的路径，这里简单保存到项目根目录下的 uploads 文件夹
	dest := fmt.Sprintf("./uploads/%s%s", SaveFileName, filepath.Ext(file.Filename))
	// 保存文件
	if err := c.SaveUploadedFile(file, dest); err != nil {
		c.JSON(200, utils.Error(400, "保存文件失败"))
		return
	}

	c.JSON(200, utils.Success(fmt.Sprintf("%s.jfif", SaveFileName), "上传成功"))
}

// GetImage 动态读取图片并返回给客户端
func (l *CertificatesController) GetImage(c *gin.Context) {

	filename := c.DefaultQuery("fileName", "")

	if filename == "" {
		c.JSON(200, utils.Error(400, "读取文件为空"))
		return
	}
	filePath := "./uploads/" + filename

	// 读取文件内容
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		c.JSON(200, utils.Error(400, "读取文件失败"))
		return
	}

	// 设置响应头
	c.Header("Content-Type", "image/jpeg") // 根据实际情况修改图片类型
	c.Data(200, "image/jpeg", fileContent)
}
