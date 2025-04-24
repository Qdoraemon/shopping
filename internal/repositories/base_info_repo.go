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

// TODO 测试用例
func (r *BaseInfoRepository) GetAllBasicInformations() ([]*models.BaseInformation, error) {
	var users []*models.BaseInformation
	err := r.engine.Table(models.BaseInformation{}.TableName()).Find(&users).Error
	return users, err
}
