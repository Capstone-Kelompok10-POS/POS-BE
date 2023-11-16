package schema

import (
	"time"

	"gorm.io/gorm"
)

type Membership struct {
	ID 		  	uint			`gorm:"primaryKey"` 
	CreatedAt time.Time      	`gorm:"autoCreateTime"`
	UpdatedAt time.Time      	`gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt	`gorm:"index"`

	CashierID	uint           	`gorm:"index"`
	Cashier Cashier				`gorm:"foreignKey:CashierID"`
	Name        string          `gorm:"name"`
	Point     	uint 			`json:"point"`
	Telephone 	string         	`json:"telephone"`
}