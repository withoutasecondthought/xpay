package xpay

type Teacher struct {
	Id       int    `json:"-" db:"id"`
	Email    string `json:"email" binding:"required" db:"email"`
	Password string `json:"password" binding:"required" db:"password"`
}
