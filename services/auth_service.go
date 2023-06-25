package services

import (
	"fmt"

	"github.com/57th-street/oza-gueser/repositories"
	"github.com/57th-street/oza-gueser/utils"
)

func (s *Service) Register(email, password string) error {
	exists, err := repositories.CheckUserExist(s.db, email)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("this email already registered")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return fmt.Errorf("fail to hash password")
	}
	if err := repositories.CreateUser(s.db, email, hashedPassword); err != nil {
		return fmt.Errorf("fail to record data")
	}
	return nil
}

func (s *Service) Login(email, password string) (bool, error) {
	hashedPassword, err := repositories.GetUserPassword(s.db, email)
	if err != nil {
		return false, fmt.Errorf("fail to get data")
	}
	if err = utils.CompareHashAndPassword(hashedPassword, password); err != nil {
		return false, fmt.Errorf("email or password is incorrect")
	}
	return true, nil
}
