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

func (r *UserRepository) GetUserById(userId int) (*models.User, error) {
	var users = new(models.User)
	err := r.engine.Table(models.User{}.TableName()).Where("id = ?", userId).First(users).Error
	return users, err
}

func (r *UserRepository) GetUserByUserName(username string) (*models.User, error) {
	var users = new(models.User)
	err := r.engine.Table(models.User{}.TableName()).Where("username = ?", username).First(users).Error
	return users, err
}

