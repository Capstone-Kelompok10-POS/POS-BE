package web

type AdminLoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type AdminResponse struct {
	ID       uint   `json:"id"`
	SuperAdminID uint `json:"superAdminID"`
	FullName string `json:"fullname"`
	Username string `json:"username"`
}

type CashierLoginResponse struct {
	Username string `json:"username"`
	Token string `json:"token"`
}

type CashierResponse struct {
	ID uint `json:"id"`
	Admin_ID uint `json:"admin_ID"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
}