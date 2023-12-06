package web

import (
	"qbills/models/domain"
	"time"
)

type StockResponse struct {
	ID              uint                        `json:"id"`
	CreatedAt       time.Time                   `json:"createdAt"`
	ProductDetailID uint                        `json:"productDetailId"`
	ProductDetail   domain.ProductDetail `json:"productDetail"`
	Stock           int                         `json:"stock"`
}

type StockResponseCustom struct {
	ID              uint                        `json:"id"`
	CreatedAt       time.Time                   `json:"createdAt"`
	ProductDetail   domain.ProductDetail `json:"productDetail"`
	Stock           int                         `json:"stock"`
}

type StockResponseCreate struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	ProductDetailID uint      `json:"productDetailId"`
	Stock           int       `json:"stock"`
}
