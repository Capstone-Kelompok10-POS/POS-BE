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
	CodeMember   uuid.UUID `gorm:"type:char(36);notnull"`
<<<<<<< Updated upstream
	Point        uint      `json:"point"`
<<<<<<< Updated upstream
	Phone_Number string    `json:"phone_number"`
=======
	PhoneNumber string    `json:"phoneNumber"`
	Barcode      string    `json:"barcode"`
=======
<<<<<<< Updated upstream
	Point        uint      `json:"point"`
	Phone_Number string    `json:"phone_number"`
=======
	Point        int      `json:"point"`
	PhoneNumber string    `json:"phoneNumber"`
	Barcode      string    `json:"barcode"`
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
}

func (membership *Membership) BeforeCreate(tx *gorm.DB) error {
	membership.CodeMember = uuid.NewV4()
	return nil
}
