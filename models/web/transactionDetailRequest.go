package web

type TransactionDetailCreateRequest struct {
	ProductDetailID uint   `json:"productDetailId" validate:"required,numeric"`
	Quantity        int    `json:"quantity" validate:"required,numeric,min=1"`
	Notes           string `json:"notes"`
}

type TransactionDetailCreate struct {
	ProductPrice map[uint]float64 `json:"productPrice"`
	SubTotal     map[uint]float64 `json:"subTotal"`
}