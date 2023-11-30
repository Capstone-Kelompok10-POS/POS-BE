package domain

type Product struct {
	ID            uint
	ProductTypeID uint
	ProductType   ProductType
	AdminID       uint
	Admin         Admin
	Name          string
	Ingredients   string
	Image         string
}

type ProductResponse struct {
	ID            uint
	ProductTypeID uint
	ProductType   ProductType
	AdminID       uint
	Admin         AdminResponse
	Name          string
	Ingredients   string
	Image         string
}

type ProductPreloadResponse struct {
	ID            uint
	ProductTypeID uint
	AdminID       uint
	Name          string
	Ingredients   string
	Image         string
}
