package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func ConvertPointCreateRequestToConvertPointDomain(request web.ConvertPointRequest) *domain.ConvertPoint {
	return &domain.ConvertPoint{
		Point : request.Point,
		ValuePoint: request.ValuePoint,
	}
}

func ConvertPointUpdateRequestToConvertPointDomain(request web.ConvertPointRequest) *domain.ConvertPoint {
	return &domain.ConvertPoint{
		Point : request.Point,
		ValuePoint: request.ValuePoint,
	}
}

func ConvertPointDomainToConvertPointSchema(request domain.ConvertPoint) *schema.ConvertPoint {
	return &schema.ConvertPoint{
		Point : request.Point,
		ValuePoint: request.ValuePoint,
	}
}