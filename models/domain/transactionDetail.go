package domain

type TransactionDetail struct {
	ID              uint
	TransactionID   uint
	ProductDetailID uint
	ProductDetail   ProductDetail
	Price           float64
	Quantity        int
	SubTotal        float64
	Notes           string
}