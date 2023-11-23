package web

import uuid "github.com/satori/go.uuid"

type MembershipCardResponse struct {
	Name         string    `json:"name"`
	CodeMember   uuid.UUID `json:"CodeMember"`
	Phone_Number string    `json:"phone_number"`
}