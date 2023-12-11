package schema

import "time"

type MembershipPoint struct {
	ID           uint       `gorm:"primaryKey"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	MembershipID uint       `gorm:"index" json:"membershipID"`
	Membership   Membership `gorm:"foreignKey:MembershipID" json:"membership"`
	Point        uint       `json:"point"`
}
