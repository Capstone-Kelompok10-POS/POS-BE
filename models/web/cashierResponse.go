package web

type CashierLoginResponse struct {
	Username string `json:"username"`
	Token string `json:"token"`
}

type CashierResponse struct {
	ID uint `json:"id"`
	AdminID uint `json:"adminID"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
}