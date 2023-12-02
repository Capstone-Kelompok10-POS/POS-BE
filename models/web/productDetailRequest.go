package web

type ProductDetailCreate struct {
<<<<<<< Updated upstream
	ProductID  uint    `json:"productID"`
	Price      float64 `json:"price" `
	TotalStock int     `json:"totalStock"`
=======
	ProductID  uint    `json:"productId"`
	Price      float64 `json:"price"`
	TotalStock int     `json:"totalStock" validate:"numeric"`
>>>>>>> Stashed changes
	Size       string  `json:"size"`
}
