package services

import (
	"errors"
	"shopping/internal/middleware"
	"shopping/internal/models"
	"shopping/internal/repositories"
	"shopping/internal/utils"
	"strings"

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

	if utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}
	return user, nil

}

func (us *UserService) GetInfo(token string) (*middleware.MyClaims, error) {
	if condition := strings.Contains(token, "Bearer"); !condition {
		return &middleware.MyClaims{}, errors.New("token格式錯誤")
	}

	// 去掉前缀Bearer
	tokenString := strings.TrimPrefix(token, "Bearer ")
	// 解析token
	Claims, err := middleware.ParseToken(tokenString)
	if err != nil {
		return &middleware.MyClaims{}, err
	}

	return Claims, err
}
