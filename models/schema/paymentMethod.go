package schema

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID            uint           `gorm:"primaryKey"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	PaymentTypeID uint           `gorm:"not null"`
	PaymentType   PaymentType    `gorm:"foreignKey:PaymentTypeID"`
	Name          string         `gorm:"not null"`
	// TransactionPayment []TransactionPayment `gorm:"foreignKey:PaymentMethodID"`
}
