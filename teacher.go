package xpay

type Teacher struct {
	Id       int    `json:"-"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
