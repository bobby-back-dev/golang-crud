package service

import (
	"errors"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/models"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) UserCreateService(user *models.User) (*models.User, error) {

	data := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	if data.Name == "" || data.Email == "" || data.Password == "" {
		return nil, errors.New("invalid data")
	}

	repo, err := us.userRepo.CreateUser(data)
	if err != nil {
		return nil, err
	}

	return repo, nil
}
