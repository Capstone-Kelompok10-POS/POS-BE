package web

type ProductCreateRequest struct {
	ProductTypeID uint    `json:"productTypeID"`
	AdminID       uint    `json:"adminID"`
	Name          string  `json:"name"`
	Description   string  `json:"description" gorm:"not null"`
	Price         float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	TotalStock    int     `json:"totalStock"`
	Size          string  `json:"size"`
	Image         string  `json:"image"`
}

type ProductUpdateRequest struct {
	ProductTypeID uint    `json:"productTypeID"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	TotalStock    int     `json:"totalStock"`
	Size          string  `json:"size"`
	Image         string  `json:"image"`
}
