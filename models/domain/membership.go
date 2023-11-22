package domain

import uuid "github.com/satori/go.uuid"

type Membership struct {
	ID           uint
	CashierID    uint
	Name         string
	CodeMember   uuid.UUID
	Point        uint
	Phone_Number string
}
