package schema

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID             uint           	  `gorm:"primaryKey"`
	CreatedAt 	   time.Time      	  `gorm:"autoCreateTime"`
	UpdatedAt 	   time.Time      	  `gorm:"autoUpdateTime"`
	DeletedAt 	   gorm.DeletedAt 	  `gorm:"index"`

	CashierID      uint 			  `gorm:"index;not null"`
	Cashier        Cashier 			  `gorm:"foreignKey:CashierID"`
	MembershipID   uint 			  `gorm:"index"`
	Membership     Membership 		  `gorm:"foreignKey:MembershipID"`
	ConvertPointID uint 			  `gorm:"index"`
	ConvertPoint   ConvertPoint 	  `gorm:"foreignKey:ConvertPointID"`

	Status         string             `gorm:"type:ENUM('CANCEL', 'SUCCESS', 'PENDING');not null;default:'PENDING'"`
	Discount       float64            `json:"discount" gorm:"type:decimal;(10,2);not null"`
	TotalPrice     float64            `json:"totalPrice" gorm:"type:decimal;(10,2);not null"`
	Tax            float64            `json:"tax" gorm:"type:decimal;(10,2);not null"`
	TotalPayment   float64            `json:"totalPayment" gorm:"type:decimal;(10,2);not null"`
	Details       []TransactionDetail `gorm:"foreignKey:TransactionID"`
}