package response

import (
	"qbills/models/domain"
	"qbills/models/web"
	"time"
)

func MembershipCardDomainToMembershipCardResponse(membership *domain.Membership) web.MembershipCardResponse {
	availableDate := time.Now().AddDate(1, 0, 0) 
	return web.MembershipCardResponse{
		Name:           membership.Name,
		Code_Member:    membership.Code_Member,
		Phone_Number:   membership.Phone_Number,
		Available_Date: availableDate.Format("2006-01-02"),
		Barcode:        membership.Barcode,
	}
}

