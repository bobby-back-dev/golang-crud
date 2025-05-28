package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/models"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/service"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userServices service.UserService) *UserHandler {
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

	var users *models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&users)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	data, err := uh.userService.UserCreateService(users)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	userDataResponse := map[string]interface{}{
		"id":    data.ID,
		"name":  data.Name,
		"email": data.Email,
	}
	response := map[string]interface{}{
		"message": "Success",
		"data":    userDataResponse,
	}

	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}
