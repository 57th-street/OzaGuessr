package testdata

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
