package web

type ProductDetailCreate struct {
	ProductID  uint    `json:"productID"`
	Price      float64 `json:"price" `
	TotalStock int     `json:"totalStock"`
	Size       string  `json:"size"`
}
