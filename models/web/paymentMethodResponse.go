package web

type PaymentMethodResponse struct {
	ID            uint   `json:"id"`
	PaymentTypeID uint   `gorm:"not null" json:"paymentTypeID"`
	Name          string `gorm:"not null" json:"name"`
}
