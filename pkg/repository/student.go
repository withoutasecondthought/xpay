package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"xpay"
)

type StudentPostgres struct {
	db *sqlx.DB
}

func (s *StudentPostgres) GetStudent(studentId int) (xpay.Student, error) {
	var student xpay.Student

	query := fmt.Sprintf(`SELECT %s.id, %s.name, %s.teacher_id, SUM(%s.sum)
								FROM %s
								LEFT JOIN %s
								ON %s.id = $1
								GROUP BY %s.name, %s.id, %s.teacher_id`,
		students, students, students, transactions, students, transactions, students, students, students, students)

	err := s.db.Get(&student, query, studentId)
	if err != nil {
		return xpay.Student{}, err
	}

	return student, nil
}

func (s *StudentPostgres) GetStudents(teacherId int) ([]xpay.Student, error) {
	var studentsArr []xpay.Student

	query := fmt.Sprintf(`SELECT %s.id, %s.name, %s.teacher_id, SUM(%s.sum)
								FROM %s 
								INNER JOIN %s 
								ON %s.teacher_id = $1 
								GROUP BY %s.id, %s.name, %s.teacher_id`,
		students, students, students, transactions, students, transactions, students, students, students, students)

	err := s.db.Select(&studentsArr, query, teacherId)
	if err != nil {
		return nil, err
	}

	return studentsArr, nil
}

func (s *StudentPostgres) NewStudent(student xpay.Student) error {
	query := fmt.Sprintf(`INSERT INTO %s (teacher_id, name) VALUES ($1, $2)`, students)

	err := s.db.QueryRow(query, student.TeacherId, student.Name).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentPostgres) Transaction(transaction xpay.Transaction) error {
	query := fmt.Sprintf(`INSERT INTO %s (teacher_id, student_id, sum) VALUES ($1, $2, $3)`, transactions)

	err := s.db.QueryRow(query, transaction.TeacherId, transaction.StudentId, transaction.Sum).Err()
	if err != nil {
		return err
	}

	return nil
}

func (s *StudentPostgres) DeleteStudent(student xpay.Student) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE teacher_id=$1 AND id=$2`, students)

	err := s.db.QueryRow(query, student.TeacherId, student.Id).Err()
	if err != nil {
		return err
	}

	return nil
}

func NewStudentPostgres(db *sqlx.DB) *StudentPostgres {
	return &StudentPostgres{db: db}
}
