package main

import (
	"fmt"
	"net/http"
	"rest-api-golang/src/config"
	"rest-api-golang/src/controllers"
	"rest-api-golang/src/repositories"
	"rest-api-golang/src/services"
)

func main() {
	db, err := config.PostgresConnect()
	if err != nil {
		panic(err)
	}
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository, db)
	userController := controllers.NewUserController(userService)

	http.HandleFunc("/users", userController.Get)
	http.HandleFunc("/users/create", userController.Insert)
	fmt.Println("Web Starting")
	http.ListenAndServe(":8085", nil)
}
