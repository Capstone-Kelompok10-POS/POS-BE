package web

type CashierCreateRequest struct {
<<<<<<< Updated upstream
	AdminID  uint   `json:"adminID"`
	Fullname string `json:"fullname" validate:"required,min=1,max=255"`
	Username string `json:"username" validate:"required,min=1"`
=======
<<<<<<< Updated upstream
	AdminID uint 	`json:"AdminID"`
=======
	AdminID  uint   `json:"adminId"`
>>>>>>> Stashed changes
	Fullname string `json:"fullname" validate:"required,min=1,max=255"`
	Username string `json:"username" validate:"required,alphanum,min=1"`
>>>>>>> Stashed changes
	Password string `json:"password" validate:"required,min=8"`
}

type CashierLoginRequest struct {
	Username string `json:"username" validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=8"`
}

type CashierUpdateRequest struct {
<<<<<<< Updated upstream
	AdminID  uint   `json:"adminID"`
	Fullname string `json:"fullname" validate:"required,min=1,max=255"`
	Username string `json:"username" validate:"required,min=1"`
=======
<<<<<<< Updated upstream
	AdminID uint 	`json:"AdminID"`
=======
	AdminID  uint   `json:"adminId"`
>>>>>>> Stashed changes
	Fullname string `json:"fullname" validate:"required,min=1,max=255"`
	Username string `json:"username" validate:"required,alphanum,min=1"`
>>>>>>> Stashed changes
	Password string `json:"password" validate:"required,min=8"`
}