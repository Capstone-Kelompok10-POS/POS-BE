package schema

import (
	"gorm.io/gorm"
	"time"
)

type ProductType struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	TypeName        string `json:"typeName"`
	TypeDescription string `json:"typeDescription"`
}
