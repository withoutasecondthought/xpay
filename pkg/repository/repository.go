package repository

import (
	"github.com/jmoiron/sqlx"
	"xpay"
)

const (
	teachers     = "teachers"
	students     = "students"
	transactions = "transactions"
)

type Auth interface {
	SignIn(teacher xpay.Teacher) (int, error)
	SignUp(teacher xpay.Teacher) (int, error)
}

type Student interface {
	GetStudent(student xpay.Student) (xpay.Student, error)
	GetStudents(teacherId int) ([]*Student, error)
	NewStudent(student xpay.Student) error
	DeleteStudent(student xpay.Student) error
}

type Repository struct {
	Auth
	Student
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
	}
}
