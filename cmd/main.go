package main

import (
	"fmt"
	"github.com/bobby-back-dev/golang-crud/config/platform/database"
	"github.com/bobby-back-dev/golang-crud/helper/crypto"
	"github.com/bobby-back-dev/golang-crud/helper/godo"
	"github.com/bobby-back-dev/golang-crud/helper/reqres/reqresuser"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/handler"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/repository"
	"github.com/bobby-back-dev/golang-crud/internal/app/user/service"
	"github.com/bobby-back-dev/golang-crud/routes"
	"log"
	"net/http"
	"time"
)

func main() {
	if err := godo.LoadEnv(); err != nil {
		fmt.Printf("env gagal di muat")
	}

	if err := database.ConnectToDb(); err != nil {
		log.Fatal("Error koneksi ke db: ", err)
	}
	dbPool := database.GetPool()

	defer database.ClosePool()

	hash := crypto.NewHash()
	userRes := reqresuser.NewUserWebRes()

	userRepository := repository.NewUserRepository(dbPool, hash)
	userService := service.NewUserService(userRepository, hash, userRes)
	userHandler := handler.NewUserHandler(userService)
	handle := routes.SetRouter(userHandler)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        handle,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
