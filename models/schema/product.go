package schema

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID            uint           `gorm:"primaryKey"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	ProductDetail []ProductDetail `gorm:"foreignKey:ProductID"`
	ProductTypeID uint        `gorm:"index;not null"`
	ProductType   ProductType `gorm:"foreignKey:ProductTypeID"`
	AdminID       uint        `gorm:"index;not null"`
	Admin         Admin       `gorm:"foreignKey:AdminID"`
	Name          string      `json:"name" gorm:"not null"`
	Ingredients   string      `json:"ingredients" gorm:"not null"`
	Image         string      `json:"image" gorm:"not null"`
}
