package schema

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Membership struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CashierID    uint      `gorm:"index"`
	Cashier      Cashier   `gorm:"foreignKey:CashierID"`
	Name         string    `gorm:"name"`
	Code_Member  uuid.UUID `gorm:"type:char(36);notnull"`
	Point        uint      `json:"point"`
	Phone_Number string    `json:"phone_number"`
}

func (membership *Membership) BeforeCreate(tx *gorm.DB) error {
	membership.Code_Member = uuid.NewV4()
	return nil
}
