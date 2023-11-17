package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func MembershipCreateRequestToMembershipDomain(request web.MembershipCreateRequest) *domain.Membership {
	return &domain.Membership{
		CashierID:    request.CashierID,
		Name:         request.Name,
		Phone_Number: request.Phone_Number,
	}
}

func MembershipUpdateRequestToMembershipDomain(request web.MembershipUpdateRequest) *domain.Membership {
	return &domain.Membership{
		CashierID:    request.CashierID,
		Name:         request.Name,
		Point:        request.Point,
		Phone_Number: request.Phone_Number,
	}
}

func MembershipDomainintoMembershipSchema(request domain.Membership) *schema.Membership {
	return &schema.Membership{
		CashierID:    request.CashierID,
		Name:         request.Name,
		Point:        request.Point,
		Phone_Number: request.Phone_Number,
	}
}
