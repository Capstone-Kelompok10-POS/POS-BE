package domain

type Admin struct {
	ID           uint
	SuperAdminID uint
	FullName     string
	Username     string
	Password     string
}
type AdminResponse struct {
	ID           uint
	SuperAdminID uint
	FullName     string
	Username     string
}
