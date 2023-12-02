package web

import (
	"time"
)

type TransactionResponse struct {
	ID             uint                          `json:"id"`
	CreatedAt      time.Time 				     `json:"createdAt"`
	Cashier        CashierTransactionResponse    `json:"cashier"`
	Membership     MembershipTransactionResponse `json:"membership"`
	ConvertPointID uint      				     `json:"convertPointId"`
	Status         string    				     `json:"status"`
	Discount       float64   				     `json:"discount"`
	TotalPrice     float64   				     `json:"totalPrice"`
	Tax     	   float64   				     `json:"tax"`
	TotalPayment   float64 					     `json:"totalPayment"`
	Details        []TransactionDetailResponse   `json:"details"`
}