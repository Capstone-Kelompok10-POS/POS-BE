package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func ProductDetailCreateToProductDomain(request web.ProductDetailCreate) *domain.ProductDetail {
	return &domain.ProductDetail{
		ProductID:  request.ProductID,
		Price:      request.Price,
		TotalStock: request.TotalStock,
		Size:       request.Size,
	}
}

func ProductDetailDomainToProductDetailSchema(request domain.ProductDetail) schema.ProductDetail {
	return schema.ProductDetail{
		ID:         request.ID,
		ProductID:  request.ProductID,
		Price:      request.Price,
		TotalStock: request.TotalStock,
		Size:       request.Size,
	}
}
