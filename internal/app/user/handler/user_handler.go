package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bobby-back-dev/golang-crud/helper/reqres"
	"github.com/bobby-back-dev/golang-crud/helper/reqres/reqresuser"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userServices *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userServices,
	}
}

func (uh *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w.Write([]byte("Hello World")))
}

func (uh *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Println("Method Not Allowed")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Header.Set("Content-Type", "application/json")

	var users reqresuser.UserRequestRegistOrUpdate

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&users)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	log.Println(users)
	defer r.Body.Close()

	data, err := uh.userService.UserCreateService(users)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	dataUser := reqres.WebResp{
		Message: "User Created",
		Data:    data,
	}

	err = json.NewEncoder(w).Encode(&dataUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func (uh *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Method Not Allowed")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	r.Header.Set("Content-Type", "application/json")

	var users *reqresuser.UserRequestLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&users)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	data, err := uh.userService.LoginUser(users)
	if err != nil {
		http.Error(w, "email or password not valid", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dataLogin := reqres.WebResp{
		Message: "User Logged In",
		Data:    data,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&dataLogin); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (uh *UserHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("Method Not Allowed")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	data, err := uh.userService.GetAllUser()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dataUser := reqres.WebResp{
		Message: "User Found",
		Data:    data,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&dataUser); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (uh *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Println("Method Not Allowed")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data, err := uh.userService.GetUserByID(idInt)
	if err != nil {
		log.Println(data)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	dataUser := reqres.WebResp{
		Message: "User Found",
		Data:    data,
	}
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(&dataUser); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
