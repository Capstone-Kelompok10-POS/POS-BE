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

	ProductTypeID uint        `gorm:"index"`
	ProductType   ProductType `gorm:"foreignKey:ProductTypeID"`
	AdminID       uint        `gorm:"index"`
	Admin         Admin       `gorm:"foreignKey:AdminID"`
	Name          string      `json:"name"`
	Description   string      `json:"description" gorm:"not null"`
	Price         float64     `json:"price" gorm:"type:decimal(10,2);not null"`
	TotalStock    int         `json:"totalStock"`
	Size          string      `json:"size" gorm:"type:ENUM('SMALL', 'MEDIUM', 'LARGE');not null;default:'SMALL'"`
	Image         string      `json:"image"`
}
