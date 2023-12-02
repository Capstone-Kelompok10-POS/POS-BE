package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func ProductTypeCreateToProductTypeDomain(request web.ProductTypeCreate) *domain.ProductType {
	return &domain.ProductType{
		TypeName:        request.TypeName,
		TypeDescription: request.TypeDescription,
	}
}

func ProductTypeUpdateToProductTypeDomain(request web.ProductTypeUpdate) *domain.ProductType {
	return &domain.ProductType{
		TypeName:        request.TypeName,
		TypeDescription: request.TypeDescription,
	}

}

func ProductTypeDomainToProductTypeSchema(request domain.ProductType) *schema.ProductType {
	return &schema.ProductType{
		ID:              request.ID,
		TypeName:        request.TypeName,
		TypeDescription: request.TypeDescription,
	}
}
