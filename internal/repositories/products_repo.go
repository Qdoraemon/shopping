package repositories

import (
	"shopping/internal/models"
	"strings"
	"time"

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

func (r *ProductRepository) AddProduct(product *models.Product, options []*models.Option, specifications []*models.Specification, variants []models.ProductVariant) error {
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

	// // 插入变体数据
	for _, variant := range variants {
		variant.ProductID = product.ID
		if err := tx.Create(&variant).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事务
	return tx.Commit().Error
}

// DeleteProduct deletes a product by its ID from the database
func (r *ProductRepository) DeleteProduct(id int) error {
	// 开启事务
	tx := r.engine.Begin()

	// 删除选项数据
	if err := tx.Where("product_id = ?", id).Delete(&models.Option{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除规格数据
	if err := tx.Where("product_id = ?", id).Delete(&models.Specification{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除变体数据
	if err := tx.Where("product_id = ?", id).Delete(&models.ProductVariant{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// 删除产品数据
	if err := tx.Delete(&models.Product{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}

// DeleteProducts deletes multiple products by their IDs from the database
func (r *ProductRepository) DeleteProducts(ids []int) error {
	// 开启事务
	tx := r.engine.Begin()
	// 删除选项数据
	if err := tx.Where("product_id IN ?", ids).Delete(&models.Option{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除规格数据
	if err := tx.Where("product_id IN ?", ids).Delete(&models.Specification{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除变体数据
	if err := tx.Where("product_id IN ?", ids).Delete(&models.ProductVariant{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 删除产品数据
	if err := tx.Where("id IN ?", ids).Delete(&models.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	return tx.Commit().Error
}

func (r *ProductRepository) CopyProduct(id int) (*models.Product, error) {
	var product models.Product
	if err := r.engine.Table(models.Product{}.TableName()).Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	// 复制产品数据
	newProduct := product
	newProduct.ID = 0                  // 重置ID以生成新记录
	newProduct.CreateTime = time.Now() // 更新创建时间
	newProduct.Status = "0"            // 更新状态

	// 判断 newProduct.Name 是否包含 "副本"
	if !strings.Contains(newProduct.Name, "副本") {
		newProduct.Name = newProduct.Name + " 副本"
	}

	// 开启事务
	tx := r.engine.Begin()

	if err := tx.Create(&newProduct).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 复制选项数据
	var options []models.Option
	if err := r.engine.Table(models.Option{}.TableName()).Where("product_id = ?", id).Find(&options).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, option := range options {
		option.ID = 0 // 重置ID以生成新记录
		option.ProductID = newProduct.ID
		if err := tx.Create(&option).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// 复制规格数据
	var specifications []models.Specification
	if err := r.engine.Table(models.Specification{}.TableName()).Where("product_id = ?", id).Find(&specifications).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, specification := range specifications {
		specification.ID = 0 // 重置ID以生成新记录
		specification.ProductID = newProduct.ID
		if err := tx.Create(&specification).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// 复制变体数据
	var variants []models.ProductVariant
	if err := r.engine.Table(models.ProductVariant{}.TableName()).Where("product_id = ?", id).Find(&variants).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, variant := range variants {
		variant.ID = 0 // 重置ID以生成新记录
		variant.ProductID = newProduct.ID
		if err := tx.Create(&variant).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &newProduct, nil
}

// UpdateProduct updates a product and its related options, specifications, and variants in the database
func (r *ProductRepository) UpdateProduct(product *models.Product, options []*models.Option, specifications []*models.Specification, variants []models.ProductVariant) error {
	// 開始一個新事務
	tx := r.engine.Begin()

	// 更新產品數據
	if err := tx.Save(product).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 刪除舊的選項數據
	if err := tx.Where("product_id = ?", product.ID).Delete(&models.Option{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 插入新的選項數據
	for _, option := range options {
		option.ProductID = product.ID // 使用產品ID
		if err := tx.Create(option).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 刪除舊的規格數據
	if err := tx.Where("product_id = ?", product.ID).Delete(&models.Specification{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 插入新的規格數據
	for _, specification := range specifications {
		specification.ProductID = product.ID // 使用產品ID
		if err := tx.Create(specification).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 刪除舊的變體數據
	if err := tx.Where("product_id = ?", product.ID).Delete(&models.ProductVariant{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 插入新的變體數據
	for _, variant := range variants {
		variant.ProductID = product.ID // 使用產品ID
		if err := tx.Create(&variant).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// 提交事務
	return tx.Commit().Error
}
