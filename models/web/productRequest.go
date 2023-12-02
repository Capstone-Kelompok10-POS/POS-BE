package web

type ProductCreateRequest struct {
	ProductTypeID uint   `json:"productTypeId"`
	AdminID       uint   `json:"adminId"`
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
