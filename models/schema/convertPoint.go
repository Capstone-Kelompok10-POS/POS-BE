package schema

import (
	"time"

	"gorm.io/gorm"
)

type ConvertPoint struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Point int `json:"point"`
	ValuePoint int `json:"valuePoint"`
}