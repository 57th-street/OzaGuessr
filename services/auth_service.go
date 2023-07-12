package services

import (
	"database/sql"
	"errors"

	"github.com/57th-street/oza-gueser/apperrors"
	"github.com/57th-street/oza-gueser/repositories"
	"github.com/57th-street/oza-gueser/utils"
)

func (s *Service) Register(email, password string) (int, error) {
	exists, err := repositories.CheckUserExist(s.db, email)
	if err != nil {
		err = apperrors.CheckDataExistFailed.Wrap(err, "fail to check user")
		return 0, err
	}
	if exists {
		err = apperrors.DataAlreadyExists.Wrap(err, "this email already exists")
		return 0, err
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		err = apperrors.HashPasswordFailed.Wrap(err, "fail to hash password")
		return 0, err
	}
	result, err := repositories.CreateUser(s.db, email, hashedPassword)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return 0, err
	}
	userID, _ := result.LastInsertId()
	return int(userID), nil
}

func (s *Service) Login(email, password string) (int, error) {
	user, err := repositories.GetUser(s.db, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no data")
			return 0, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return 0, err
	}
	if err = utils.CompareHashAndPassword(user.HashedPassword, password); err != nil {
		err = apperrors.ComparePasswordFailed.Wrap(err, "password does not match")
		return 0, err
	}
	return user.ID, nil
}
