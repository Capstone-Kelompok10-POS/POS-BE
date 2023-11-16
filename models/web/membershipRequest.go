package web

//Membership Request 
type MembershipCreateRequest struct {
	CashierID	uint	`json:"CashierID"`
	Name	    string 	`json:"name" validate:"required,min=1,max=255"`
	Telephone	string	`json:"telephone" validate:"required,min=1,max=15"`
}

type MembershipUpdateRequest struct {
	CashierID	uint	`json:"CashierID"`
	Name		string 	`json:"name" validate:"required,min=1,max=255"`
	Point       uint    `json:"point"`
	Telephone	string	`json:"telephone" validate:"required,min=1,max=15"`
}