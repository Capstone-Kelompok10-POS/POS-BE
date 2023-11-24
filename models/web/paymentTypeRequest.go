package web

type PaymentTypeRequest struct {
	TypeName string `gorm:"not null" json:"typeName"`
}
