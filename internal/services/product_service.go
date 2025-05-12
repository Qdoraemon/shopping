package services

import (
	"shopping/internal/models"
	"shopping/internal/repositories"

	"gorm.io/gorm"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductsService(engine *gorm.DB) *ProductService {
	return &ProductService{productRepo: repositories.NewProductRepository(engine)}
}

func (us *ProductService) GetAllProducts() ([]*models.Product, error) {
	return us.productRepo.GetAllProducts()
}

func (s *ProductService) GetProductDetails(id int) (*models.Product, []map[string]interface{}, error) {
	product, specifications, err := s.productRepo.GetProductById(id)
	if err != nil {
		return nil, nil, err
	}

	// 将 specifications 按 category 分组
	specMap := make(map[string][]map[string]string)
	for _, spec := range specifications {
		item := map[string]string{"name": spec.Name, "value": spec.Value}
		specMap[spec.Category] = append(specMap[spec.Category], item)
	}

	// 将分组后的 specifications 转换为所需格式
	var formattedSpecifications []map[string]interface{}
	for category, items := range specMap {
		formattedSpecifications = append(formattedSpecifications, map[string]interface{}{
			"category": category,
			"items":    items,
		})
	}

	return product, formattedSpecifications, nil
}

func (us *ProductService) GetProductById(id int) (*models.FrontendProduct, error) {
	product, specifications, err := us.GetProductDetails(id)

	if err != nil {
		return nil, err
	}

	// 将 color 和 storage 合并为 options
	options := []models.Option{
		{Name: "顔色", Values: product.Color},
		{Name: "存儲", Values: product.Storage},
	}

	// 创建新的 Product 结构体并赋值
	newProduct := models.FrontendProduct{
		ID:             product.ID,
		Name:           product.Name,
		CoverImage:     product.CoverImage,
		DetailImages:   product.DetailImages,
		Description:    product.Description,
		SalePrice:      product.SalePrice,
		CostPrice:      product.CostPrice,
		StockQuantity:  product.StockQuantity,
		Brand:          product.Brand,
		Options:        options,
		Category:       models.CategoryMap[product.CategoryID],
		Features:       product.Features,
		UpdateTime:     product.UpdateTime,
		CreateTime:     product.CreateTime,
		IsAvailable:    product.IsAvailable,
		IsDeleted:      product.IsDeleted,
		Type:           product.Type,
		Specifications: specifications,
	}
	// fmt.Println("product:", newProduct.Specifications)
	// return us.productRepo.GetProductById(id)
	return &newProduct, nil
}
