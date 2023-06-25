package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/57th-street/oza-gueser/controllers/services"
)

type AuthController struct {
	service services.AuthServicer
}

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewAuthController(s services.AuthServicer) *AuthController {
	return &AuthController{service: s}
}

func (c *AuthController) RegisterController(w http.ResponseWriter, req *http.Request) {
	var input UserInput
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}
	err := c.service.Register(input.Email, input.Password)
	if err != nil {
		http.Error(w, "fail to register\n", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *AuthController) LoginController(w http.ResponseWriter, req *http.Request) {
	var input UserInput
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}
	isAuthUser, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		http.Error(w, "fail to login\n", http.StatusInternalServerError)
		return
	}
	if isAuthUser {
		http.Error(w, "Authentication failed\n", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}
