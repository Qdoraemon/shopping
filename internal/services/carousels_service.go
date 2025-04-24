package services

import (
	"shopping/internal/models"
	"shopping/internal/repositories"

	"gorm.io/gorm"
)

type CarouselsService struct {
	carouselsRepo *repositories.CarouselsRepository
}

func NewCarouselsService(engine *gorm.DB) *CarouselsService {
	return &CarouselsService{carouselsRepo: repositories.NewCarouselsRepository(engine)}
}

func (us *CarouselsService) GetAllCarousels() ([]*models.Carousel, error) {
	return us.carouselsRepo.GetAllCarousels()
}
