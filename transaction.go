package xpay

type Transaction struct {
	StudentID int `json:"student_id" binding:"required"`
	Sum       int `json:"sum" binding:"required"`
}
