package domain

type Admin struct {
	ID           uint   `json:"id"`
	SuperAdminID uint   `json:"superAdminID"`
	FullName     string `json:"fullName"`
	Username     string `json:"username"`
	Password     string `json:"password"`
}
type AdminResponse struct {
	ID           uint   `json:"id"`
	SuperAdminID uint   `json:"superAdminID"`
	FullName     string `json:"fullName"`
	Username     string `json:"username"`
}
