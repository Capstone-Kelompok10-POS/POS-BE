package schema

import (
	"time"
)

type Stock struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	ProductID uint    `gorm:"index" json:"productID"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	Stock     int     `json:"stock"`
}
