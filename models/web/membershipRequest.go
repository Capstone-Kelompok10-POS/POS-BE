package web

<<<<<<< Updated upstream
//Membership Request
=======
<<<<<<< Updated upstream
//Membership Request 
>>>>>>> Stashed changes
type MembershipCreateRequest struct {
	CashierID    uint   `json:"CashierID"`
	Name         string `json:"name" validate:"required,min=1,max=255"`
	Phone_Number string `json:"phone_number" validate:"required,min=1,max=15"`
}

type MembershipUpdateRequest struct {
<<<<<<< Updated upstream
	CashierID    uint   `json:"CashierID"`
=======
	CashierID	uint	`json:"CashierID"`
	Name		string 	`json:"name" validate:"required,min=1,max=255"`
	Point       uint    `json:"point"`
	Telephone	string	`json:"telephone" validate:"required,min=1,max=15"`
}
=======
//Membership Request
type MembershipCreateRequest struct {
	CashierID    uint   `json:"cashierID"`
	Name         string `json:"name" validate:"required,min=1,max=255"`
	Phone_Number string `json:"phone_number" validate:"required,min=1,max=15"`
}

type MembershipUpdateRequest struct {
	CashierID    uint   `json:"cashierID"`
>>>>>>> Stashed changes
	Name         string `json:"name" validate:"required,min=1,max=255"`
	Point        uint   `json:"point"`
	Phone_Number string `json:"phone_number" validate:"required,min=1,max=15"`
}
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
