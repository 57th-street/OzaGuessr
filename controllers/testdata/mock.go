package testdata

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) Register(email, password string) error {
	return nil
}

func (s *serviceMock) Login(email, password string) error {
	return nil
}
