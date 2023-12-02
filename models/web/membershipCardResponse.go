package web

import (
	uuid "github.com/satori/go.uuid"
)

type MembershipCardResponse struct {
	Name          string    `json:"name"`
<<<<<<< Updated upstream
	CodeMember    uuid.UUID `json:"CodeMember"`
=======
	CodeMember    uuid.UUID `json:"codeMember"`
>>>>>>> Stashed changes
	AvailableDate string    `json:"availableDate"`
	Barcode       string    `json:"barcode"`
}
