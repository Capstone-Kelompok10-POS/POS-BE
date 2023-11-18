package schema

import (
	"gorm.io/gorm"
	"time"
)

type PaymentMethod struct {
	ID          uint           `gorm:"primaryKey"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	PaymentType uint           `gorm:"not null"`
	Name        string         `gorm:"not null"`
}
