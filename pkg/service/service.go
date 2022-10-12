package service

import "xpay/pkg/repository"

type Auth interface {
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
