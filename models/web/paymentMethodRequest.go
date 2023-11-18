package web

type PaymentMethodRequest struct {
	PaymentType uint   `gorm:"not null" json:"paymentType"`
	Name        string `gorm:"not null" json:"name"`
}
