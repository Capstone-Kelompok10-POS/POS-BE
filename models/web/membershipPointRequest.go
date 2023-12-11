package web

type MembershipPointCreate struct {
	MembershipID uint `json:"membershipID"`
	Point        uint `json:"point"`
}
