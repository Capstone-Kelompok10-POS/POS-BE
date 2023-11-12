package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func MembershipCreateRequestToMembershipDomain(request web.MembershipCreateRequest) *domain.Membership {
	return &domain.Membership{
		CashierID: request.CashierID,
		Name:      request.Name,
		Telephone: request.Telephone,
	}
}

func MembershipUpdateRequestToMembershipDomain(request web.MembershipUpdateRequest) *domain.Membership {
	return &domain.Membership{
		CashierID: request.CashierID,
		Name:      request.Name,
		Telephone: request.Telephone,
	}
}

func MembershipDomainintoMembershipSchema(request domain.Membership) *schema.Membership {
	return &schema.Membership{
		CashierID: request.CashierID,
		Name:      request.Name,
		Telephone: request.Telephone,
	}
}