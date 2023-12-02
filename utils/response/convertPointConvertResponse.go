package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)


func ConvertPointSchemaToConvertPointDomain(convertPoint *schema.ConvertPoint) *domain.ConvertPoint {
	return &domain.ConvertPoint{
		ID:             convertPoint.ID,
		Point: convertPoint.Point,
		ValuePoint: convertPoint.ValuePoint,
	}
}

func ConvertPointDomainToConvertPointResponse(convertPoint *domain.ConvertPoint) web.ConvertPointResponse {
	return web.ConvertPointResponse{
		ID:             convertPoint.ID,
		Point: convertPoint.Point,
		ValuePoint: convertPoint.ValuePoint,
	}
}

func ConvertCPointResponse(convertPoints []domain.ConvertPoint) []web.ConvertPointResponse {
	var results []web.ConvertPointResponse
	for _, convertPoint := range convertPoints {
		convertPointResponse := web.ConvertPointResponse{
			ID:             convertPoint.ID,
			Point: convertPoint.Point,
			ValuePoint: convertPoint.ValuePoint,
		}
		results = append(results, convertPointResponse)
	}
	return results
}