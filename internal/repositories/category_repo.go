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
