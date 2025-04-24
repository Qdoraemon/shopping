package repositories

import (
	"shopping/internal/models"

	"gorm.io/gorm"
)

type BaseInfoRepository struct {
	engine *gorm.DB
}

func NewBaseInfoRepository(engine *gorm.DB) *BaseInfoRepository {
	return &BaseInfoRepository{engine: engine}
}

func (r *BaseInfoRepository) GetAllBasicInformations() ([]*models.BaseInformation, error) {
	var baseInfo []*models.BaseInformation
	err := r.engine.Table(models.BaseInformation{}.TableName()).Find(&baseInfo).Error
	return baseInfo, err
}
