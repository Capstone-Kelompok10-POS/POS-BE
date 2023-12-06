package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func PaymentMethodDomainToPaymentMethodResponse(response *domain.PaymentMethod) web.PaymentMethodResponse {
	return web.PaymentMethodResponse{
		ID: response.ID,
		PaymentTypeID: response.PaymentTypeID,
		Name:          response.Name,
	}
}

func PaymentMethodSchemaToPaymentMethodDomain(response schema.PaymentMethod) *domain.PaymentMethod {
	return &domain.PaymentMethod{
		ID:            response.ID,
		PaymentTypeID: response.PaymentTypeID,
		Name:          response.Name,
	}
}

func ConvertPaymentMethodResponse(response []domain.PaymentMethod) []web.PaymentMethodResponse {
	var results []web.PaymentMethodResponse

	for _, paymentMethod := range response {

		paymentMethodResponse := web.PaymentMethodResponse{
			ID:            paymentMethod.ID,
			PaymentTypeID: paymentMethod.PaymentTypeID,
			Name:          paymentMethod.Name,
		}
		results = append(results, paymentMethodResponse)
	}
	return results
}
