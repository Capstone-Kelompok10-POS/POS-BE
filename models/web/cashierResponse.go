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
<<<<<<< Updated upstream
=======
	AdminID  uint   `json:"adminId"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
=======
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
}

type CashierTransactionResponse struct {
	ID       uint   `json:"id"`
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
	Fullname string `json:"fullname"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	Username string `json:"username"`
>>>>>>> Stashed changes
}