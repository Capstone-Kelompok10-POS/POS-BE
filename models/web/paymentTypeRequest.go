package web

type PaymentTypeRequest struct {
<<<<<<< Updated upstream
	TypeName string `gorm:"not null" json:"typeName"`
=======
<<<<<<< Updated upstream
<<<<<<< Updated upstream
	TypeName string `gorm:"not null" json:"typeName"`
=======
	TypeName string `json:"typeName" validate:"required"`
=======
	TypeName string `json:"typeName" validate:"required"`
=======
	TypeName string `json:"typeName"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}
