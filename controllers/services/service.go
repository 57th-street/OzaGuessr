package services

type AuthServicer interface {
	Register(email, password string) error
	Login(email, password string) error
}
