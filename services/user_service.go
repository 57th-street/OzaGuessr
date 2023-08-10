package services

import (
	"github.com/57th-street/oza-gueser/apperrors"
	"github.com/57th-street/oza-gueser/models"
	"github.com/57th-street/oza-gueser/repositories"
)

func (s *Service) GetProfileService(userId int) (models.Profile, error) {
	profile, err := repositories.GetProfile(s.db, userId)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Profile{}, err
	}
	return profile, nil
}

func (s *Service) UpdateProfileService(userId int, profile models.Profile) error {
	if err := repositories.UpdateProfile(s.db, userId, profile); err != nil {
		return apperrors.UpdateDataFailed.Wrap(err, "fail to update data")
	}
	return nil
}
