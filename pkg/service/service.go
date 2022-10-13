package service

import (
	"xpay"
	"xpay/pkg/repository"
)

type Auth interface {
	SignIn(teacher xpay.Teacher) (int, error)
	SignUp(teacher xpay.Teacher) (int, error)
}

type Student interface {
}

type Service struct {
	Auth
	Student
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo.Auth),
	}
}
