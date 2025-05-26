package services

import (
	"shopping/internal/models"
	"shopping/internal/repositories"

	"gorm.io/gorm"
)

type CategoryService struct {
	repositoriesRepo *repositories.CategoriesRepository
}

func NewCategoryService(engine *gorm.DB) *CategoryService {
	return &CategoryService{repositoriesRepo: repositories.NewCategoriesRepository(engine)}
}

func (cs *CategoryService) GetAllCategories() ([]models.Categories, error) {
	return cs.repositoriesRepo.GetAllCategories()
}
