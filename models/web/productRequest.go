package web

import "qbills/models/domain"

type ProductCreateRequest struct {
<<<<<<< Updated upstream
	ProductTypeID uint   `json:"productTypeId"`
	AdminID       uint   `json:"adminId"`
=======
	ProductTypeID uint   `json:"productTypeId" gorm:"index;not null"`
	AdminID       uint   `json:"adminId" gorm:"index;not null"`
	Name          string `json:"name" gorm:"not null"`
	Ingredients   string `json:"ingredients" gorm:"not null"`
	Image         string `json:"image" gorm:"not null"`
	ProductDetail domain.ProductDetail `json:"productDetail"`
}

type ProductUpdateRequest struct {
	AdminID       uint   `json:"adminID" gorm:"index;not null"`
	ProductTypeID uint   `json:"productTypeID"`
>>>>>>> Stashed changes
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients"`
	Image         string `json:"image"`
}

type ProductUpdateRequest struct {
	AdminID       uint   `json:"adminId"`
	ProductTypeID uint   `json:"productTypeId"`
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients"`
	Image         string `json:"image"`
}