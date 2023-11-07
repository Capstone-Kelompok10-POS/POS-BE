package web

type SuperAdminLoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type SuperAdminResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}