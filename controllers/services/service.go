package services

import "github.com/57th-street/oza-gueser/models"

type AuthServicer interface {
	RegisterService(user models.User) (models.User, error)
	LoginService(user models.User) (models.User, error)
}
