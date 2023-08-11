package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/57th-street/oza-gueser/apperrors"
	"github.com/57th-street/oza-gueser/controllers/services"
	"github.com/57th-street/oza-gueser/utils"
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
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	userID, err := c.service.Register(input.Email, input.Password)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	tokenString, err := utils.GenerateToken(userID)
	if err != nil {
		err = apperrors.GenerateTokenFailed.Wrap(err, "fail to generate token")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	resp := AuthResponse{AccessToken: tokenString}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (c *AuthController) LoginController(w http.ResponseWriter, req *http.Request) {
	var input UserInput
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	userID, err := c.service.Login(input.Email, input.Password)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	tokenString, err := utils.GenerateToken(userID)
	if err != nil {
		err = apperrors.GenerateTokenFailed.Wrap(err, "fail to generate token")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	resp := AuthResponse{AccessToken: tokenString}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
