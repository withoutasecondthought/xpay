package repository

import "github.com/jmoiron/sqlx"

type AuthPostgres interface {

}

type StudentPostgres interface {

}

type Repository struct {
	AuthPostgres
	StudentPostgres
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthPostgres:,
		StudentPostgres:,
	}
}