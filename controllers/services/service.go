package services

type AuthServicer interface {
	Register(email, password string) (int, error)
	Login(email, password string) (int, error)
}
