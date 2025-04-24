package repositories

import (
	"shopping/internal/models"

	"gorm.io/gorm"
)

type CarouselsRepository struct {
	engine *gorm.DB
}

func NewCarouselsRepository(engine *gorm.DB) *CarouselsRepository {
	return &CarouselsRepository{engine: engine}
}

func (r *CarouselsRepository) GetAllCarousels() ([]*models.Carousel, error) {
	var carousels []*models.Carousel
	err := r.engine.Table(models.Carousel{}.TableName()).Find(&carousels).Error
	return carousels, err
}
