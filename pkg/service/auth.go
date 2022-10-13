package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/mail"
	"os"
	"time"
	"xpay"
	"xpay/pkg/repository"
)

var salt = "nonigrep"

type AuthService struct {
	repo repository.Auth
}

type CustomJWTClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

func (a *AuthService) SignIn(teacher xpay.Teacher) (string, error) {
	_, err := mail.ParseAddress(teacher.Email)
	if err != nil {
		return "", err
	}
	teacher.Password = hashPassword(teacher.Password)
	id, err := a.repo.SignIn(teacher)
	if err != nil {
		return "", err
	}

	return NewJWT(id)
}

func (a *AuthService) SignUp(teacher xpay.Teacher) (string, error) {
	_, err := mail.ParseAddress(teacher.Email)
	if err != nil {
		return "", err
	}
	teacher.Password = hashPassword(teacher.Password)
	id, err := a.repo.SignUp(teacher)
	if err != nil {
		return "", err
	}

	return NewJWT(id)
}

func NewJWT(id int) (string, error) {
	var claims = CustomJWTClaims{
		id,
		jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

func hashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}
