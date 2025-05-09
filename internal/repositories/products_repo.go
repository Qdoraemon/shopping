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

func (r *ProductRepository) GetProductById(id int) (*models.Product, error) {
	var product = new(models.Product)
	err := r.engine.Table(models.Product{}.TableName()).Where("id = ?", id).Find(product).Error
	if err != nil {
		return nil, err
	}

	var specifications []*models.Specification
	err = r.engine.Table(models.Specification{}.TableName()).Where("product_id = ?", id).Find(&specifications).Error
	if err != nil {
		return nil, err
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

	product.Specifications = formattedSpecifications

	return product, nil
}
