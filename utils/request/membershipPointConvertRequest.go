package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func MembershipPointCreateToMembershipPointDomain(request web.MembershipPointCreate) *domain.MembershipPoint {
	return &domain.MembershipPoint{
		MembershipID: request.MembershipID,
		Point:        request.Point,
	}
}

func MembershipPointDomainToMembershipPointSchema(request *domain.MembershipPoint) schema.MembershipPoint {
	return schema.MembershipPoint{
		ID:           request.ID,
		CreatedAt:    request.CreatedAt,
		MembershipID: request.MembershipID,
		Point:        request.Point,
	}
}
