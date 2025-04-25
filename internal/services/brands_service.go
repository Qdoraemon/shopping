package services

import (
	"errors"
	"shopping/internal/models"
	"shopping/internal/repositories"
	"time"

	"gorm.io/gorm"
)

type BrandsService struct {
	brandsRepo *repositories.BrandsRepository
}

func NewBrandsService(engine *gorm.DB) *BrandsService {
	return &BrandsService{brandsRepo: repositories.NewBrandsRepository(engine)}
}

func (us *BrandsService) GetBrandsByPage(page, pageSize int) ([]*models.Brands, error) {
	return us.brandsRepo.GetBrandsByPage(page, pageSize)
}

func (us *BrandsService) AddBrands(certificate *models.Brands) error {
	if certificate == nil {
		return errors.New("certificate cannot be nil")
	}
	certificate.CreateTime = time.Now()
	certificate.UpdateTime = time.Now()
	return us.brandsRepo.AddBrands(certificate)
}

// UpdateBrands 调用仓库层方法更新证书信息
func (us *BrandsService) UpdateBrands(certificate *models.Brands) error {
	certificate.UpdateTime = time.Now()
	return us.brandsRepo.UpdateBrands(certificate)
}

// DeleteBrands 调用仓库层方法标记证书为已删除
func (us *BrandsService) DeleteBrands(id interface{}) error {
	return us.brandsRepo.DeleteBrands(id)
}
