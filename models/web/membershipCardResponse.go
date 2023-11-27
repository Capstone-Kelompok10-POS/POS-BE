package web

import (
	uuid "github.com/satori/go.uuid"
)

type MembershipCardResponse struct {
	Name           string    `json:"name"`
	CodeMember     uuid.UUID `json:"CodeMember"`
	AvailableDate string    `json:"availableDate"`
	Barcode        string    `json:"barcode"`
}
