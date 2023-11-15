package response

import (
	"qbills/models/domain"
	"qbills/models/schema"
	"qbills/models/web"
)

func StockDomainToStockResponse(response *domain.Stock) web.StockResponse {
	return web.StockResponse{
		ID:        response.ID,
		CreatedAt: response.CreatedAt,
		ProductID: response.ProductID,
		Product:   response.Product,
		Stock:     response.Stock,
	}
}

func StockSchemaToStockDomain(response *schema.Stock) *domain.Stock {
	return &domain.Stock{
		ID:        response.ID,
		CreatedAt: response.CreatedAt,
		ProductID: response.ProductID,
		Stock:     response.Stock,
	}
}

func StockResponseToStockResponseCustom(response web.StockResponse) web.StockResponseCustom {

	res := ProductResponseToProductPreloadResponse(response.Product)
	return web.StockResponseCustom{
		ID:        response.ID,
		CreatedAt: response.CreatedAt,
		ProductID: response.ProductID,
		Product:   res,
		Stock:     response.Stock,
	}
}

func StockResponseToStockResponseCreate(response web.StockResponse) web.StockResponseCreate {

	return web.StockResponseCreate{
		ID:        response.ID,
		CreatedAt: response.CreatedAt,
		ProductID: response.ProductID,
		Stock:     response.Stock,
	}
}

func ConvertStockResponse(Stock []domain.Stock) []web.StockResponseCustom {
	var results []web.StockResponseCustom
	for _, stock := range Stock {
		res := DomainProductToDomainResponseProduct(stock.Product)
		stockResponse := web.StockResponseCustom{
			ID:        stock.ID,
			CreatedAt: stock.CreatedAt,
			ProductID: stock.ProductID,
			Product:   res,
			Stock:     stock.Stock,
		}
		results = append(results, stockResponse)
	}
	return results
}
