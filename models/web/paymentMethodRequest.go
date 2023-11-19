package web

type PaymentMethodRequest struct {
	PaymentTypeID uint   `gorm:"not null" json:"paymentType"`
	Name          string `gorm:"not null" json:"name"`
}
