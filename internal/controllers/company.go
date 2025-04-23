package controllers

import "github.com/gin-gonic/gin"

// Company 定义公司详情结构体
type Company struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}

// 模拟公司详情数据
var CompanyInfo = Company{
	Name:        "示例购物公司",
	Description: "一家专注于提供优质购物服务的公司",
	Address:     "示例市示例路 123 号",
	Phone:       "13800138000",
	Email:       "info@example.com",
}

// GetCompanyInfo 处理公司详情展示请求
func GetCompanyInfo(c *gin.Context) {
	c.JSON(200, CompanyInfo)
}
