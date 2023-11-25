package web

type MembershipCreateRequest struct {
	CashierID   uint   `json:"CashierID"`
	Name        string `json:"name" validate:"required,min=1,max=255"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=1,max=15"`
}

type MembershipUpdateRequest struct {
	CashierID   uint   `json:"CashierID"`
	Name        string `json:"name" validate:"required,min=1,max=255"`
	Point       uint   `json:"point"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=1,max=15"`
}
