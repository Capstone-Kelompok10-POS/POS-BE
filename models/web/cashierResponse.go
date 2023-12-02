package web

type CashierLoginResponse struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type CashierResponse struct {
	ID       uint   `json:"id"`
<<<<<<< Updated upstream
	AdminID  uint   `json:"adminID"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
=======
	AdminID  uint   `json:"adminId"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
}

type CashierTransactionResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
>>>>>>> Stashed changes
}