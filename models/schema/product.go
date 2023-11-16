package schema

import (
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ProductTypeID uint        `gorm:"index;not null"`
	ProductType   ProductType `gorm:"foreignKey:ProductTypeID"`
	AdminID       uint        `gorm:"index;not null"`
	Admin         Admin       `gorm:"foreignKey:AdminID"`
	Name          string      `json:"name" gorm:"not null"`
	Ingredients   string      `json:"ingredients" gorm:"not null"`
	Price         float64     `json:"price" gorm:"type:decimal(10,2);not null"`
	TotalStock    int         `json:"totalStock" gorm:"not null"`
	Size          string      `json:"size" gorm:"type:ENUM('SMALL', 'NORMAL', 'LARGE');not null;default:'NORMAL'"`
	Image         string      `json:"image" gorm:"not null"`
}
