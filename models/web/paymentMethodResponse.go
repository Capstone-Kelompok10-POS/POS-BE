package web

type PaymentMethodResponse struct {
	ID            uint   `json:"id"`
<<<<<<< Updated upstream
	PaymentTypeID uint   `gorm:"not null" json:"paymentTypeID"`
	Name          string `gorm:"not null" json:"name"`
=======
<<<<<<< Updated upstream
	PaymentTypeID uint   `gorm:"not null" json:"paymentTypeID"`
	Name          string `gorm:"not null" json:"name"`
=======
	PaymentTypeID uint   `json:"paymentTypeId"`
	Name          string `json:"name"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}
