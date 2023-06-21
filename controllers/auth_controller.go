package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/57th-street/oza-gueser/controllers/services"
	"github.com/57th-street/oza-gueser/models"
)

type AuthController struct {
	service services.AuthServicer
}

func NewAuthController(s services.AuthServicer) *AuthController {
	return &AuthController{service: s}
}

func (c *AuthController) RegisterController(w http.ResponseWriter, req *http.Request) {
	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}
	user, err := c.service.RegisterService(user)
	if err != nil {
		http.Error(w, "fail to register\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (c *AuthController) LoginController(w http.ResponseWriter, req *http.Request) {
	var user models.User
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}
	user, err := c.service.LoginService(user)
	if err != nil {
		http.Error(w, "fail to register\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
