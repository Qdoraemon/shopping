package controllers

import "github.com/gin-gonic/gin"

// Product 定义商品结构体
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Desc  string  `json:"desc"`
}

// 模拟商品数据
var Products = []Product{
	{ID: 1, Name: "苹果", Price: 5.99, Desc: "新鲜的苹果"},
	{ID: 2, Name: "香蕉", Price: 3.99, Desc: "香甜的香蕉"},
}

// GetProducts 处理商品展示请求
func GetProducts(c *gin.Context) {
	c.JSON(200, Products)
}
