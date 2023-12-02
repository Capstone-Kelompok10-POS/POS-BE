package domain

import "time"

type Stock struct {
	ID        uint
	CreatedAt time.Time
	ProductID uint
	Product   Product
	Stock     int
}

type StockResponse struct {
	ID        uint
	CreatedAt time.Time
	ProductID uint
	Product   Product
	Stock     int
}
