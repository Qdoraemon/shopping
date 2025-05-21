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

func (r *ProductRepository) GetProductById(id int) (*models.Product, []*models.Specification, []*models.Option, []models.ProductVariant, error) {
	var product = new(models.Product)
	err := r.engine.Table(models.Product{}.TableName()).Where("id = ?", id).Find(product).Error
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var specifications []*models.Specification
	err = r.engine.Table(models.Specification{}.TableName()).Where("product_id = ?", id).Find(&specifications).Error
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var options []*models.Option
	err = r.engine.Table(models.Option{}.TableName()).Where("product_id = ?", id).Find(&options).Error
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var variants []models.ProductVariant
	err = r.engine.Table(models.ProductVariant{}.TableName()).Where("product_id = ?", id).Find(&variants).Error
	if err != nil {
		return nil, nil, nil, nil, err
	}

	// product.Variants = variants

	return product, specifications, options, variants, nil
}

func (r *ProductRepository) AddProduct(product *models.Product, options []*models.Option, specifications []*models.Specification) error {
	// 开启事务
	tx := r.engine.Begin()

	// 插入产品数据
	if err := tx.Create(product).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 插入选项数据
	for _, option := range options {
		option.ProductID = product.ID
		if err := tx.Create(option).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 插入规格数据
	for _, specification := range specifications {
		specification.ProductID = product.ID
		if err := tx.Create(specification).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	return tx.Commit().Error
}
