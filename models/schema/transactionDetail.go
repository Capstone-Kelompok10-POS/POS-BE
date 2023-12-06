package schema

type TransactionDetail struct {
	ID              uint      	  `gorm:"primaryKey"`
	TransactionID   uint          `gorm:"index;not null"`
	ProductDetailID uint          `gorm:"index;not null"`
	ProductDetail   ProductDetail `gorm:"foreignKey:ProductDetailID"`
	Price           float64       `json:"price" gorm:"type:decimal(10,2);not null"`
	Quantity        int           `json:"quantity" gorm:"not null"`
	SubTotal        float64       `json:"subtotal" gorm:"type:decimal(10,2);not null"`
	Notes           string        `json:"notes"`
}