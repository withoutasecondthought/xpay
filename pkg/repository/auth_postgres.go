package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"xpay"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func (a *AuthPostgres) SignIn(teacher xpay.Teacher) (int, error) {
	var id int

	query := fmt.Sprintf(`SELECT id FROM %s WHERE email=$1 AND password_hash=$2`, teachers)
	var err = a.db.Get(&id, query, teacher.Email, teacher.Password)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthPostgres) SignUp(teacher xpay.Teacher) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (email, password_hash) VALUES ($1, $2) RETURNING id`, teachers)
	var row = a.db.QueryRow(query, teacher.Email, teacher.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
