package domain

type ProductDetail struct {
	ID         uint
	ProductID  uint
	Product    Product
	Price      float64
	TotalStock int
	Size       string
}
