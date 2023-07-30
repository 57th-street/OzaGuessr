package services

import "github.com/57th-street/oza-gueser/models"

type AuthServicer interface {
	Register(email, password string) (int, error)
	Login(email, password string) (int, error)
}

type UserServicer interface {
	GetProfileService(userId int) (models.Profile, error)
	UpdateProfileService(userId int, profile models.Profile) error
}
