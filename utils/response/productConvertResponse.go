package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func ProductSchemaToProductDomain(product *schema.Product) *domain.Product {
	return &domain.Product{
		ID:            product.ID,
		ProductTypeID: product.ProductTypeID,
		AdminID:       product.AdminID,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		Stock:         product.Stock,
		Size:          product.Size,
		Image:         product.Image,
	}
}

func ProductDomainToProductResponse(product *domain.Product) web.ProductResponse {
	return web.ProductResponse{
		ID:            product.ID,
		ProductTypeID: product.ProductTypeID,
		ProductType:   product.ProductType,
		AdminID:       product.AdminID,
		Admin:         product.Admin,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		Stock:         product.Stock,
		Size:          product.Size,
		Image:         product.Image,
	}
}

func ProductDomainToProductCreateResponse(product *domain.Product) web.ProductCreateResponse {
	return web.ProductCreateResponse{
		ID:            product.ID,
		ProductTypeID: product.ProductTypeID,
		AdminID:       product.AdminID,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		Stock:         product.Stock,
		Size:          product.Size,
		Image:         product.Image,
	}
}

func ProductDomainToProductUpdateResponse(product *domain.Product) web.ProductUpdateResponse {
	return web.ProductUpdateResponse{
		ID:            product.ID,
		ProductTypeID: product.ProductTypeID,
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		Stock:         product.Stock,
		Size:          product.Size,
		Image:         product.Image,
	}
}

func ConvertProductResponse(products []domain.Product) []web.ProductResponse {
	var results []web.ProductResponse

	for _, product := range products {
		productResponse := web.ProductResponse{
			ID:            product.ID,
			ProductTypeID: product.ProductTypeID,
			ProductType:   product.ProductType,
			AdminID:       product.AdminID,
			Admin:         product.Admin,
			Name:          product.Name,
			Description:   product.Description,
			Price:         product.Price,
			Stock:         product.Stock,
			Size:          product.Size,
			Image:         product.Image,
		}
		results = append(results, productResponse)
	}

	return results

}
