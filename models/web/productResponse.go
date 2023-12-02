package web

import "qbills/models/domain"

type ProductResponse struct {
	ID            uint               `json:"ID"`
	ProductTypeID uint               `json:"productTypeID"`
	ProductType   domain.ProductType `json:"productType"`
	AdminID       uint               `json:"adminID"`
	Admin         domain.Admin       `json:"admin"`
<<<<<<< Updated upstream
	Name          string             `json:"name"`
<<<<<<< Updated upstream
	Ingredients   string             `json:"ingredients" gorm:"not null"`
=======
	Ingredients   string             `json:"Ingredients" gorm:"not null"`
=======
<<<<<<< Updated upstream
	Name          string             `json:"name"`
	Ingredients   string             `json:"ingredients" gorm:"not null"`
<<<<<<< Updated upstream
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	Price         float64            `json:"price" gorm:"type:decimal(10,2);not null"`
	TotalStock    int                `json:"totalStock"`
>>>>>>> Stashed changes
	Size          string             `json:"size"`
>>>>>>> Stashed changes
	Image         string             `json:"image"`
}

type ProductResponseCustom struct {
	ID            uint                 `json:"ID"`
	ProductTypeID uint                 `json:"productTypeID"`
	ProductType   domain.ProductType   `json:"productType"`
<<<<<<< Updated upstream
	AdminID       uint                 `json:"adminID"`
	Admin         domain.AdminResponse `json:"admin"`
	Name          string               `json:"name"`
	Ingredients   string               `json:"ingredients" gorm:"not null"`
<<<<<<< Updated upstream
=======
	Price         float64              `json:"price" gorm:"type:decimal(10,2);not null"`
	TotalStock    int                  `json:"totalStock"`
>>>>>>> Stashed changes
	Size          string               `json:"size"`
=======
=======
	ProductDetail []domain.ProductDetail
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients" gorm:"not null"`
	Image         string `json:"image"`
}

type ProductResponseCustom struct {
	ID            uint               `json:"ID"`
	ProductTypeID uint               `json:"productTypeID"`
	ProductType   domain.ProductType `json:"productType"`
	ProductDetail []domain.ProductDetail
>>>>>>> Stashed changes
	AdminID       uint                 `json:"adminID"`
	Admin         domain.AdminResponse `json:"admin"`
	Name          string               `json:"name"`
	Ingredients   string               `json:"ingredients" gorm:"not null"`
<<<<<<< Updated upstream
	Price         float64              `json:"price" gorm:"type:decimal(10,2);not null"`
	TotalStock    int                  `json:"totalStock"`
	Size          string               `json:"size"`
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	Image         string               `json:"image"`
}

type ProductCreateResponse struct {
<<<<<<< Updated upstream
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
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
>>>>>>> Stashed changes
	ID            uint    `json:"ID"`
	ProductTypeID uint    `json:"productTypeID"`
	AdminID       uint    `json:"adminID"`
	Name          string  `json:"name"`
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	Description   string  `json:"description" gorm:"not null"`
	Price         float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock         uint    `json:"stock"`
=======
=======
>>>>>>> Stashed changes
	Ingredients   string  `json:"ingredients" gorm:"not null"`
=======
<<<<<<< Updated upstream
	Ingredients   string  `json:"Ingredients" gorm:"not null"`
=======
	Ingredients   string  `json:"ingredients" gorm:"not null"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	Price         float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	TotalStock    int     `json:"totalStock"`
>>>>>>> Stashed changes
	Size          string  `json:"size"`
	Image         string  `json:"image"`
}

type ProductUpdateResponse struct {
	ID            uint    `json:"ID"`
	ProductTypeID uint    `json:"productTypeID"`
	Name          string  `json:"name"`
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	Description   string  `json:"description" gorm:"not null"`
	Price         float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock         uint    `json:"stock"`
=======
=======
>>>>>>> Stashed changes
	Ingredients   string  `json:"ingredients" gorm:"not null"`
=======
<<<<<<< Updated upstream
	Ingredients   string  `json:"Ingredients" gorm:"not null"`
=======
	Ingredients   string  `json:"ingredients" gorm:"not null"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	Price         float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	TotalStock    int     `json:"totalStock"`
>>>>>>> Stashed changes
	Size          string  `json:"size"`
	Image         string  `json:"image"`
>>>>>>> Stashed changes
}
<<<<<<< Updated upstream
=======
=======
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

type ProductTransactionResponse struct {
	ID            uint    `json:"id"`
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients" gorm:"not null"`
	Image         string `json:"image"`
}
>>>>>>> Stashed changes
>>>>>>> Stashed changes
