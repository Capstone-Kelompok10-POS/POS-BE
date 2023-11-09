package web

type AdminLoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type AdminResponse struct {
	ID           uint   `json:"id"`
	SuperAdminID uint   `json:"superAdminID"`
	FullName     string `json:"fullname"`
	Username     string `json:"username"`
}
