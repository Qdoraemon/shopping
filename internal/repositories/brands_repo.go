package repositories

import (
	"shopping/internal/models"

	"gorm.io/gorm"
)

type BrandsRepository struct {
	engine *gorm.DB
}

func NewBrandsRepository(engine *gorm.DB) *BrandsRepository {
	return &BrandsRepository{engine: engine}
}

// GetBrandsByPage 分页查询未删除的证书
func (r *BrandsRepository) GetBrandsByPage(page, pageSize int) ([]*models.Brands, error) {
	var brands []*models.Brands
	err := r.engine.
		Table(models.Brands{}.TableName()).
		Where("is_deleted = 1").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&brands).Error
	return brands, err
}

func (r *BrandsRepository) AddBrands(certificate *models.Brands) error {
	err := r.engine.Table(models.Brands{}.TableName()).Create(certificate).Error
	return err
}

// UpdateBrands 根据证书 ID 更新证书信息
func (r *BrandsRepository) UpdateBrands(certificate *models.Brands) error {
	// 这里假设使用 ID 作为唯一标识进行更新
	err := r.engine.Table(models.Brands{}.TableName()).
		Where("id = ?", certificate.ID).
		Updates(certificate).
		Error
	return err
}

// DeleteBrands 根据证书 ID 标记证书为已删除
func (r *BrandsRepository) DeleteBrands(id interface{}) error {
	updateData := map[string]interface{}{
		"is_deleted": 0,
	}
	err := r.engine.Table(models.Brands{}.TableName()).
		Where("id = ?", id).
		Updates(updateData).
		Error
	return err
}
