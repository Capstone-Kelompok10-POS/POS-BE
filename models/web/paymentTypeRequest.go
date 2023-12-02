package web

type PaymentTypeRequest struct {
<<<<<<< Updated upstream
	TypeName string `gorm:"not null" json:"typeName"`
=======
<<<<<<< Updated upstream
	TypeName string `gorm:"not null" json:"typeName"`
=======
	TypeName string `json:"typeName" validate:"required"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}
