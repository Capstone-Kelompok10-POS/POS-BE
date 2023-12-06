package web

type MembershipCreateRequest struct {
	CashierID    uint   `json:"CashierId"`
	Name         string `json:"name" validate:"required,min=1,max=255"`
	PhoneNumber string `json:"phoneNumber" validate:"required,number,min=1,max=15"`
}


type MembershipUpdateRequest struct {
<<<<<<< Updated upstream
	CashierID   uint   `json:"cashierId"`
	Name        string `json:"name" validate:"required,min=1,max=255"`
	Point       int    `json:"point" validate:"numeric"`
	PhoneNumber string `json:"phoneNumber" validate:"required,number,min=1,max=15"`
}

=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
	CashierID   uint   `json:"CashierID"`
	Name        string `json:"name" validate:"required,min=1,max=255"`
	Point       uint   `json:"point"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=1,max=15"`
}
=======
<<<<<<< Updated upstream
>>>>>>> Stashed changes
	CashierID	uint	`json:"CashierID"`
	Name		string 	`json:"name" validate:"required,min=1,max=255"`
	Point       uint    `json:"point"`
	Telephone	string	`json:"telephone" validate:"required,min=1,max=15"`
<<<<<<< Updated upstream
}
=======
}
=======
	CashierID   uint   `json:"CashierID"`
	Name        string `json:"name" validate:"required,min=1,max=255"`
<<<<<<< Updated upstream
	Point       uint   `json:"point"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=1,max=15"`
=======
<<<<<<< Updated upstream
	Point       int    `json:"point" validate:"numeric"`
	PhoneNumber string `json:"phoneNumber" validate:"required,number,min=1,max=15"`
>>>>>>> Stashed changes
}

=======
	Point       uint   `json:"point"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=1,max=15"`
<<<<<<< Updated upstream
}
=======
}


>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
