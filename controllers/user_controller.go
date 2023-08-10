package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/57th-street/oza-gueser/apperrors"
	"github.com/57th-street/oza-gueser/controllers/services"
	"github.com/57th-street/oza-gueser/models"
	"github.com/57th-street/oza-gueser/utils"
)

type UserController struct {
	service services.UserServicer
}

func NewUserController(s services.UserServicer) *UserController {
	return &UserController{service: s}
}

func (c *UserController) GetProfileController(w http.ResponseWriter, req *http.Request) {
	userID := utils.GetUserID(req.Context())
	profile, err := c.service.GetProfileService(userID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(profile)
}

func (c *UserController) UpdateProfileController(w http.ResponseWriter, req *http.Request) {
	var profile models.Profile
	if err := json.NewDecoder(req.Body).Decode(&profile); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}
	userID := utils.GetUserID(req.Context())
	if err := c.service.UpdateProfileService(userID, profile); err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
