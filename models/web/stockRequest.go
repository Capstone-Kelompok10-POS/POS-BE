package web

type StockCreateRequest struct {
	ProductDetailID uint `json:"productDetailID"`
	Stock           int  `json:"stock"`
}
