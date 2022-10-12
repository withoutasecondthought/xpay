package xpay

type Student struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name" binding:"required"`
	TeacherID int    `json:"teacher_id" binding:"required"`
	Money     int    `json:"money,omitempty"`
}
