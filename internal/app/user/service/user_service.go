package service

import (
	"context"
	"errors"
	"github.com/bobby-back-dev/golang-crud/helper/crypto"
	"github.com/bobby-back-dev/golang-crud/helper/reqres/reqresuser"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/models"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/repository"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type UserServices interface {
	UserCreateService(ctx context.Context, user reqresuser.UserRequestRegisOrUpdate) (*reqresuser.UserResponseLogin, error)
	LoginUser(ctx context.Context, user *reqresuser.UserRequestLogin) (*reqresuser.UserResponseLogin, error)
}

type userService struct {
	userRepo repository.UserRepo
	hash     *crypto.Hash
	resp     *reqresuser.UserWebRes
}

func NewUserService(userRepo repository.UserRepo, hash *crypto.Hash, resp *reqresuser.UserWebRes) UserServices {
	return &userService{
		userRepo: userRepo,
		hash:     hash,
		resp:     resp,
	}
}

var stringToken = []byte("datatahasiauntukjwt")

func (us *userService) UserCreateService(ctx context.Context, user reqresuser.UserRequestRegisOrUpdate) (*reqresuser.UserResponseLogin, error) {

	data := &models.User{
		Username:     user.Username,
		PasswordHash: user.PasswordHash,
		DisplayName:  user.DisplayName,
	}
	if data.Username == "" || data.PasswordHash == "" || data.DisplayName == "" {
		return nil, errors.New("username and password are required")
	}

	dataPassword, err := us.hash.HashPassword(data.PasswordHash)
	if err != nil {
		return nil, err
	}
	log.Println("password plaintext : ", data.PasswordHash)

	data.PasswordHash = dataPassword

	log.Println("password hash : ", data.PasswordHash)

	repo, err := us.userRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	return us.resp.GetDataUser(*repo), nil
}

func (us *userService) LoginUser(ctx context.Context, user *reqresuser.UserRequestLogin) (*reqresuser.UserResponseLogin, error) {

	if user.Username == "" || user.PasswordHash == "" {
		return nil, errors.New("username is required")
	}
	data, err := us.userRepo.Login(ctx, user.Username)
	if err != nil {
		return nil, errors.New("username or password is wrong")
	}

	passworIsValid := us.hash.CheckPasswordHash(user.PasswordHash, data.PasswordHash)
	if !passworIsValid {
		return nil, errors.New("password is invalid")
	}

	claim := jwt.MapClaims{
		"id":          data.ID,
		"username":    data.Username,
		"displayname": data.DisplayName,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenData, err := token.SignedString(stringToken)
	if err != nil {
		return nil, errors.New("token is invalid")
	}

	log.Println("token : ", tokenData)

	log.Println("password input : ", data.PasswordHash)
	log.Println("password hash : ", passworIsValid)

	return us.resp.GetDataUser(*data), nil
}

//
//func (us *UserService) GetAllUser() (*[]reqresuser.UserResponseLogin, error) {
//	//perbaiki error internal server error
//	data, err := us.userRepo.GetAll()
//	if err != nil {
//		return nil, errors.New("user not found")
//	}
//
//	return us.resp.AppendUser(data), nil
//}
//
//func (us *UserService) GetUserByID(id int) (*reqresuser.UserResponseLogin, error) {
//
//	data, err := us.userRepo.GetUserByID(id)
//	if err != nil {
//		return nil, errors.New("user not found")
//	}
//
//	return us.resp.GetDataUser(*data), nil
//}
