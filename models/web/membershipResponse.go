package web

// Membership Response
type MembershipResponse struct {
	ID          uint   `json:"id"`
	CashierID   uint   `json:"CashierID"`
	Name        string `json:"name"`
	Point       uint   `json:"point"`
	PhoneNumber string `json:"phoneNumber"`
}
