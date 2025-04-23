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
func (r *BaseInfoRepository) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := r.engine.Table(models.User{}.TableName()).Find(&users).Error
	return users, err
}
