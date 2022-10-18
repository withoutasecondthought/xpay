package xpay

type Student struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name" binding:"required"`
	TeacherId int    `json:"-"`
	Money     int    `json:"money"`
}
