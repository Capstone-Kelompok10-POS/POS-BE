package web

import "qbills/models/domain"

type ProductCreateRequest struct {
	ProductTypeID uint                 `json:"productTypeId"`
	AdminID       uint                 `json:"adminId"`
	Name          string               `json:"name"`
	Ingredients   string               `json:"ingredients"`
	Image         string               `json:"image"`
	ProductDetail domain.ProductDetail `json:"productDetail"`
}

type ProductUpdateRequest struct {
	AdminID       uint   `json:"adminId"`
	ProductTypeID uint   `json:"productTypeId"`
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients"`
	Image         string `json:"image"`
	ProductDetail domain.ProductDetail `json:"productDetail"`
}