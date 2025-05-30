package routes

import (
	"github.com/bobby-back-dev/golang-crud/internal/app/user/handler"
	"github.com/gorilla/mux"
)

func SetRouter(userHandler *handler.UserHandler) *mux.Router {
	r := mux.NewRouter()

	//r.HandleFunc("/get/user", userHandler.GetUserHandler).Methods("get")
	r.HandleFunc("/user/post", userHandler.CreateUserHandler).Methods("post")
	r.HandleFunc("/user/login", userHandler.LoginUser).Methods("post")
	r.HandleFunc("/get/all/user", userHandler.GetAllUser).Methods("get")
	r.HandleFunc("/user/get/{id}", userHandler.GetUserById).Methods("get")
	return r
}
