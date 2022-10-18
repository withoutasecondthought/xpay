package xpay

type Transaction struct {
	TeacherId int `json:"-"`
	StudentId int `json:"student_id" binding:"required"`
	Sum       int `json:"sum" binding:"required"`
}
