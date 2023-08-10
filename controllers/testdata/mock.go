package testdata

import (
	"github.com/57th-street/oza-gueser/models"
)

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) Register(email, password string) (int, error) {
	return 1, nil
}

func (s *serviceMock) Login(email, password string) (int, error) {
	return 1, nil
}

func (s *serviceMock) GetProfileService(userId int) (models.Profile, error) {
	return models.Profile{Name: ProfileTestData.Name, ImageUrl: ProfileTestData.ImageUrl, Intro: ProfileTestData.Intro}, nil
}

func (s *serviceMock) UpdateProfileService(userId int, profile models.Profile) error {
	return nil
}
