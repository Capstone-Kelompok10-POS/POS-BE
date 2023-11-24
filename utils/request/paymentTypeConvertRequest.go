package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func PaymentTypeRequestToPaymentTypeDomain(request web.PaymentTypeRequest) *domain.PaymentType {
	return &domain.PaymentType{
		TypeName: request.TypeName,
	}
}

func PaymentTypeDomainToPaymentTypeRequest(request *domain.PaymentType) web.PaymentTypeRequest {
	return web.PaymentTypeRequest{
		TypeName: request.TypeName,
	}
}

func PaymentTypeDomainToPaymentTypeSchema(request domain.PaymentType) schema.PaymentType {
	return schema.PaymentType{
		ID:       request.ID,
		TypeName: request.TypeName,
	}
}
