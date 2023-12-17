package web

import uuid "github.com/satori/go.uuid"

// Membership Response
type MembershipResponse struct {
	ID          uint      `json:"id"`
	CashierID   uint      `json:"cashierId"`
	Name        string    `json:"name"`
	CodeMember  uuid.UUID `json:"codeMember"`
	TotalPoint  uint      `json:"totalPoint"`
	PhoneNumber string    `json:"phoneNumber"`
}

type MembershipTransactionResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	CodeMember  uuid.UUID `json:"codeMember"`
	TotalPoint  uint      `json:"totalPoint"`
	PhoneNumber string    `json:"phoneNumber"`
}
