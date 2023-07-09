package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/57th-street/oza-gueser/apperrors"
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
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}
	if err := c.service.Register(input.Email, input.Password); err != nil {
		apperrors.ErrorHandler(w, req, err)
	}
	w.WriteHeader(http.StatusOK)
}

func (c *AuthController) LoginController(w http.ResponseWriter, req *http.Request) {
	var input UserInput
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}
	err := c.service.Login(input.Email, input.Password)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
	}
	w.WriteHeader(http.StatusOK)
}
