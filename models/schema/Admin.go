package schema

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	SuperAdminID uint `gorm:"index"`
	SuperAdmin SuperAdmin `gorm:"foreignKey:SuperAdminID"`
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}