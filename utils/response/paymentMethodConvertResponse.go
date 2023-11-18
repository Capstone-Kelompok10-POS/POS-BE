package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func PaymentMethodDomainToPaymentMethodResponse(response *domain.PaymentMethod) web.PaymentMethodResponse {
	return web.PaymentMethodResponse{
		PaymentType: response.PaymentType,
		Name:        response.Name,
	}
}

func PaymentMethodSchemaToPaymentMethodDomain(response schema.PaymentMethod) *domain.PaymentMethod {
	return &domain.PaymentMethod{
		ID:          response.ID,
		PaymentType: response.PaymentType,
		Name:        response.Name,
	}
}

func ConvertPaymentMethodResponse(response []domain.PaymentMethod) []web.PaymentMethodResponse {
	var results []web.PaymentMethodResponse

	for _, paymentMethod := range response {

		paymentMethodResponse := web.PaymentMethodResponse{
			ID:          paymentMethod.ID,
			PaymentType: paymentMethod.PaymentType,
			Name:        paymentMethod.Name,
		}
		results = append(results, paymentMethodResponse)
	}
	return results
}
