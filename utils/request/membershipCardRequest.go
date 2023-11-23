package request

import (
	"qbills/models/domain"
	"qbills/models/web"
)

func MembershipCardPrintRequestToMembershipDomain(request web.MembershipCardPrintRequest) *domain.Membership{
	return &domain.Membership{
		CashierID: request.CashierID,
		Name: request.Name,
		Phone_Number: request.Phone_Number,
	}
}
