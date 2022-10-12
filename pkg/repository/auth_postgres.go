package repository

import (
	"github.com/jmoiron/sqlx"
	"xpay"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (a AuthPostgres) SignIn(teacher xpay.Teacher) (int, error) {
	return 0, nil
}

func (a AuthPostgres) SignUp(teacher xpay.Teacher) (int, error) {
	return 0, nil
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
