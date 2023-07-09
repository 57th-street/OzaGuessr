package services

import (
	"database/sql"
	"errors"

	"github.com/57th-street/oza-gueser/apperrors"
	"github.com/57th-street/oza-gueser/repositories"
	"github.com/57th-street/oza-gueser/utils"
)

func (s *Service) Register(email, password string) error {
	exists, err := repositories.CheckUserExist(s.db, email)
	if err != nil {
		err = apperrors.CheckDataExistFailed.Wrap(err, "fail to check if email already exists")
		return err
	}
	if exists {
		err = apperrors.DataAlreadyExists.Wrap(err, "this email already exists")
		return err
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		err = apperrors.HashPasswordFailed.Wrap(err, "fail to hash password")
		return err
	}
	if err = repositories.CreateUser(s.db, email, hashedPassword); err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
	}
	return nil
}

func (s *Service) Login(email, password string) error {
	hashedPassword, err := repositories.GetUserPassword(s.db, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return err
	}
	if err = utils.CompareHashAndPassword(hashedPassword, password); err != nil {
		err = apperrors.ComparePasswordFailed.Wrap(err, "password does not match")
		return err
	}
	return nil
}
