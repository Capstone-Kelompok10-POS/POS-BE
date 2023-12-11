package web

import (
	"qbills/models/domain"
	"time"
)

type MembershipPointResponse struct {
	ID           uint              `json:"id"`
	CreatedAt    time.Time         `json:"createdAt"`
	MembershipID uint              `json:"membershipID"`
	Membership   domain.Membership `json:"membership"`
	Point        int               `json:"point"`
}

type MembershipPointCreateResponse struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	MembershipID uint      `json:"membershipID"`
	Point        int       `json:"point"`
}
