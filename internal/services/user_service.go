package services

import (
	"errors"
	"shopping/internal/models"
	"shopping/internal/repositories"
	"shopping/internal/utils"

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

/*
用于登录使用的方法
*/
func (us *UserService) LoginUser(username string, password string) (*models.User, error) {
	user, err := us.userRepo.GetUserByUserName(username)
	if err != nil {
		return nil, err
	}

	if utils.CheckPasswordHash(password, user.Password) == false {
		return nil, errors.New("invalid password")
	}
	return user, nil

}
