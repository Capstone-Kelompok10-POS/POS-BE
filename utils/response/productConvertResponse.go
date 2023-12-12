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
		Ingredients:   product.Ingredients,
		Image:         product.Image,
	}
}

func ProductDomainToProductResponse(product *domain.Product) web.ProductResponse {
	productDetail := ProductDetailDomainToProductDetailPreload(product.ProductDetail)
	admin := AdminDomainToAdminDomainResponse(product.Admin)

	return web.ProductResponse{
		ID:            product.ID,
		ProductType:   product.ProductType,
		Admin:         admin,
		Name:          product.Name,
		Ingredients:   product.Ingredients,
		Image:         product.Image,
		ProductDetail: productDetail,
	}
}

func ProductDomainToProductCreateResponse(product *domain.Product) web.ProductCreateResponse {
	return web.ProductCreateResponse{
		ID:            product.ID,
		ProductTypeID: product.ProductTypeID,
		AdminID:       product.AdminID,
		Name:          product.Name,
		Ingredients:   product.Ingredients,
		Image:         product.Image,
	}
}

func ProductDomainToProductUpdateResponse(product *domain.Product) web.ProductUpdateResponse {
	return web.ProductUpdateResponse{
		ID:            product.ID,
		ProductTypeID: product.ProductTypeID,
		Name:          product.Name,
		Ingredients:   product.Ingredients,
		Image:         product.Image,
	}
}

func BestProductsDomainToProductsResponse(bestProduct *domain.BestSellingProduct) *web.BestProductsResponse {
	response := &web.BestProductsResponse{
		ProductID: bestProduct.ProductID,
		ProductName:bestProduct.ProductName,
		ProductImage:bestProduct.ProductImage,
		ProductSize: bestProduct.ProductSize,
		ProductPrice:bestProduct.ProductPrice,
		ProductTypeName:bestProduct.ProductTypeName,
		TotalQuantity:bestProduct.TotalQuantity,
		Amount:bestProduct.Amount,
	}
	return response
}

func ProductsDomainToProductsResponse(product *domain.Product) *web.ProductsResponse {
	response := &web.ProductsResponse{
		ID: product.ID,
		ProductType: web.ProductTypeResponse{
			ID:              product.ProductTypeID,
			TypeName:        product.ProductType.TypeName,
			TypeDescription: product.ProductType.TypeDescription,
		},
		Name:          product.Name,
		Ingredients:   product.Ingredients,
		Image:         product.Image,
		ProductDetail: make([]web.ProductDetailResponse, 0),
	}

	for _, detail := range product.ProductDetail {
		response.ProductDetail = append(response.ProductDetail, web.ProductDetailResponse{
			ID:         detail.ID,
			ProductID:  detail.ProductID,
			Price:      detail.Price,
			TotalStock: detail.TotalStock,
			Size:       detail.Size,
		})
	}
	return response
}
func ConvertBestProductResponse(bestProducts []domain.BestSellingProduct) []web.BestProductsResponse {
	var results []web.BestProductsResponse

	for _, bestProduct := range bestProducts {
		results = append(results, *BestProductsDomainToProductsResponse(&bestProduct))
	}
	return results
}

func ConvertProductResponse(products []domain.Product) []web.ProductsResponse {
	var results []web.ProductsResponse

	for _, product := range products {
		results = append(results, *ProductsDomainToProductsResponse(&product))
	}
	return results
}

func ProductResponseToProductCostumResponse(product web.ProductResponse) web.ProductResponseCustom {
	return web.ProductResponseCustom{
		ID:          product.ID,
		ProductType: product.ProductType,
		Admin:       product.Admin,
		Name:        product.Name,
		Ingredients: product.Ingredients,
		Image:       product.Image,
		//ProductDetail: product.ProductDetail,
	}
}

func ProductResponseToProductsCostumResponse(product web.ProductResponse) web.ProductResponseCustom {
	return web.ProductResponseCustom{
		ID:          product.ID,
		ProductType: product.ProductType,
		Admin:       product.Admin,
		Name:        product.Name,
		Ingredients: product.Ingredients,
		Image:       product.Image,
	}
}

func ProductResponseToProductPreloadResponse(response domain.Product) domain.ProductPreloadResponse {
	return domain.ProductPreloadResponse{
		ID:            response.ID,
		ProductTypeID: response.ProductTypeID,
		AdminID:       response.AdminID,
		Name:          response.Name,
		Ingredients:   response.Ingredients,
		Image:         response.Image,
		ProductDetail: response.ProductDetail,
	}
}

func DomainProductToDomainResponseProduct(response domain.Product) domain.ProductPreloadResponse {
	return domain.ProductPreloadResponse{
		ID:            response.ID,
		ProductTypeID: response.ProductTypeID,
		AdminID:       response.AdminID,
		Name:          response.Name,
		Ingredients:   response.Ingredients,
		Image:         response.Image,
		ProductDetail: response.ProductDetail,
	}
}
