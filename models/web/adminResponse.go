package web

type AdminLoginResponse struct {
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type AdminResponse struct {
	ID           uint   `json:"id"`
	SuperAdminID uint   `json:"superAdminID"`
	FullName     string `json:"fullname"`
	Username     string `json:"username"`
}