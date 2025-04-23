package services

import (
	"shopping/internal/models"
	"shopping/internal/repositories"

	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(engine *gorm.DB) *UserService {
	return &UserService{userRepo: repositories.NewUserRepository(engine)}
}

func (us *UserService) GetUsers() ([]*models.User, error) {
	return us.userRepo.GetUsers()
}
