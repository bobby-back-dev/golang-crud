package reqresuser

import (
	"github.com/bobby-back-dev/golang-crud/internal/app/user/models"
)

type UserWebRes struct{}

func NewUserWebRes() *UserWebRes {
	return &UserWebRes{}
}

func (uws *UserWebRes) GetDataUser(user models.User) *UserResponseLogin {
	return &UserResponseLogin{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (uws *UserWebRes) GetAllDataUser(user models.User) *UserResponseLogin {
	return &UserResponseLogin{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func (uws *UserWebRes) AppendUser(user *[]models.User) *[]UserResponseLogin {
	var userList []UserResponseLogin
	for _, v := range *user {
		userList = append(userList, *uws.GetAllDataUser(v))
	}
	return &userList
}
