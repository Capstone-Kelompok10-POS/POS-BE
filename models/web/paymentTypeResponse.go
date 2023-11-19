package web

type PaymentTypeResponse struct {
	ID       uint   `json:"id"`
	TypeName string `gorm:"not null" json:"typeName"`
}
