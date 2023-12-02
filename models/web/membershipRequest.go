package web

type MembershipCreateRequest struct {
	CashierID    uint   `json:"CashierId"`
	Name         string `json:"name" validate:"required,min=1,max=255"`
	PhoneNumber string `json:"phoneNumber" validate:"required,number,min=1,max=15"`
}


type MembershipUpdateRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=255"`
	Point       int    `json:"point" validate:"numeric"`
	PhoneNumber string `json:"phoneNumber" validate:"required,number,min=1,max=15"`
}

