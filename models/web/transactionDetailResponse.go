package web

type TransactionDetailResponse struct {
	ID              uint 							 `json:"id"`
	TransactionID   uint 							 `json:"transactionId"`
	ProductDetailID uint 							 `json:"productDetailId"`
	ProductDetail   ProductDetailTransactionResponse `json:"productDetail"`
	Price           float64 						 `json:"price"`
	Quantity        int     						 `json:"quantity"`
	SubTotal        float64 						 `json:"subTotal"`
	Notes           string  						 `json:"notes"`
}