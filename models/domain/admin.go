package domain

type Admin struct {
	ID           uint
	SuperAdminID uint
	FullName     string
	Username     string
	Password     string
}
type AdminResponse struct {
	ID           uint   `json:"id"`
	SuperAdminID uint   `json:"superAdminID"`
	FullName     string `json:"fullName"`
	Username     string `json:"username"`
}
