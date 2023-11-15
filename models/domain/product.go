package domain

type Product struct {
	ID            uint
	ProductTypeID uint
	ProductType   ProductType
	AdminID       uint
	Admin         Admin
	Name          string
	Description   string
	Price         float64
	Stock         uint
	Size          string
	Image         string
}

type ProductResponse struct {
	ID            uint
	ProductTypeID uint
	ProductType   ProductType
	AdminID       uint
	Admin         AdminResponse
	Name          string
	Description   string
	Price         float64
	Stock         uint
	Size          string
	Image         string
}

type ProductPreloadResponse struct {
	ID            uint
	ProductTypeID uint
	AdminID       uint
	Name          string
	Description   string
	Price         float64
	Stock         uint
	Size          string
	Image         string
}
