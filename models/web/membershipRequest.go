package web

//Membership Request
type MembershipCreateRequest struct {
<<<<<<< Updated upstream
	CashierID    uint   `json:"CashierID"`
	Name         string `json:"name" validate:"required,min=1,max=255"`
	Phone_Number string `json:"phone_number" validate:"required,min=1,max=15"`
}

type MembershipUpdateRequest struct {
	CashierID    uint   `json:"CashierID"`
	Name         string `json:"name" validate:"required,min=1,max=255"`
	Point        uint   `json:"point"`
	Phone_Number string `json:"phone_number" validate:"required,min=1,max=15"`
}
=======
	CashierID   uint   `json:"cashierId"`
	Name        string `json:"name" validate:"required,min=1,max=255"`
	PhoneNumber string `json:"phoneNumber" validate:"required,number,min=1,max=15"`
}

type MembershipUpdateRequest struct {
	CashierID   uint   `json:"cashierId"`
	Name        string `json:"name" validate:"required,min=1,max=255"`
	Point       int    `json:"point" validate:"numeric"`
	PhoneNumber string `json:"phoneNumber" validate:"required,number,min=1,max=15"`
}
>>>>>>> Stashed changes
