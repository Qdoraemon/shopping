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

func (bs *BrandsService) GetBrandsByPage(page, pageSize int) ([]*models.Brands, error) {
	return bs.brandsRepo.GetBrandsByPage(page, pageSize)
}

func (bs *BrandsService) AddBrands(certificate *models.Brands) error {
	if certificate == nil {
		return errors.New("certificate cannot be nil")
	}
	certificate.CreateTime = time.Now()
	certificate.UpdateTime = time.Now()
	return bs.brandsRepo.AddBrands(certificate)
}

// UpdateBrands 调用仓库层方法更新证书信息
func (bs *BrandsService) UpdateBrands(certificate *models.Brands) error {
	certificate.UpdateTime = time.Now()
	return bs.brandsRepo.UpdateBrands(certificate)
}

// DeleteBrands 调用仓库层方法标记证书为已删除
func (bs *BrandsService) DeleteBrands(id interface{}) error {
	return bs.brandsRepo.DeleteBrands(id)
}

func (bs *BrandsService) GetAllBrands() ([]*models.Brands, error) {
	return bs.brandsRepo.GetAllBrands()
}
