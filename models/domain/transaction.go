package domain

import "time"

type Transaction struct {
	ID             uint
	CreatedAt      time.Time
	CashierID      uint
	Cashier        Cashier
	MembershipID   uint
	Membership     Membership
	ConvertPointID uint
	ConvertPoint   ConvertPoint

	Discount     float64
	TotalPrice   float64
	Tax          float64
	TotalPayment float64
	Details      []TransactionDetail
	TransactionPayment TransactionPayment
}