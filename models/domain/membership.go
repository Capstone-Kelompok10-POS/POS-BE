package domain

import uuid "github.com/satori/go.uuid"

type Membership struct {
	ID           uint
	CashierID    uint
	Name         string
	CodeMember   uuid.UUID
	Point        int
	PhoneNumber string
	Barcode      string
}
