package web

import (
	"qbills/models/domain"
	"time"
)

type StockResponse struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	ProductID uint           `json:"productID"`
	Product   domain.Product `json:"product"`
	Stock     int            `json:"stock"`
}

type StockResponseCustom struct {
	ID        uint                          `json:"id"`
	CreatedAt time.Time                     `json:"createdAt"`
	ProductID uint                          `json:"productID"`
	Product   domain.ProductPreloadResponse `json:"product"`
	Stock     int                           `json:"stock"`
}
