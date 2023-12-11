package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func MembershipDomainToMembershipPointResponse(response *domain.MembershipPoint) web.MembershipPointResponse {
	return web.MembershipPointResponse{
		ID:           response.ID,
		CreatedAt:    response.CreatedAt,
		MembershipID: response.MembershipID,
		Membership:   response.Membership,
		Point:        response.Point,
	}
}

func MembershipPointSchemaToMembershipPointDomain(response schema.MembershipPoint) *domain.MembershipPoint {
	return &domain.MembershipPoint{
		ID:           response.ID,
		CreatedAt:    response.CreatedAt,
		MembershipID: response.MembershipID,
		Point:        response.Point,
	}
}

func ConvertMembershipPointResponse(point []domain.MembershipPoint) []web.MembershipPointResponse {
	var results []web.MembershipPointResponse
	for _, point := range point {

		pointResponse := web.MembershipPointResponse{
			//ID:           p   oint.ID,
			CreatedAt:    point.CreatedAt,
			MembershipID: point.MembershipID,
			Membership:   point.Membership,
			Point:        point.Point,
		}
		results = append(results, pointResponse)
	}
	return results
}
