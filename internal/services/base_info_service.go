package services

import (
	"shopping/internal/models"
	"shopping/internal/repositories"

	"gorm.io/gorm"
)

type BaseInfoService struct {
	baseInfoRepo *repositories.BaseInfoRepository
}

func NewBaseInfoService(engine *gorm.DB) *BaseInfoService {
	return &BaseInfoService{baseInfoRepo: repositories.NewBaseInfoRepository(engine)}
}

func (us *BaseInfoService) GetAllBasicInformation() ([]*models.BaseInformation, error) {
	return us.baseInfoRepo.GetAllBasicInformations()
}
