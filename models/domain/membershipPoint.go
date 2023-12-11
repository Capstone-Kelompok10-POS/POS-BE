package domain

import "time"

type MembershipPoint struct {
	ID           uint
	CreatedAt    time.Time
	MembershipID uint
	Membership   Membership
	Point        uint
}
