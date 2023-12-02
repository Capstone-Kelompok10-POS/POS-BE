package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func PaymentTypeDomainToPaymentTypeRespone(response *domain.PaymentType) web.PaymentTypeResponse {
	return web.PaymentTypeResponse{
		ID:       response.ID,
		TypeName: response.TypeName,
	}
}

func PaymentTypeSchemaToPaymentTypeDomain(response schema.PaymentType) *domain.PaymentType {
	return &domain.PaymentType{
		ID:       response.ID,
		TypeName: response.TypeName,
	}
}

func ConvertPaymentTypeResponse(response []domain.PaymentType) []web.PaymentTypeResponse {
	var results []web.PaymentTypeResponse

	for _, paymentType := range response {

		paymentTypeResponse := web.PaymentTypeResponse{
			ID:       paymentType.ID,
			TypeName: paymentType.TypeName,
		}
		results = append(results, paymentTypeResponse)
	}
	return results
}
