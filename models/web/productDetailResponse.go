package web

type ProductDetailResponse struct {
	ID         uint    `json:"id"`
	ProductID  uint    `json:"productId"`
	Price      float64 `json:"price" `
	TotalStock int     `json:"totalStock"`
	Size       string  `json:"size"`
}

type ProductDetailCreateResponse struct {
	ID         uint    `json:"id"`
	ProductID  uint    `json:"productId"`
<<<<<<< Updated upstream
	Price      float64 `json:"price" `
	TotalStock int     `json:"totalStock"`
	Size       string  `json:"size"`
}
type ProductDetailTransactionResponse struct {
<<<<<<< Updated upstream
	ID         uint `json:"id"`
	ProductID  uint `json:"productId"`
	Product    ProductTransactionResponse `json:"product"`
	Price      float64 `json:"price"`
=======
	ID        uint `json:"id"`
	ProductID uint `json:"productId"`
	// Product    ProductTransactionResponse `json:"product"`
=======
>>>>>>> Stashed changes
	Price      float64 `json:"price" `
>>>>>>> Stashed changes
	TotalStock int     `json:"totalStock"`
	Size       string  `json:"size"`
}
type ProductDetailTransactionResponse struct {
	ID         uint                       `json:"id"`
	ProductID  uint                       `json:"productId"`
	Product    ProductTransactionResponse `json:"product"`
	Price      float64                    `json:"price"`
	TotalStock int                        `json:"totalStock"`
	Size       string                     `json:"size"`
}