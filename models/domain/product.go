package domain

type Product struct {
	ID            uint
	ProductTypeID uint
	ProductType   ProductType
	AdminID       uint
	Admin         Admin
	ProductDetail []ProductDetail
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
	ProductDetail []ProductDetail
}

type ProductPreloadResponse struct {
	ID            uint
	ProductTypeID uint
	AdminID       uint
	Name          string
	Ingredients   string
	Image         string
	ProductDetail []ProductDetail
}
