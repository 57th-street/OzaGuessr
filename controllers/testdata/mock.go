package testdata

import "github.com/57th-street/oza-gueser/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) RegisterService(user models.User) (models.User, error) {
	return authUser, nil
}

func (s *serviceMock) LoginService(user models.User) (models.User, error) {
	return authUser, nil
}
