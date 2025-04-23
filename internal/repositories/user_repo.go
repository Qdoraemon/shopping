package repositories

import (
	"shopping/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	engine *gorm.DB
}

func NewUserRepository(engine *gorm.DB) *UserRepository {
	return &UserRepository{engine: engine}
}

// GetUsers 获取所有用户
func (r *UserRepository) GetUsers() ([]*models.User, error) {
	var users []*models.User
	err := r.engine.Table(models.User{}.TableName()).Find(&users).Error
	return users, err
}
