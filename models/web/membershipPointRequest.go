package web

type MembershipPointCreate struct {
	MembershipID uint `json:"membershipID"`
	Point        int  `json:"point"`
}
