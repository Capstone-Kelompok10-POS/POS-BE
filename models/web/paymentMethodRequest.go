package web

type PaymentMethodRequest struct {
	PaymentTypeID uint   `gorm:"not null" json:"paymentTypeId"`
	Name          string `gorm:"not null" json:"name"`
}
