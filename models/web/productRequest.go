package web

type ProductCreateRequest struct {
	ProductTypeID uint    `json:"productTypeID" gorm:"index;not null"`
	AdminID       uint    `json:"adminID" gorm:"index;not null"`
	Name          string  `json:"name" gorm:"not null"`
	Ingredients   string  `json:"ingredients" gorm:"not null"`
	Price         float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	TotalStock    int     `json:"totalStock" gorm:"not null"`
	Size          string  `json:"size"`
	Image         string  `json:"image" gorm:"not null"`
}

type ProductUpdateRequest struct {
	ProductTypeID uint    `json:"productTypeID"`
	Name          string  `json:"name"`
	Ingredients   string  `json:"ingredients"`
	Price         float64 `json:"price"`
	TotalStock    int     `json:"totalStock"`
	Size          string  `json:"size"`
	Image         string  `json:"image"`
}
