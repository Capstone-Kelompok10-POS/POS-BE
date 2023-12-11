package domain

import "time"

type Stock struct {
	ID              uint
	CreatedAt       time.Time
	ProductDetailID uint
	ProductDetail   ProductDetail
	Stock           int
}
