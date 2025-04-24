package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Carousel struct {
	// ID
	ID int `json:"id"`
	// 標題
	Title string `json:"title"`
	// 排序
	Order int `json:"order"`
	// 图片URL地址
	ImageUrl string `json:"imageUrl"`
	// 图片URL地址
	RedirectUrl string `json:"redirectUrl"`
	// 是否启用
	IsEnabled bool `json:"isEnabled"`
	// 创建时间
	CreateTime string `json:"createTime"`
	// 修改时间
	UpdateTime string `json:"updateTime"`
	// 是否删除
	IsDeleted bool `json:"isDeleted"`
}

type Product struct {
	// 商品ID
	ID int `json:"id"`
	// 商品名稱
	Name string `json:"name"`
	// 商品封面圖片
	CoverImage string `json:"coverImage"`
	// 商品详情图片
	DetailImages string `json:"detailImages"`
	// 商品描述
	Description string `json:"description"`
	// 销售价格
	SalePrice int `json:"salePrice"`
	// 成本价格
	CostPrice int `json:"costPrice"`
	// 库存数量
	StockQuantity string `json:"stockQuantity"`
	// 商品品牌
	Brand string `json:"brand"`
	// 商品分类ID
	CategoryId int `json:"categoryId"`
	// 更新时间
	UpdateTime string `json:"updateTime"`
	// 创建时间
	CreateTime string `json:"createTime"`
	// 上下架状态
	IsAvailable bool `json:"isAvailable"`
	// 是否删除
	IsDeleted bool `json:"isDeleted"`
}

type BasicInformation struct {
	// ID
	ID int `json:"id"`
	// 首頁標題
	HomeTitle string `json:"homeTitle"`
	// 首页描述
	HomeDescription string `json:"homeDescription"`
	// 联系方式
	Phone string `json:"phone"`
	// 邮箱
	Email string `json:"email"`
	// 地址
	Address string `json:"address"`
	// 微信图片
	WeChatImage string `json:"weChatImage"`
	// 是否刪除
	IsDeleted bool `json:"isDeleted"`
	// 是否启用
	IsEnable bool `json:"isEnable"`
}

func main() {
	r := gin.Default()
	var Carousels []Carousel
	var Products []Product
	var BasicInformations []BasicInformation
	Carousels = append(Carousels, Carousel{
		ID:       1,
		Title:    "測試",
		Order:    1,
		ImageUrl: "https://i.ebayimg.com/images/g/EcIAAOSwropmIhRN/s-l1200.jpg",
	})

	Products = append(Products, Product{
		ID:          1,
		Name:        "ipad 10",
		Brand:       "Apple",
		CoverImage:  "https://i.ebayimg.com/images/g/EcIAAOSwropmIhRN/s-l1200.jpg",
		Description: "ipad 10 商品",
		SalePrice:   1000,
		CostPrice:   500,
		IsAvailable: true,
	})
	Products = append(Products, Product{
		ID:          2,
		Name:        "ipad 10",
		Brand:       "Apple",
		CoverImage:  "https://i.ebayimg.com/images/g/EcIAAOSwropmIhRN/s-l1200.jpg",
		Description: "ipad 10 商品",
		SalePrice:   1000,
		CostPrice:   500,
		IsAvailable: true,
	})
	Products = append(Products, Product{
		ID:          3,
		Name:        "ipad 10",
		Brand:       "Apple",
		CoverImage:  "https://i.ebayimg.com/images/g/EcIAAOSwropmIhRN/s-l1200.jpg",
		Description: "ipad 10 商品",
		SalePrice:   1000,
		CostPrice:   500,
		IsAvailable: true,
	})
	Products = append(Products, Product{
		ID:          4,
		Name:        "ipad 10",
		Brand:       "Apple",
		CoverImage:  "https://i.ebayimg.com/images/g/EcIAAOSwropmIhRN/s-l1200.jpg",
		Description: "ipad 10 商品",
		SalePrice:   1000,
		CostPrice:   500,
		IsAvailable: true,
	})

	BasicInformations = append(BasicInformations, BasicInformation{
		ID:          123,
		HomeTitle:   "測試用戶",
		Phone:       "1234567890",
		Email:       "123@test.com",
		Address:     "測試地址",
		WeChatImage: "@assets/barcode",
	})

	// Carousels = append(Carousels, Carousel{
	// 	ID:       1,
	// 	Title:    "測試",
	// 	Order:    1,
	// 	ImageUrl: "@/assets/2.png",
	// })
	// Carousels = append(Carousels, Carousel{
	// 	ID:       1,
	// 	Title:    "測試",
	// 	Order:    1,
	// 	ImageUrl: "https://i.ebayimg.com/images/g/EcIAAOSwropmIhRN/s-l1200.jpg",
	// })

	r.GET("/v1/carousels/getAllCarousels", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": Carousels,
		})
		// fmt.Println(Carousels)
	})

	r.GET("/v1/basicInformation/getAllBasicInformation", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": 100,
		})
	})

	r.GET("/v1/certificates/getAllCertificates", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": 100,
		})
	})

	r.GET("/v1/products/getAllProducts", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": Products,
		})
	})

	r.GET("/v1/products/getProductById/:id", func(c *gin.Context) {
		id := c.Param("id")
		i, _ := strconv.ParseInt(id, 10, 64)
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": Products[i],
		})
	})

	r.Run("127.0.0.1:8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
