package xpay

type Transaction struct {
	Id        int `json:"id,omitempty"`
	TeacherID int `json:"-"`
	StudentID int `json:"student_id" binding:"required"`
	Sum       int `json:"sum" binding:"required"`
}
