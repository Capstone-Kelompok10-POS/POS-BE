package schema

import (
	"time"

	"gorm.io/gorm"
)

type TransactionPayment struct {
	ID              uint      	   `gorm:"primaryKey"`
	CreatedAt       time.Time 	   `gorm:"autoCreateTime"`
	UpdatedAt 		time.Time 	   `gorm:"autoUpdateTime"`
	DeletedAt 		gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	
	TransactionID   uint           `gorm:"index;not null"`
	PaymentMethodID uint 		   `gorm:"index;not null"`
	PaymentMethod PaymentMethod    `gorm:"foreignKey:PaymentMethodID"`
	Invoice         string         `json:"invoice" gorm:"not null"`
	PaymentStatus   string         `gorm:"type:varchar(10)" json:"paymentStatus"`
	VANumber      string         `json:"vaNumber"`
}

