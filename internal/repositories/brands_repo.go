package repositories

import (
	"shopping/internal/models"

	"gorm.io/gorm"
)

type BrandRepository struct {
	engine *gorm.DB
}

func NewBrandRepository(engine *gorm.DB) *BrandRepository {
	return &BrandRepository{engine: engine}
}

// GetBrandByPage 分页查询未删除的证书
func (r *BrandRepository) GetBrandByPage(page, pageSize int) ([]*models.Brand, error) {
	var Brand []*models.Brand
	err := r.engine.
		Table(models.Brand{}.TableName()).
		Where("is_deleted = 1").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&Brand).Error
	return Brand, err
}

func (r *BrandRepository) AddBrand(certificate *models.Brand) error {
	err := r.engine.Table(models.Brand{}.TableName()).Create(certificate).Error
	return err
}

// UpdateBrand 根据证书 ID 更新证书信息
func (r *BrandRepository) UpdateBrand(certificate *models.Brand) error {
	// 这里假设使用 ID 作为唯一标识进行更新
	err := r.engine.Table(models.Brand{}.TableName()).
		Where("id = ?", certificate.ID).
		Updates(certificate).
		Error
	return err
}

// DeleteBrand 根据证书 ID 标记证书为已删除
func (r *BrandRepository) DeleteBrand(id interface{}) error {
	updateData := map[string]interface{}{
		"is_deleted": 0,
	}
	err := r.engine.Table(models.Brand{}.TableName()).
		Where("id = ?", id).
		Updates(updateData).
		Error
	return err
}

// 獲取所有的品牌
func (r *BrandRepository) GetAllBrand() ([]*models.Brand, error) {
	var Brand []*models.Brand
	err := r.engine.Table(models.Brand{}.TableName()).Find(&Brand).Error
	return Brand, err
}
