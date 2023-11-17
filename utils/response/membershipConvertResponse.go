package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func MembershipDomainToMembershipResponse(membership *domain.Membership) web.MembershipResponse {
	return web.MembershipResponse{
		ID:           membership.ID,
		CashierID:    membership.CashierID,
		Name:         membership.Name,
		Point:        membership.Point,
		Phone_Number: membership.Phone_Number,
	}
}

func MembershipSchemaToMembershipDomain(membership *schema.Membership) *domain.Membership {
	return &domain.Membership{
		ID:           membership.ID,
		CashierID:    membership.CashierID,
		Name:         membership.Name,
		Point:        membership.Point,
		Phone_Number: membership.Phone_Number,
	}
}

func ConvertMembershipResponse(memberships []domain.Membership) []web.MembershipResponse {
	var results []web.MembershipResponse
	for _, membership := range memberships {
		membershipResponse := web.MembershipResponse{
			ID:           membership.ID,
			CashierID:    membership.CashierID,
			Name:         membership.Name,
			Point:        membership.Point,
			Phone_Number: membership.Phone_Number,
		}
		results = append(results, membershipResponse)
	}
	return results
}
