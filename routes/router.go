package routes

import (
	"github.com/bobby-back-dev/golang-crud/internal/app/user/handler"
	"github.com/gorilla/mux"
)

func SetRouter(userHandler *handler.UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/get/user", userHandler.GetUserHandler)
	r.HandleFunc("/user/post", userHandler.CreateUserHandler)
	return r
}
