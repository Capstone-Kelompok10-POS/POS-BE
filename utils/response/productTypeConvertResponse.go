package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func ProductTypeDomainToProductTypeResponse(productType *domain.ProductType) *web.ProductTypeResponse {
	return &web.ProductTypeResponse{
		ID:              productType.ID,
		TypeName:        productType.TypeName,
		TypeDescription: productType.TypeDescription,
	}

}

func ProductTypeSchemaToProductTypeDomain(productType *schema.ProductType) *domain.ProductType {
	return &domain.ProductType{
		ID:              productType.ID,
		TypeName:        productType.TypeName,
		TypeDescription: productType.TypeDescription,
	}
}

func ConvertProductTypeResponse(productTypes []domain.ProductType) []web.ProductTypeResponse {
	var results []web.ProductTypeResponse

	for _, productType := range productTypes {
		ProductTypeResponse := web.ProductTypeResponse{
			ID:              productType.ID,
			TypeName:        productType.TypeName,
			TypeDescription: productType.TypeDescription,
		}
		results = append(results, ProductTypeResponse)
	}
	return results
}
