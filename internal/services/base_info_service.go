package services

import (
	"shopping/internal/models"
	"shopping/internal/repositories"

	"gorm.io/gorm"
)

type BaseInfoService struct {
	baseInfoRepo *repositories.UserRepository
}

func NewBaseInfoService(engine *gorm.DB) *BaseInfoService {
	return &BaseInfoService{baseInfoRepo: repositories.NewUserRepository(engine)}
}

func (us *BaseInfoService) GetUsers() ([]*models.User, error) {
	return us.baseInfoRepo.GetUsers()
}
