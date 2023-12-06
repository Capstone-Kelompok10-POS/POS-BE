package web

import "qbills/models/domain"

type ProductResponse struct {
	ID            uint               `json:"Id"`
	ProductTypeID uint               `json:"productTypeId"`
	ProductType   domain.ProductType `json:"productType"`
	AdminID       uint               `json:"adminId"`
	Admin         domain.Admin       `json:"admin"`
	ProductDetail []domain.ProductDetail
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients"`
	Image         string `json:"image"`
}

type ProductResponseCustom struct {
	ID            uint               `json:"Id"`
	ProductTypeID uint               `json:"productTypeId"`
	ProductType   domain.ProductType `json:"productType"`
	ProductDetail []domain.ProductDetail
	AdminID       uint                 `json:"adminId"`
	Admin         domain.AdminResponse `json:"admin"`
	Name          string               `json:"name"`
	Ingredients   string               `json:"ingredients"`
	Image         string               `json:"image"`
}

type ProductCreateResponse struct {
	ID            uint   `json:"Id"`
	ProductTypeID uint   `json:"productTypeId"`
<<<<<<< Updated upstream
	AdminID       uint   `json:"adminID"`
=======
	AdminID       uint   `json:"adminId"`
>>>>>>> Stashed changes
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients"`
	Image         string `json:"image"`
}

type ProductUpdateResponse struct {
<<<<<<< Updated upstream
	ID            uint   `json:"Id"`
	ProductTypeID uint   `json:"productTypeId"`
=======
	ID            uint   `json:"id"`
	ProductTypeID uint   `json:"productTypeid"`
>>>>>>> Stashed changes
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients"`
	Image         string `json:"image"`
}

type ProductTransactionResponse struct {
	ID            uint   `json:"id"`
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients"`
	Image         string `json:"image"`
}

type ProductsResponse struct {
	ID            uint               `json:"Id"`
	ProductType   ProductTypeResponse `json:"productType"`
	Name          string               `json:"name"`
	Ingredients   string               `json:"ingredients"`
	Image         string               `json:"image"`
	ProductDetail []ProductDetailResponse `json:"productDetail"`
}