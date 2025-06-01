package repositories

import (
	"shopping/internal/models"

	"gorm.io/gorm"
)

type CategoriesRepository struct {
	engine *gorm.DB
}

func NewCategoriesRepository(engine *gorm.DB) *CategoriesRepository {
	return &CategoriesRepository{engine: engine}
}

func (r *CategoriesRepository) GetAllCategories() ([]models.Categories, error) {
	var categories []models.Categories
	result := r.engine.Find(&categories)
	return categories, result.Error
}

func (r *CategoriesRepository) GetCategoriesByPage(request models.SearchCategoriesRequest) ([]*models.Categories, int64, error) {
	var Categories []*models.Categories
	var total int64

	// 初始化查询条件
	query := r.engine.Model(&models.Categories{})

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
	err := query.Find(&Categories).Error

	if err != nil {
		return nil, 0, err
	}
	return Categories, total, nil
}

func (r *CategoriesRepository) AddCategory(category *models.Categories) error {
	err := r.engine.Table(models.Categories{}.TableName()).Create(category).Error
	return err
}

// UpdateBrand 根据ID 更新信息
func (r *CategoriesRepository) UpdateCategory(category *models.Categories) error {
	// 这里假设使用 ID 作为唯一标识进行更新

	err := r.engine.Table(models.Categories{}.TableName()).
		Save(category).
		Error

	return err
}

// DeleteBrand 根据 ID 标记为已删除
func (r *CategoriesRepository) DeleteCategory(id string) error {
	// 开启事务
	tx := r.engine.Begin()
	// 删除选项数据
	if err := tx.Where("id = ?", id).Delete(&models.Categories{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
