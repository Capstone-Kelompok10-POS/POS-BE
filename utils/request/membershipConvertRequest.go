package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func MembershipCreateRequestToMembershipDomain(request web.MembershipCreateRequest) *domain.Membership {
	return &domain.Membership{
		CashierID:   request.CashierID,
		Name:        request.Name,
		PhoneNumber: request.PhoneNumber,
	}
}

func MembershipUpdateRequestToMembershipDomain(request web.MembershipUpdateRequest) *domain.Membership {
	return &domain.Membership{
		Name:        request.Name,
		TotalPoint:  request.TotalPoint,
		PhoneNumber: request.PhoneNumber,
	}
}

func MembershipDomainintoMembershipSchema(request domain.Membership) *schema.Membership {
	return &schema.Membership{
		CashierID:   request.CashierID,
		Name:        request.Name,
		CodeMember:  request.CodeMember,
		TotalPoint:  request.TotalPoint,
		PhoneNumber: request.PhoneNumber,
	}
}
