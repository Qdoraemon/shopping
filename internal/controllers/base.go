package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"shopping/internal/utils"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

// UploadImage 处理图片上传请求
func (l *BaseController) UploadImage(c *gin.Context) {
	// 从表单中获取上传的文件
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(200, utils.Error(400, "获取文件失败"))
		return
	}
	files := form.File["files"]

	var uploadedFiles []string

	for _, file := range files {
		// 检查文件类型，这里简单判断文件后缀是否为图片
		if file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png" {
			c.JSON(200, utils.Error(400, "文件类型不支持"))
			return
		}
		// // 检查文件大小，这里简单判断文件大小不超过 2MB
		if file.Size > 5*1024*1024 {
			c.JSON(200, utils.Error(400, "文件大小超过限制"))
			return
		}
		// 修改图片名称，这里简单使用UUID戳作为文件名
		SaveFileName := utils.GenFileNameByUUID("Image")

		// 定义保存文件的路径，UUID生成對應的文件名
		dest := fmt.Sprintf("./uploads/%s%s", SaveFileName, filepath.Ext(file.Filename))
		// 保存文件
		if err := c.SaveUploadedFile(file, dest); err != nil {
			c.JSON(200, utils.Error(400, "保存文件失败"))
			return
		}

		uploadedFiles = append(uploadedFiles, fmt.Sprintf("%s%s", SaveFileName, filepath.Ext(file.Filename)))
	}

	// fmt.Println(uploadedFiles)
	c.JSON(200, utils.Success(uploadedFiles, "上传成功"))
}

// GetImage 动态读取图片并返回给客户端
func (l *BaseController) GetImage(c *gin.Context) {

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

// 刪除圖片
func (l *BaseController) DeleteImage(c *gin.Context) {
	var req struct {
		ImageUrl string `json:"imageUrl"`
	}

	// 解析请求体中的参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, utils.Error(400, "无效请求"))
		return
	}

	// 删除文件
	err := os.Remove("./uploads/" + req.ImageUrl) // 这里假设 URL 是文件的路径
	if err != nil {
		c.JSON(200, utils.Error(400, "删除文件失败"))
		return
	}

	// 假设这里还会删除数据库中的记录

	c.JSON(200, utils.Success(nil, "删除成功"))
}
