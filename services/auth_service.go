package services

import (
	"fmt"

	"github.com/57th-street/oza-gueser/models"
	"github.com/57th-street/oza-gueser/repositories"
)

func (s *Service) RegisterService(user models.User) (models.User, error) {
	exists, err := repositories.CheckUserExist(s.db, user.Email)
	if err != nil {
		return models.User{}, err
	}
	if exists {
		return models.User{}, fmt.Errorf("this email already registered")
	}
	registeredUser, err := repositories.Register(s.db, user)
	if err != nil {
		return models.User{}, err
	}
	return registeredUser, nil
}

func (s *Service) LoginService(user models.User) (models.User, error) {
	loggedInUser, err := repositories.Login(s.db, user)
	if err != nil {
		return models.User{}, err
	}
	return loggedInUser, nil
}
