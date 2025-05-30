package services

import (
	"errors"
	"shopping/internal/models"
	"shopping/internal/repositories"
	"time"

	"gorm.io/gorm"
)

type BrandService struct {
	BrandRepo *repositories.BrandRepository
}

func NewBrandService(engine *gorm.DB) *BrandService {
	return &BrandService{BrandRepo: repositories.NewBrandRepository(engine)}
}

func (bs *BrandService) GetBrandByPage(page, pageSize int) ([]*models.Brand, error) {
	return bs.BrandRepo.GetBrandByPage(page, pageSize)
}

func (bs *BrandService) AddBrand(brand *models.Brand) error {
	if brand == nil {
		return errors.New("brand cannot be nil")
	}
	brand.CreateTime = time.Now()
	brand.UpdateTime = time.Now()
	return bs.BrandRepo.AddBrand(brand)
}

// UpdateBrand 调用仓库层方法更新证书信息
func (bs *BrandService) UpdateBrand(brand *models.Brand) error {
	brand.UpdateTime = time.Now()
	return bs.BrandRepo.UpdateBrand(brand)
}

// DeleteBrand 调用仓库层方法标记证书为已删除
func (bs *BrandService) DeleteBrand(id interface{}) error {
	return bs.BrandRepo.DeleteBrand(id)
}

func (bs *BrandService) GetAllBrand() ([]*models.Brand, error) {
	return bs.BrandRepo.GetAllBrand()
}
