package repositories

import (
	"shopping/internal/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	engine *gorm.DB
}

func NewProductRepository(engine *gorm.DB) *ProductRepository {
	return &ProductRepository{engine: engine}
}

func (r *ProductRepository) GetAllProducts() ([]*models.Product, error) {
	var products []*models.Product
	err := r.engine.Table(models.Product{}.TableName()).Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetProductById(id int) (*models.Product, []*models.Specification, error) {
	var product = new(models.Product)
	err := r.engine.Table(models.Product{}.TableName()).Where("id = ?", id).Find(product).Error
	if err != nil {
		return nil, nil, err
	}

	var specifications []*models.Specification
	err = r.engine.Table(models.Specification{}.TableName()).Where("product_id = ?", id).Find(&specifications).Error
	if err != nil {
		return nil, nil, err
	}

	return product, specifications, nil
}
