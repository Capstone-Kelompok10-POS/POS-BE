package web

import "qbills/models/domain"


type TransactionCreateRequest struct {
	CashierID 	   uint 					     `json:"cashierId" validate:"numeric"`
	MembershipID   uint                          `json:"membershipId" validate:"numeric"`
	ConvertPointID uint                          `json:"convertPointId" validate:"numeric"`
	Details        []domain.TransactionDetail    `json:"details" validate:"required"`
	TotalPrice 	   float64 				         `json:"totalPrice" validate:"numeric"`
	Discount 	   float64 						 `json:"discount" validate:"numeric"`
	Tax 		   float64 					     `json:"tax" validate:"numeric"`
	TotalPayment   float64 						 `json:"totalPayment" validate:"numeric"`
	TransactionPayment domain.TransactionPayment `json:"transactionPayment" validate:"required"` 
}

type TransactionCreate struct {
	CashierID 	 uint 		`json:"cashierId"`
	Discount 	 float64 	`json:"discount"`
	TotalPrice   float64    `json:"total"`
	Tax 		 float64 	`json:"tax"`
	TotalPayment float64 	`json:"totalPayment"`
}