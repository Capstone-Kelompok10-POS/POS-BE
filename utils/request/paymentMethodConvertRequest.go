package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func PaymentMethodRequestToPaymentMethodDomain(request web.PaymentMethodRequest) *domain.PaymentMethod {
	return &domain.PaymentMethod{
		PaymentTypeID: request.PaymentTypeID,
		Name:          request.Name,
	}
}

func PaymentMethodDomainToPaymentMethodRequest(request *domain.PaymentMethod) web.PaymentMethodRequest {
	return web.PaymentMethodRequest{
		PaymentTypeID: request.PaymentTypeID,
		Name:          request.Name,
	}
}

func PaymentMethodDomainToPaymentMethodSchema(request *domain.PaymentMethod) schema.PaymentMethod {
	return schema.PaymentMethod{
		ID:            request.ID,
		PaymentTypeID: request.PaymentTypeID,
		Name:          request.Name,
	}
}
