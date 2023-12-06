package request

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func ProductCreateRequestToProductDomain(request web.ProductCreateRequest) *domain.Product {
	newProductDetail := domain.ProductDetail{
		Price: request.ProductDetail.Price,
		Size: request.ProductDetail.Size,
	}
	return &domain.Product{
		ProductTypeID: request.ProductTypeID,
		AdminID:       request.AdminID,
		Name:          request.Name,
		Ingredients:   request.Ingredients,
		Image:         request.Image,
		ProductDetail: []domain.ProductDetail{newProductDetail},
	}
}

func ProductUpdateRequestToProductDomain(request web.ProductUpdateRequest) *domain.Product {
	return &domain.Product{
		ProductTypeID: request.ProductTypeID,
		AdminID:       request.AdminID,
		Name:          request.Name,
		Ingredients:   request.Ingredients,
		Image:         request.Image,
	}
}

func ProductDomainToProductUpdateRequest(request *domain.Product) web.ProductUpdateRequest {
	return web.ProductUpdateRequest{
		ProductTypeID: request.ProductTypeID,
		AdminID:       request.AdminID,
		Name:          request.Name,
		Ingredients:   request.Ingredients,
		Image:         request.Image,
	}
}

func ProductDomainToProductSchema(request domain.Product) *schema.Product {
	return &schema.Product{
		ID:            request.ID,
		ProductTypeID: request.ProductTypeID,
		AdminID:       request.AdminID,
		Name:          request.Name,
		Ingredients:   request.Ingredients,
		Image:         request.Image,
	}
}
