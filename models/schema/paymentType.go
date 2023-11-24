package schema

type PaymentType struct {
	ID       uint   `gorm:"primaryKey"`
	TypeName string `gorm:"not null"`
}
