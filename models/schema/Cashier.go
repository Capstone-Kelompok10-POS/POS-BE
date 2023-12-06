package schema 

import (
	"time"

	"gorm.io/gorm"
)

type Cashier struct {
	ID			uint		 `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	AdminID    uint	`gprm:"index"`
	Admin Admin `gorm:"foreignKey:AdminID"`
	Fullname	string `json:"fullname"`
	Username	string `json:"username"`
	Password	string `json:"password"`
}