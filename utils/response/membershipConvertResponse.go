package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func MembershipDomainToMembershipResponse(membership *domain.Membership) web.MembershipResponse {
	return web.MembershipResponse{
		ID:        membership.ID,
		CashierID: membership.CashierID,
		Name:      membership.Name,
		Point:     membership.Point,
		Telephone: membership.Telephone,
	}
}

func MembershipSchemaToMembershipDomain(membership *schema.Membership) *domain.Membership {
	return &domain.Membership{
		ID:        membership.ID,
		CashierID: membership.CashierID,
		Name:      membership.Name,
		Point:     membership.Point,
		Telephone: membership.Telephone,
	}
}

func ConvertMembershipResponse(memberships []domain.Membership) []web.MembershipResponse {
	var results []web.MembershipResponse
	for _, membership := range memberships {
		membershipResponse := web.MembershipResponse{
			ID:        membership.ID,
			CashierID: membership.CashierID,
			Name:      membership.Name,
			Point:     membership.Point,
			Telephone: membership.Telephone,
		}
		results = append(results, membershipResponse)
	}
	return results
}
