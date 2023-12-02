package web

type ProductDetailResponse struct {
<<<<<<< Updated upstream
	ID         uint
	ProductID  uint
=======
	ID         uint    `json:"id"`
	ProductID  uint    `json:"productId"`
>>>>>>> Stashed changes
	Price      float64 `json:"price" `
	TotalStock int     `json:"totalStock"`
	Size       string  `json:"size"`
}

type ProductDetailCreateResponse struct {
<<<<<<< Updated upstream
	ID         uint
	ProductID  uint
=======
	ID         uint    `json:"id"`
	ProductID  uint    `json:"productId"`
>>>>>>> Stashed changes
	Price      float64 `json:"price" `
	TotalStock int     `json:"totalStock"`
	Size       string  `json:"size"`
}
<<<<<<< Updated upstream
=======

type ProductDetailTransactionResponse struct {
	ID         uint 					  `json:"id"`
	ProductID  uint 					  `json:"productId"`
	Product    ProductTransactionResponse `json:"product"`
	Price      float64 					  `json:"price" `
	TotalStock int     					  `json:"totalStock"`
	Size       string  					  `json:"size"`
}
>>>>>>> Stashed changes
