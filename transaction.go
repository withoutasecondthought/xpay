package xpay

type Transaction struct {
	TeacherId int `json:"-" db:"teacher_id"`
	StudentId int `json:"student_id" binding:"required" db:"student_id"`
	Sum       int `json:"sum" binding:"required" db:"sum"`
}
