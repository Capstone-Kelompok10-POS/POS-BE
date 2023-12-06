package domain

import "time"

type TransactionPayment struct {
	ID            uint
	CreatedAt     time.Time
	UpdatedAt time.Time
	TransactionID uint
	PaymentMethodID uint
	PaymentMethod PaymentMethod
	Invoice       string
	PaymentStatus string
	VANumber    string
}

type PaymentTransactionStatus struct {
	OrderID           string    `json:"order_id"`
	TransactionStatus string    `json:"transaction_status"`
	FraudStatus       string    `json:"fraud_status"`
	SettlementTime    time.Time `json:"settlement_time"`
}