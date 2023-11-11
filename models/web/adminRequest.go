package web

type AdminCreateRequest struct {
	SuperAdminID uint   `json:"superAdminID"`
	FullName     string `json:"fullname" validate:"required,min=1,max=255"`
	Username     string `json:"username" validate:"required,min=1"`
	Password     string `json:"password" validate:"required,min=8"`
}

type AdminLoginRequest struct {
	Username string `json:"username" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=8"`
}

type AdminUpdateRequest struct {
	SuperAdminID uint   `json:"superAdminID"`
	FullName     string `json:"fullname" validate:"required,min=1,max=255"`
	Username     string `json:"username" validate:"required,min=1"`
	Password     string `json:"password" validate:"required,min=8"`
}
