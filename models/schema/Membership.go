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

	CashierID   uint      `gorm:"index"`
	Cashier     Cashier   `gorm:"foreignKey:CashierID"`
	Name        string    `gorm:"name"`
	CodeMember  uuid.UUID `gorm:"type:char(36);notnull"`
	TotalPoint  uint      `json:"totalPoint"`
	PhoneNumber string    `json:"phoneNumber"`
	Barcode     string    `json:"barcode"`
	MembershipPoint []MembershipPoint `gorm:"foreignKey:MembershipID"`

}

func (membership *Membership) BeforeCreate(tx *gorm.DB) error {
	membership.CodeMember = uuid.NewV4()
	return nil
}
