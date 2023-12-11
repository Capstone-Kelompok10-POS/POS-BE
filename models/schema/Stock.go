package schema

import (
	"time"
)

type Stock struct {
	ID              uint          `gorm:"primaryKey"`
	CreatedAt       time.Time     `gorm:"autoCreateTime"`
	ProductDetailID uint          `gorm:"index" json:"productID"`
	ProductDetail   ProductDetail `gorm:"foreignKey:ProductDetailID" json:"productDetail"`
	Stock           int           `json:"stock"`
}
