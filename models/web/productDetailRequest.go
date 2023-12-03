package web

type ProductDetailCreate struct {
	ProductID  uint    `json:"productId"`
	Price      float64 `json:"price"`
	TotalStock int     `json:"totalStock" validate:"numeric"`
	Size       string  `json:"size"`
}
