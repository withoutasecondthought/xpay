package service

import (
	"xpay"
	"xpay/pkg/repository"
)

type StudentService struct {
	repos repository.Student
}

func (s *StudentService) GetStudent(studentId int) (xpay.Student, error) {
	return s.repos.GetStudent(studentId)
}

func (s *StudentService) GetStudents(teacherId int) ([]xpay.Student, error) {
	return s.repos.GetStudents(teacherId)

}

func (s *StudentService) NewStudent(student xpay.Student) error {
	return s.repos.NewStudent(student)

}

func (s *StudentService) Transaction(transaction xpay.Transaction) error {
	return s.repos.Transaction(transaction)

}

func (s *StudentService) DeleteStudent(student xpay.Student) error {
	return s.repos.DeleteStudent(student)
}

func NewStudentService(repos repository.Student) *StudentService {
	return &StudentService{repos: repos}
}
