package domain

type ProductDetail struct {
	ID         uint    `json:"id"`
	ProductID  uint    `json:"productID"`
	Price      float64 `json:"price"`
	TotalStock int     `json:"totalStock"`
	Size       string  `json:"size"`
}
