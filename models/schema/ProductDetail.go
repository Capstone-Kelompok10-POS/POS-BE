package schema

import (
	"time"

	"gorm.io/gorm"
)

type ProductDetail struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ProductID  uint `gorm:"index"`
	Product    Product `gorm:"foreignKey:ProductID"`
	Price      float64 `json:"price" gorm:"type:decimal(10,2);not null"`
	TotalStock int     `json:"totalStock" gorm:"not null"`
	Size       string  `json:"size" gorm:"type:ENUM('SMALL', 'NORMAL', 'LARGE');not null;default:'NORMAL'"`
}
