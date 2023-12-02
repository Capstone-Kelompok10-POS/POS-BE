package schema

import (
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
>>>>>>> Stashed changes
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

<<<<<<< Updated upstream
	ProductTypeID uint        `gorm:"index;not null"`
=======
<<<<<<< Updated upstream
	ProductTypeID uint        `gorm:"index"`
=======
<<<<<<< Updated upstream
=======
=======
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID            uint           `gorm:"primaryKey"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	ProductDetail []ProductDetail `gorm:"foreignKey:ProductID"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	ProductTypeID uint        `gorm:"index;not null"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	ProductType   ProductType `gorm:"foreignKey:ProductTypeID"`
	AdminID       uint        `gorm:"index;not null"`
	Admin         Admin       `gorm:"foreignKey:AdminID"`
<<<<<<< Updated upstream
	Name          string      `json:"name" gorm:"not null"`
	Ingredients   string      `json:"ingredients" gorm:"not null"`
	Image         string      `json:"image" gorm:"not null"`
=======
<<<<<<< Updated upstream
	Name          string      `json:"name"`
	Description   string      `json:"description" gorm:"not null"`
	Price         float64     `json:"price" gorm:"type:decimal(10,2);not null"`
	Stock         uint        `json:"stock"`
	Size          string      `json:"size" gorm:"type:ENUM('SMALL', 'MEDIUM', 'LARGE');not null;default:'SMALL'"`
	Image         string      `json:"image"`
=======
	Name          string      `json:"name" gorm:"not null"`
	Ingredients   string      `json:"ingredients" gorm:"not null"`
<<<<<<< Updated upstream
	Price         float64     `json:"price" gorm:"type:decimal(10,2);not null"`
	TotalStock    int         `json:"totalStock" gorm:"not null"`
	Size          string      `json:"size" gorm:"type:ENUM('SMALL', 'NORMAL', 'LARGE');not null;default:'NORMAL'"`
=======
<<<<<<< Updated upstream
	Price         float64     `json:"price" gorm:"type:decimal(10,2);not null"`
	TotalStock    int         `json:"totalStock" gorm:"not null"`
	Size          string      `json:"size" gorm:"type:ENUM('SMALL', 'NORMAL', 'LARGE');not null;default:'NORMAL'"`
=======
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	Image         string      `json:"image" gorm:"not null"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}
