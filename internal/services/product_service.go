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

func (us *ProductService) GetProductById(id int) (*models.Product, error) {
	return us.productRepo.GetProductById(id)
}
