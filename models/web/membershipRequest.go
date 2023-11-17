package web

//Membership Request
type MembershipCreateRequest struct {
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
