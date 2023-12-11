package domain

import "time"

type ProductDetail struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ProductID  uint
	Product    Product
	Price      float64
	TotalStock int
	Size       string
}
