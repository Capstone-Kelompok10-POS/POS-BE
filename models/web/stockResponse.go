package web

import (
	"time"
)

type StockResponse struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	ProductDetailID uint      `json:"productDetailID"`
	Stock           int       `json:"stock"`
}

type StockResponseCustom struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	ProductDetailID uint      `json:"productDetailID"`
	Stock           int       `json:"stock"`
}

type StockResponseCreate struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	ProductDetailID uint      `json:"productDetailID"`
	Stock           int       `json:"stock"`
}
