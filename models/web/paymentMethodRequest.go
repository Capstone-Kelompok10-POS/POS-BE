package web

type PaymentMethodRequest struct {
<<<<<<< Updated upstream
	PaymentTypeID uint   `gorm:"not null" json:"paymentTypeId"`
	Name          string `gorm:"not null" json:"name"`
=======
<<<<<<< Updated upstream
	PaymentTypeID uint   `gorm:"not null" json:"paymentTypeId"`
	Name          string `gorm:"not null" json:"name"`
=======
	PaymentTypeID uint   `json:"paymentTypeId" validate:"required,number"`
	Name          string `json:"name" validate:"required"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}
