package services

import (
	"shopping/internal/models"
	"shopping/internal/repositories"
	"time"

	"gorm.io/gorm"
)

type option struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

var TypeMap = map[int]string{
	1: "平板",
	2: "電腦",
	3: "手機",
	4: "智能手錶",
}

type Product struct {
	ID            int64     `json:"id"`
	Name          string    `json:"name"`
	CoverImage    string    `json:"coverImage"`
	DetailImages  []string  `json:"detailImages"`
	Description   string    `json:"description"`
	SalePrice     float64   `json:"salePrice"`
	CostPrice     float64   `json:"costPrice"`
	StockQuantity int       `json:"stockQuantity"`
	Brand         *string   `json:"brand"`
	Options       []option  `json:"options"`
	CategoryID    int       `json:"categoryId"`
	UpdateTime    time.Time `json:"updateTime"`
	CreateTime    time.Time `json:"createTime"`
	IsAvailable   bool      `json:"isAvailable"`
	IsDeleted     bool      `json:"isDeleted"`
	Type          string    `json:"type"`
	Model         string    `json:"model"`
}

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductsService(engine *gorm.DB) *ProductService {
	return &ProductService{productRepo: repositories.NewProductRepository(engine)}
}

func (us *ProductService) GetAllProducts() ([]*models.Product, error) {
	return us.productRepo.GetAllProducts()
}

func (us *ProductService) GetProductById(id int) (*Product, error) {
	product, err := us.productRepo.GetProductById(id)

	if err != nil {
		return nil, err
	}

	// 将 color 和 storage 合并为 options
	options := []option{
		{Name: "顔色", Values: product.Color},
		{Name: "存儲", Values: product.Storage},
	}

	// 创建新的 Product 结构体并赋值
	newProduct := Product{
		ID:            product.ID,
		Name:          product.Name,
		CoverImage:    product.CoverImage,
		DetailImages:  product.DetailImages,
		Description:   product.Description,
		SalePrice:     product.SalePrice,
		CostPrice:     product.CostPrice,
		StockQuantity: product.StockQuantity,
		Brand:         product.Brand,
		Options:       options,
		CategoryID:    product.CategoryID,
		UpdateTime:    product.UpdateTime,
		CreateTime:    product.CreateTime,
		IsAvailable:   product.IsAvailable,
		IsDeleted:     product.IsDeleted,
		Type:          TypeMap[product.Type],
	}

	// return us.productRepo.GetProductById(id)
	return &newProduct, nil
}
