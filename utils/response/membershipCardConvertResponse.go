package response

import (
	"qbills/models/domain"
	"qbills/models/web"
	"time"
)

func MembershipCardDomainToMembershipCardResponse(membership *domain.Membership) web.MembershipCardResponse {
	availableDate := time.Now().AddDate(1, 0, 0)
	return web.MembershipCardResponse{
		Name:          membership.Name,
		CodeMember:    membership.CodeMember,
		AvailableDate: availableDate.Format("2006-01-02"),
		Barcode:       membership.Barcode,
	}
}
