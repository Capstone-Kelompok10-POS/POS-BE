package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func ProductDetailSchemaToProductDetailDomain(response schema.ProductDetail) *domain.ProductDetail {
	return &domain.ProductDetail{
		ID:         response.ID,
		ProductID:  response.ProductID,
		Price:      response.Price,
		TotalStock: response.TotalStock,
		Size:       response.Size,
	}
}

func ProductDetailDomainToProductDetailResponses(response *domain.ProductDetail) web.ProductDetailResponse {
	return web.ProductDetailResponse{
		ID:         response.ID,
		ProductID:  response.ProductID,
		Price:      response.Price,
		TotalStock: response.TotalStock,
		Size:       response.Size,
	}
}

func ProductDetailDomainToProductDetailCreateResponses(response *domain.ProductDetail) web.ProductDetailCreateResponse {
	return web.ProductDetailCreateResponse{
		ID:         response.ID,
		ProductID:  response.ProductID,
		Price:      response.Price,
		TotalStock: response.TotalStock,
		Size:       response.Size,
	}
}

func ProductDetailDomainToProductDetailPreload(response []domain.ProductDetail) []domain.ProductDetailPreload {
	var results []domain.ProductDetailPreload
	for _, productDetail := range response {
		productDetailPreload := domain.ProductDetailPreload{
			ID:         productDetail.ID,
			ProductID:  productDetail.ProductID,
			Price:      productDetail.Price,
			TotalStock: productDetail.TotalStock,
			Size:       productDetail.Size,
		}
		results = append(results, productDetailPreload)
	}
	return results
}

func ConvertProductDetailResponse(response []domain.ProductDetail) []web.ProductDetailResponse {
	var results []web.ProductDetailResponse
	for _, productDetail := range response {
		membershipResponse := web.ProductDetailResponse{
			ID:         productDetail.ID,
			ProductID:  productDetail.ProductID,
			Price:      productDetail.Price,
			TotalStock: productDetail.TotalStock,
			Size:       productDetail.Size,
		}
		results = append(results, membershipResponse)
	}
	return results
}
