package xpay

type Student struct {
	Id        int    `json:"id,omitempty" db:"id"`
	Name      string `json:"name" binding:"required" db:"name"`
	TeacherId int    `json:"-" db:"teacher_id"`
	Money     int    `json:"money" db:"sum"`
}
