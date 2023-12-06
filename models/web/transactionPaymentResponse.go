package web

import "time"

type TransactionPaymentResponse struct {
	ID            uint                  `json:"id"`
	TransactionID uint                  `json:"transactionId"`
	CreatedAt     time.Time             `json:"createdAt"`
	UpdateAt      time.Time             `json:"updateAt"`
	PaymentMethod PaymentMethodResponse `json:"paymentmethod"`
	Invoice       string                `json:"invoice"`
	VANumber      string                `json:"vaNumber"`
	PaymentStatus string                `json:"paymentStatus"`
}