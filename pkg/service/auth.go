package service

import (
	"crypto/sha1"
	"fmt"
	"net/mail"
	"xpay"
	"xpay/pkg/repository"
)

var salt = "nonigrep"

type AuthService struct {
	repo repository.Auth
}

func (a *AuthService) SignIn(teacher xpay.Teacher) (int, error) {
	_, err := mail.ParseAddress(teacher.Email)
	if err != nil {
		return 0, err
	}
	teacher.Password = hashPassword(teacher.Password)
	return a.repo.SignIn(teacher)
}

func (a *AuthService) SignUp(teacher xpay.Teacher) (int, error) {
	_, err := mail.ParseAddress(teacher.Email)
	if err != nil {
		return 0, err
	}
	teacher.Password = hashPassword(teacher.Password)
	return a.repo.SignUp(teacher)
}

func hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}
