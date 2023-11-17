package web

type StockCreateRequest struct {
	ProductID uint `json:"productID"`
	Stock     int  `json:"stock"`
}
