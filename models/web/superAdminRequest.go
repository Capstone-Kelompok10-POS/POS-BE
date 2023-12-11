package web

type SuperAdminLoginRequest struct {
	Username string `json:"username" validate:"required,alphanum,min=1"`
	Password string `json:"password" validate:"required,min=8"`
}
