package services

import (
	"errors"
	"shopping/internal/models"
	"shopping/internal/repositories"
	"time"

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

func (cs *CategoryService) GetCategoriesByPage(request models.SearchCategoriesRequest) (models.SearchCategoriesResponse, error) {
	categories, total, err := cs.repositoriesRepo.GetCategoriesByPage(request)

	if err != nil {
		return models.SearchCategoriesResponse{
			Categories: nil,
			Total:      0,
		}, err
	}

	return models.SearchCategoriesResponse{
		Categories: categories,
		Total:      total,
	}, nil
}

func (cs *CategoryService) AddCategory(category *models.Categories) error {
	if category == nil {
		return errors.New("category cannot be nil")
	}
	category.CreateTime = time.Now()
	category.UpdateTime = time.Now()
	return cs.repositoriesRepo.AddCategory(category)
}

// UpdateBrand 调用仓库层方法更新证书信息
func (cs *CategoryService) UpdateCategory(category *models.Categories) error {
	category.UpdateTime = time.Now()
	return cs.repositoriesRepo.UpdateCategory(category)
}

// DeleteBrand 调用仓库层方法标记证书为已删除
func (cs *CategoryService) DeleteCategory(id string) error {
	return cs.repositoriesRepo.DeleteCategory(id)
}
