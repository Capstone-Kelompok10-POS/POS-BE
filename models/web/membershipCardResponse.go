package web

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type MembershipCardResponse struct {
	Name           string    `json:"name"`
	Code_Member    uuid.UUID `json:"Code_Member"`
	Phone_Number   string    `json:"phone_number"`
	Available_Date time.Time `json:"available_date"`
	Barcode        string    `json:"barcode"`
}
