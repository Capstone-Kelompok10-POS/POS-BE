package domain

import "time"

type ProductDetail struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	ProductID  uint      `json:"productID"`
	Product    Product   `json:"product"`
	Price      float64   `json:"price"`
	TotalStock int       `json:"totalStock"`
	Size       string    `json:"size"`
}

type ProductDetailPreload struct {
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	ProductID  uint      `json:"productID"`
	Price      float64   `json:"price"`
	TotalStock int       `json:"totalStock"`
	Size       string    `json:"size"`
}
