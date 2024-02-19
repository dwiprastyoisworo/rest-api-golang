package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-golang/src/services"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (c UserController) Get(x http.ResponseWriter, r *http.Request) {
	x.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		users, err := c.UserService.Get(r.Context())
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
		}
		response := Response{
			Code: http.StatusOK,
			Data: users,
		}
		responseJson, err := json.Marshal(response)
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
		}
		x.Write(responseJson)
		return
	}
	http.Error(x, "", http.StatusBadRequest)
}

type Response struct {
	Code  int
	Data  interface{}
	Error string
}
