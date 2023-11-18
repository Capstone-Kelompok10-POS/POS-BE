package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func PaymentMethodRequestToPaymentMethodDomain(request web.PaymentMethodRequest) *domain.PaymentMethod {
	return &domain.PaymentMethod{
		PaymentType: request.PaymentType,
		Name:        request.Name,
	}
}

func PaymentMethodDomainToPaymentMethodRequest(request *domain.PaymentMethod) web.PaymentMethodRequest {
	return web.PaymentMethodRequest{
		PaymentType: request.PaymentType,
		Name:        request.Name,
	}
}

func PaymentMethodDomainToPaymentMethodSchema(request *domain.PaymentMethod) schema.PaymentMethod {
	return schema.PaymentMethod{
		ID:          request.ID,
		PaymentType: request.PaymentType,
		Name:        request.Name,
	}
}
