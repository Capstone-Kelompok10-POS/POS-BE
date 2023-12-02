package domain

import uuid "github.com/satori/go.uuid"

type Membership struct {
	ID           uint
	CashierID    uint
	Name         string
	CodeMember   uuid.UUID
<<<<<<< Updated upstream
	Point        uint
	Phone_Number string
=======
	Point        int
	PhoneNumber string
	Barcode      string
>>>>>>> Stashed changes
}
