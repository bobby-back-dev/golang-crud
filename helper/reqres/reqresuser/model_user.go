package user

import (
	"github.com/bobby-back-dev/golang-crud/internal/app/user/models"
)

func RegistRespUser(user models.User) UserRequestRegistOrUpdate {
	return UserRequestRegistOrUpdate{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func GetAllDataUser(user models.User) UserResponseLogin {
	return UserResponseLogin{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func AppendUser(user []models.User) []UserResponseLogin {
	var userList []UserResponseLogin
	for _, v := range user {
		userList = append(userList, GetAllDataUser(v))
	}
	return userList
}
