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

// GetBrandByPage 分页查询的证书
func (r *BrandRepository) GetBrandByPage(request models.SearchBrandsRequest) ([]*models.Brand, int64, error) {
	// var Brand []*models.Brand
	// err := r.engine.
	// 	Table(models.Brand{}.TableName()).
	// 	Limit(pageSize).
	// 	Offset((page - 1) * pageSize).
	// 	Find(&Brand).Error
	// return Brand, err
	var brands []*models.Brand
	var total int64

	// 初始化查询条件
	query := r.engine.Model(&models.Brand{})

	// 根据名称进行模糊查询
	if request.Name != "" {
		query = query.Where("name LIKE ? OR name_en LIKE ?", "%"+request.Name+"%", "%"+request.Name+"%")
	}

	// 根据状态筛选
	if request.IsEnabled != "" {
		query = query.Where("is_enabled = ?", request.IsEnabled)
	}

	// 获取总记录数
	query.Count(&total)

	// 分页查询
	query = query.Offset((request.Page - 1) * request.PageSize).Limit(request.PageSize)

	// 执行查询
	err := query.Find(&brands).Error

	if err != nil {
		return nil, 0, err
	}
	return brands, total, nil
}

func (r *BrandRepository) AddBrand(brand *models.Brand) error {
	err := r.engine.Table(models.Brand{}.TableName()).Create(brand).Error
	return err
}

// UpdateBrand 根据ID 更新信息
func (r *BrandRepository) UpdateBrand(brand *models.Brand) error {
	// 这里假设使用 ID 作为唯一标识进行更新
	// fmt.Println(brand.ID)
	err := r.engine.Table(models.Brand{}.TableName()).
		Save(brand).
		Error

	return err
}

// DeleteBrand 根据 ID 标记为已删除
func (r *BrandRepository) DeleteBrand(id string) error {
	// 开启事务
	tx := r.engine.Begin()
	// 删除选项数据
	if err := tx.Where("id = ?", id).Delete(&models.Brand{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

// 獲取所有的品牌
func (r *BrandRepository) GetAllBrand() ([]*models.Brand, error) {
	var Brand []*models.Brand
	err := r.engine.Table(models.Brand{}.TableName()).Find(&Brand).Error
	return Brand, err
}
