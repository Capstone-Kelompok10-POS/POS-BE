package web

// Membership Response
type MembershipResponse struct {
	ID           uint   `json:"id"`
	CashierID    uint   `json:"cashierID"`
	Name         string `json:"name"`
	Point        uint   `json:"point"`
	Phone_Number string `json:"phone_number"`
}
