package service

import (
	"errors"
	"github.com/bobby-back-dev/golang-crud/helper/crypto"
	"github.com/bobby-back-dev/golang-crud/helper/reqres/reqresuser"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/models"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/repository"
	"log"
)

type UserService struct {
	userRepo *repository.UserRepository
	hash     *crypto.Hash
	resp     *reqresuser.UserWebRes
}

func NewUserService(userRepo *repository.UserRepository, hash *crypto.Hash, resp *reqresuser.UserWebRes) *UserService {
	return &UserService{
		userRepo: userRepo,
		hash:     hash,
		resp:     resp,
	}
}

func (us *UserService) UserCreateService(user reqresuser.UserRequestRegistOrUpdate) (*reqresuser.UserResponseLogin, error) {

	data := &models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	if data.Name == "" || data.Email == "" || data.Password == "" {
		return nil, errors.New("invalid data")
	}

	dataPassword, err := us.hash.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}
	log.Println("password plaintext : ", data.Password)

	data.Password = dataPassword

	log.Println("password hash : ", data.Password)

	repo, err := us.userRepo.CreateUser(data)
	if err != nil {
		return nil, err
	}
	return us.resp.GetDataUser(*repo), nil
}

func (us *UserService) LoginUser(user *reqresuser.UserRequestLogin) (*reqresuser.UserResponseLogin, error) {

	data := &models.User{
		Email:    user.Email,
		Password: user.Password,
	}

	if data.Email == "" || data.Password == "" {
		return nil, errors.New("invalid data")
	}
	dataPassword, err := us.userRepo.Login(data.Email, data.Password)
	if err != nil {
		return nil, err
	}

	log.Println("password input : ", data.Password)
	log.Println("password hash : ", dataPassword)

	return us.resp.GetDataUser(*dataPassword), nil
}

func (us *UserService) GetAllUser() (*[]reqresuser.UserResponseLogin, error) {
	//perbaiki error internal server error
	data, err := us.userRepo.GetAll()
	if err != nil {
		return nil, errors.New("user not found")
	}

	return us.resp.AppendUser(data), nil
}

func (us *UserService) GetUserByID(id int) (*reqresuser.UserResponseLogin, error) {

	data, err := us.userRepo.GetUserByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return us.resp.GetDataUser(*data), nil
}
