package service

import (
	"xpay"
	"xpay/pkg/repository"
)

type Auth interface {
	SignIn(teacher xpay.Teacher) (string, error)
	SignUp(teacher xpay.Teacher) (string, error)
}

type Student interface {
	GetStudent(studentId int) (xpay.Student, error)
	GetStudents(teacherId int) ([]*xpay.Student, error)
	NewStudent(student xpay.Student) error
	Transaction(transaction xpay.Transaction) error
	DeleteStudent(student xpay.Student) error
}

type Service struct {
	Auth
	Student
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth:    NewAuthService(repo.Auth),
		Student: NewStudentService(repo.Student),
	}
}
