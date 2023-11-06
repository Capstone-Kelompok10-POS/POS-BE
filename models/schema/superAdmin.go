package schema

import (
	"time"

	"gorm.io/gorm"
)

type SuperAdmin struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name           string `json:"name"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Admin 	   []Admin `gorm:"foreignKey:SuperAdminID"`
}