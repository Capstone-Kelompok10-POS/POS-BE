package web

import "qbills/models/domain"

type ProductResponse struct {
	ID            uint               `json:"ID"`
	ProductTypeID uint               `json:"productTypeID"`
	ProductType   domain.ProductType `json:"productType"`
	AdminID       uint               `json:"adminID"`
	Admin         domain.Admin       `json:"admin"`
	Name          string             `json:"name"`
	Ingredients   string             `json:"ingredients" gorm:"not null"`
	Image         string             `json:"image"`
}

type ProductResponseCustom struct {
	ID            uint                 `json:"ID"`
	ProductTypeID uint                 `json:"productTypeID"`
	ProductType   domain.ProductType   `json:"productType"`
	AdminID       uint                 `json:"adminID"`
	Admin         domain.AdminResponse `json:"admin"`
	Name          string               `json:"name"`
	Ingredients   string               `json:"ingredients" gorm:"not null"`
	Image         string               `json:"image"`
}

type ProductCreateResponse struct {
	ID            uint   `json:"ID"`
	ProductTypeID uint   `json:"productTypeID"`
	AdminID       uint   `json:"adminID"`
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients" gorm:"not null"`
	Image         string `json:"image"`
}

type ProductUpdateResponse struct {
	ID            uint   `json:"ID"`
	ProductTypeID uint   `json:"productTypeID"`
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients" gorm:"not null"`
	Image         string `json:"image"`
}
