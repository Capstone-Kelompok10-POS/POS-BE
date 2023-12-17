package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func StockDomainToStockResponse(response *domain.Stock) web.StockResponse {
	return web.StockResponse{
		ID:              response.ID,
		CreatedAt:       response.CreatedAt,
		ProductDetailID: response.ProductDetailID,
		Stock:           response.Stock,
	}
}

func StockSchemaToStockDomain(response *schema.Stock) *domain.Stock {
	return &domain.Stock{
		ID:              response.ID,
		CreatedAt:       response.CreatedAt,
		ProductDetailID: response.ProductDetailID,
		Stock:           response.Stock,
	}
}

func StockResponseToStockResponseCustom(response web.StockResponse) web.StockResponseCustom {

	return web.StockResponseCustom{
		ID:              response.ID,
		CreatedAt:       response.CreatedAt,
		ProductDetailID: response.ProductDetailID,
		Stock:           response.Stock,
	}
}

func StockResponseToStockResponseCreate(response web.StockResponse) web.StockResponseCreate {

	return web.StockResponseCreate{
		ID:              response.ID,
		CreatedAt:       response.CreatedAt,
		ProductDetailID: response.ProductDetailID,
		Stock:           response.Stock,
	}
}

func ConvertStockResponse(Stock []domain.Stock) []web.StockResponseCustom {
	var results []web.StockResponseCustom
	for _, stock := range Stock {

		stockResponse := web.StockResponseCustom{
			ID:              stock.ID,
			CreatedAt:       stock.CreatedAt,
			ProductDetailID: stock.ProductDetailID,
			Stock:           stock.Stock,
		}
		results = append(results, stockResponse)
	}
	return results
}
