package web

type ProductDetailResponse struct {
	ID         uint
	ProductID  uint
	Price      float64 `json:"price" `
	TotalStock int     `json:"totalStock"`
	Size       string  `json:"size"`
}

type ProductDetailCreateResponse struct {
	ID         uint
	ProductID  uint
	Price      float64 `json:"price" `
	TotalStock int     `json:"totalStock"`
	Size       string  `json:"size"`
}
