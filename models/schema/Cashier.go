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

	Admin_ID    uint	`gprm:"index"`
	Admin Admin `gorm:"foreignKey:Admin_ID"`
	Fullname	string `json:"fullname"`
	Username	string `json:"username"`
	Password	string `json:"password"`
}