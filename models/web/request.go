package web

type AdminCreateRequest struct {
	SuperAdminID uint `json:"superAdminID"`
	FullName     string `json:"fullname" validate:"required,min=1,max=255"`
	Username     string `json:"username" validate:"required,min=1"`
	Password     string `json:"password" validate:"required,min=8"`
}

type AdminLoginRequest struct {
	Username string `json:"username" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=8"`
}

type AdminUpdateRequest struct {
	SuperAdminID uint `json:"superAdminID"`
	FullName string `json:"fullname" validate:"required,min=1,max=255"`
	Username string `json:"username" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=8"`
}

//Cashier Request
type CashierCreateRequest struct {
	Admin_ID uint 	`json:"Admin_ID"`
	Fullname string `json:"fullname" validate:"required, min=1, max=255"`
	Username string `json:"username" validate:"required, min=1"`
	Password string `json:"password" validate:"required, min=8"`
}

type CashierLoginRequest struct {
	Username string `json:"username" validate:"required, min=1"`
	Password string `json:"password" validate:"required, min=8"`
}

type CashierUpdateRequest struct {
	Admin_ID uint 	`json:"Admin_ID"`
	Fullname string `json:"fullname" validate:"required, min=1, max=255"`
	Username string `json:"username" validate:"required, min=1"`
	Password string `json:"password" validate:"required, min=8"`
}

//Membership Request 
type MembershipCreateRequest struct {
	CashierID	uint	`json:"CashierID"`
	Name	    string 	`json:"name" validate:"required, min=1, max=255"`
	Telephone	string	`json:"telephone" validate:"required, min=1, max=15"`
}

type MembershipUpdateRequest struct {
	CashierID	uint	`json:"CashierID"`
	Name		string 	`json:"name" validate:"required, min=1, max=255"`
	Point       uint    `json:"point"`
	Telephone	string	`json:"telephone" validate:"required, min=1, max=15"`
}