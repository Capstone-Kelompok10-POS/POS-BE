package web

type ProductCreateRequest struct {
<<<<<<< Updated upstream
	ProductTypeID uint   `json:"productTypeID" gorm:"index;not null"`
	AdminID       uint   `json:"adminID" gorm:"index;not null"`
	Name          string `json:"name" gorm:"not null"`
	Ingredients   string `json:"ingredients" gorm:"not null"`
	Image         string `json:"image" gorm:"not null"`
}

type ProductUpdateRequest struct {
	AdminID       uint   `json:"adminID" gorm:"index;not null"`
	ProductTypeID uint   `json:"productTypeID"`
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients"`
	Image         string `json:"image"`
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
	AdminID       uint    `json:"adminID" gorm:"index;not null"`
=======
<<<<<<< Updated upstream
=======
	AdminID       uint    `json:"adminID" gorm:"index;not null"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	ProductTypeID uint    `json:"productTypeID"`
	AdminID       uint    `json:"adminID"`
	Name          string  `json:"name"`
	Description   string  `json:"description" gorm:"not null"`
	Price         float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock         uint    `json:"stock"`
	Size          string  `json:"size"`
	Image         string  `json:"image"`
}

type ProductUpdateRequest struct {
	ProductTypeID uint    `json:"productTypeID"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	Stock         uint    `json:"stock"`
	Size          string  `json:"size"`
	Image         string  `json:"image"`
<<<<<<< Updated upstream
=======
=======
	ProductTypeID uint   `json:"productTypeId"`
	AdminID       uint   `json:"adminId"`
	Name          string `json:"name" validate:"required"`
	Ingredients   string `json:"ingredients"`
	Image         string `json:"image"`
}

type ProductUpdateRequest struct {
	AdminID       uint   `json:"adminId"`
	ProductTypeID uint   `json:"productTypeId"`
	Name          string `json:"name"`
	Ingredients   string `json:"ingredients"`
	Image         string `json:"image"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}
