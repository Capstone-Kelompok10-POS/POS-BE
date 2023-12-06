package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func StockCreateRequestToStockDomain(request web.StockCreateRequest) *domain.Stock {
	return &domain.Stock{
		ProductDetailID: request.ProductDetailID,
		Stock:           request.Stock,
	}
}

func StockDomainToStockSchema(request domain.Stock) *schema.Stock {
	return &schema.Stock{
		ID:              request.ID,
		CreatedAt:       request.CreatedAt,
		ProductDetailID: request.ProductDetailID,
		Stock:           request.Stock,
	}
}
