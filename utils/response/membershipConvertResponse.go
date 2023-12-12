package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func MembershipDomainToMembershipResponse(membership *domain.Membership) web.MembershipResponse {
	return web.MembershipResponse{
		ID:          membership.ID,
		CashierID:   membership.CashierID,
		Name:        membership.Name,
		CodeMember:  membership.CodeMember,
		TotalPoint:  membership.TotalPoint,
		PhoneNumber: membership.PhoneNumber,
	}
}

func MembershipSchemaToMembershipDomain(membership *schema.Membership) *domain.Membership {
	return &domain.Membership{
		ID:          membership.ID,
		CashierID:   membership.CashierID,
		Name:        membership.Name,
		CodeMember:  membership.CodeMember,
		TotalPoint:  membership.TotalPoint,
		PhoneNumber: membership.PhoneNumber,
	}
}

func ConvertMembershipResponse(memberships []domain.Membership) []web.MembershipResponse {
	var results []web.MembershipResponse
	for _, membership := range memberships {
		membershipResponse := web.MembershipResponse{
			ID:          membership.ID,
			CashierID:   membership.CashierID,
			Name:        membership.Name,
			CodeMember:  membership.CodeMember,
			TotalPoint:  membership.TotalPoint,
			PhoneNumber: membership.PhoneNumber,
		}
		results = append(results, membershipResponse)
	}
	return results
}
