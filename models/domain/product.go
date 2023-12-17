package domain

import "time"

type Product struct {
	ID            uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
	ProductTypeID uint
	ProductType   ProductType
	AdminID       uint
	Admin         Admin
	ProductDetail []ProductDetail
	Name          string
	Ingredients   string
	Image         string
}

type ProductResponse struct {
	ID            uint
	ProductTypeID uint
	ProductType   ProductType
	AdminID       uint
	Admin         AdminResponse
	Name          string
	Ingredients   string
	Image         string
	ProductDetail []ProductDetail
}

type ProductPreloadResponse struct {
	ID            uint
	ProductTypeID uint
	AdminID       uint
	Name          string
	Ingredients   string
	Image         string
	ProductDetail []ProductDetail
}

type BestSellingProduct struct {
	ProductID 		uint
	ProductName 	string
	ProductImage 	string
	ProductSize 	string
	ProductPrice 	float64
	ProductTypeName string
	TotalQuantity 	int
	Amount 			float64
}

type ProductRecommendation struct {
	Reply string
}