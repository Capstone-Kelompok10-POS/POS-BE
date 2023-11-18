package web

type PaymentMethodResponse struct {
	ID          uint   `json:"id"`
	PaymentType uint   `gorm:"not null" json:"paymentType"`
	Name        string `gorm:"not null" json:"name"`
}
